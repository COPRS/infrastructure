package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"

	"github.com/urfave/cli/v2"

	"google.golang.org/grpc"

	eg "github.com/COPRS/rs-infra-scaler/protos"
)

var (
	inventory_path      = "/opt/rs-infra-scaler/inventory"
	infrastructure_path = "/opt/rs-infra-scaler/infrastructure"
)

type RSInfraScaler struct {
	eg.UnimplementedCloudProviderServer
	infraShellClient     *InfraShellClient
	safescaleShellClient *SafescaleShellClient
}

func main() {

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    "listen-port",
				Aliases: []string{"l"},
				Value:   8086,
				Usage:   "grpc service listen port",
			},
			&cli.IntFlag{
				Name:    "verbosity",
				Aliases: []string{"v"},
				Value:   1,
				Usage:   "verbosity",
			},
			&cli.StringFlag{
				Name:    "tenant",
				Aliases: []string{"t"},
				Value:   "tenant",
				Usage:   "safescale tenant",
			},
		},
		Action: startRSInfraScaler,
	}

	err := app.Run(os.Args)
	if err != nil {
		klog.Fatal(err)
	}
}

func startRSInfraScaler(c *cli.Context) error {

	fs := flag.NewFlagSet("", flag.PanicOnError)
	klog.InitFlags(fs)
	fs.Set("v", strconv.Itoa(c.Int("verbosity")))

	grpcServer := grpc.NewServer()
	ssc := NewSafescaleShellClient(c.String("tenant"))
	isc := NewInfraShellClient(inventory_path, infrastructure_path, ssc)
	ia := &RSInfraScaler{
		infraShellClient:     isc,
		safescaleShellClient: ssc,
	}

	eg.RegisterCloudProviderServer(grpcServer, ia)

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", c.Int("listen-port")))
	if err != nil {
		klog.Fatalf("failed to listen: %v", err)
	}

	klog.V(1).Infof("RSInfraScaler gRPC service listening on %s", lis.Addr())
	grpcServer.Serve(lis)

	return nil
}

// NodeGroups returns all node groups configured for this cloud provider.
func (c *RSInfraScaler) NodeGroups(ctx context.Context, req *eg.NodeGroupsRequest) (*eg.NodeGroupsResponse, error) {
	klog.V(5).Infof("NodeGroups gRPC function called")

	var nodeGroups []*eg.NodeGroup

	for nodeGroupName, infraNodeGroup := range c.infraShellClient.InfraNodeGroups {
		nodeGroups = append(nodeGroups, &eg.NodeGroup{
			Id:      nodeGroupName,
			MinSize: int32(infraNodeGroup.MinSize),
			MaxSize: int32(infraNodeGroup.MaxSize),
		})
	}

	return &eg.NodeGroupsResponse{
		NodeGroups: nodeGroups,
	}, nil
}

// NodeGroupForNode returns the node group for the given node.
// The node group id is an empty string if the node should not
// be processed by cluster autoscaler.
func (c *RSInfraScaler) NodeGroupForNode(ctx context.Context, req *eg.NodeGroupForNodeRequest) (*eg.NodeGroupForNodeResponse, error) {
	klog.V(5).Infof("NodeGroupForNode gRPC function called for: %s", req.Node.Name)

	var nodeGroupName string
	var ok bool

	if strings.Contains(req.Node.Name, c.infraShellClient.GeneratedInventory.Cluster.Name) {
		nodeGroupName, ok = c.infraShellClient.NodesNodeGroupByName[req.Node.Name]
	} else {
		nodeGroupName, ok = c.infraShellClient.NodesNodeGroupByID[req.Node.Name]
	}

	if nodeGroupName == "" || !ok {
		klog.V(1).Infof("could not find nodegroup for node:  %s", req.Node.Name)
		return nil, fmt.Errorf("could not find nodegroup for node:  %s", req.Node.Name)
	}

	return &eg.NodeGroupForNodeResponse{
		NodeGroup: &eg.NodeGroup{
			Id:      nodeGroupName,
			MinSize: int32(c.infraShellClient.InfraNodeGroups[nodeGroupName].MinSize),
			MaxSize: int32(c.infraShellClient.InfraNodeGroups[nodeGroupName].MaxSize),
		},
	}, nil

}

// PricingNodePrice returns a theoretical minimum price of running a node for
// a given period of time on a perfectly matching machine.
// Implementation optional.
// func (c *RSInfraScaler) PricingNodePrice(context.Context, *eg.PricingNodePriceRequest) (*eg.PricingNodePriceResponse, error) {
// 	return &eg.PricingNodePriceResponse{}, nil
// }

// PricingPodPrice returns a theoretical minimum price of running a pod for a given
// period of time on a perfectly matching machine.
// Implementation optional.
// func (c *RSInfraScaler) PricingPodPrice(context.Context, *eg.PricingPodPriceRequest) (*eg.PricingPodPriceResponse, error) {
// 	return &eg.PricingPodPriceResponse{}, nil
// }

// GPULabel returns the label added to nodes with GPU resource.
func (c *RSInfraScaler) GPULabel(context.Context, *eg.GPULabelRequest) (*eg.GPULabelResponse, error) {
	klog.V(5).Infof("GPULabel gRPC function called")
	return &eg.GPULabelResponse{Label: "nodeHasGPU"}, nil
}

// GetAvailableGPUTypes return all available GPU types cloud provider supports.
func (c *RSInfraScaler) GetAvailableGPUTypes(context.Context, *eg.GetAvailableGPUTypesRequest) (*eg.GetAvailableGPUTypesResponse, error) {
	klog.V(5).Infof("GetAvailableGPUTypes gRPC function called")
	return &eg.GetAvailableGPUTypesResponse{
		GpuTypes: nil,
	}, nil
}

// Cleanup cleans up open resources before the cloud provider is destroyed, i.e. go routines etc.
func (c *RSInfraScaler) Cleanup(context.Context, *eg.CleanupRequest) (*eg.CleanupResponse, error) {
	klog.V(5).Infof("Cleanup gRPC function called")
	return &eg.CleanupResponse{}, nil
}

// Refresh is called before every main loop and can be used to dynamically update cloud provider state.
func (c *RSInfraScaler) Refresh(context.Context, *eg.RefreshRequest) (*eg.RefreshResponse, error) {
	klog.V(5).Infof("Refresh gRPC function called")

	klog.V(1).Infof("Syncing local nodegroups info with remote safescaled data")
	go c.infraShellClient.updateNodesNodeGroup()

	return &eg.RefreshResponse{}, nil
}

// NodeGroupTargetSize returns the current target size of the node group. It is possible
// that the number of nodes in Kubernetes is different at the moment but should be equal
// to the size of a node group once everything stabilizes (new nodes finish startup and
// registration or removed nodes are deleted completely).
func (c *RSInfraScaler) NodeGroupTargetSize(ctx context.Context, req *eg.NodeGroupTargetSizeRequest) (*eg.NodeGroupTargetSizeResponse, error) {
	klog.V(5).Infof("NodeGroupTargetSize gRPC function called")

	if infraNodeGroup, ok := c.infraShellClient.InfraNodeGroups[req.Id]; ok {
		return &eg.NodeGroupTargetSizeResponse{
			TargetSize: int32(infraNodeGroup.TargetSize),
		}, nil
	} else {
		return nil, fmt.Errorf("could not find node group info for node: %s", req.Id)
	}
}

// NodeGroupIncreaseSize increases the size of the node group. To delete a node you need
// to explicitly name it and use NodeGroupDeleteNodes. This function should wait until
// node group size is updated.
func (c *RSInfraScaler) NodeGroupIncreaseSize(_ context.Context, req *eg.NodeGroupIncreaseSizeRequest) (*eg.NodeGroupIncreaseSizeResponse, error) {
	klog.V(5).Infof("NodeGroupIncreaseSize gRPC function called")

	if req.Delta < 0 {
		klog.V(1).Infof("delta cannot be negative on nodegroup increase")
		return &eg.NodeGroupIncreaseSizeResponse{}, nil
	}

	klog.V(1).Infof("Starting increase of nodegroup %s of %d nodes", req.Id, req.Delta)
	go c.infraShellClient.IncreaseNodeGroupSize(req.Id, int(req.Delta))

	return &eg.NodeGroupIncreaseSizeResponse{}, nil
}

// NodeGroupDeleteNodes deletes nodes from this node group (and also decreasing the size
// of the node group with that). Error is returned either on failure or if the given node
// doesn't belong to this node group. This function should wait until node group size is updated.

func (c *RSInfraScaler) NodeGroupDeleteNodes(_ context.Context, req *eg.NodeGroupDeleteNodesRequest) (*eg.NodeGroupDeleteNodesResponse, error) {
	klog.V(5).Infof("NodeGroupDeleteNodes gRPC function called")

	nodes_to_delete := make([]string, 0)
	for _, node := range req.Nodes {
		nodes_to_delete = append(nodes_to_delete, node.Name)
	}

	klog.V(1).Infof("Deleting nodes %v of %s nodegroup", nodes_to_delete, req.Id)
	go c.infraShellClient.RemoveNodesFromNodeGroup(req.Id, nodes_to_delete)

	return &eg.NodeGroupDeleteNodesResponse{}, nil
}

// NodeGroupDecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the request
// for new nodes that have not been yet fulfilled. Delta should be negative. It is assumed
// that cloud provider will not delete the existing nodes if the size when there is an option
// to just decrease the target.
func (c *RSInfraScaler) NodeGroupDecreaseTargetSize(context.Context, *eg.NodeGroupDecreaseTargetSizeRequest) (*eg.NodeGroupDecreaseTargetSizeResponse, error) {
	klog.V(5).Infof("NodeGroupDecreaseTargetSize gRPC function called")
	return &eg.NodeGroupDecreaseTargetSizeResponse{}, nil
}

// NodeGroupNodes returns a list of all nodes that belong to this node group.
func (c *RSInfraScaler) NodeGroupNodes(ctx context.Context, req *eg.NodeGroupNodesRequest) (*eg.NodeGroupNodesResponse, error) {
	klog.V(5).Infof("NodeGroupNodes gRPC function called")

	var instances []*eg.Instance

	for _, nodeID := range c.infraShellClient.InfraNodeGroups[req.Id].NodesIDs {
		instances = append(instances, &eg.Instance{
			Id: nodeID,
			Status: &eg.InstanceStatus{
				// TODO: get real value
				InstanceState: eg.InstanceStatus_instanceRunning,
			},
		})
	}
	return &eg.NodeGroupNodesResponse{
		Instances: instances,
	}, nil
}

// NodeGroupTemplateNodeInfo returns a structure of an empty (as if just started) node,
// with all of the labels, capacity and allocatable information. This will be used in
// scale-up simulations to predict what would a new node look like if a node group was expanded.
// Implementation optional.
func (c *RSInfraScaler) NodeGroupTemplateNodeInfo(ctx context.Context, req *eg.NodeGroupTemplateNodeInfoRequest) (*eg.NodeGroupTemplateNodeInfoResponse, error) {
	klog.V(5).Infof("NodeGroupTemplateNodeInfo gRPC function called")

	return &eg.NodeGroupTemplateNodeInfoResponse{
		NodeInfo: &c.infraShellClient.InfraNodeGroups[req.Id].NodeTemplateInfo,
	}, nil
}

// GetOptions returns NodeGroupAutoscalingOptions that should be used for this particular
// NodeGroup. Returning a grpc error will result in using default options.
// Implementation optional.
func (c *RSInfraScaler) NodeGroupGetOptions(context.Context, *eg.NodeGroupAutoscalingOptionsRequest) (*eg.NodeGroupAutoscalingOptionsResponse, error) {
	klog.V(5).Infof("NodeGroupGetOptions gRPC function called")
	return &eg.NodeGroupAutoscalingOptionsResponse{
		NodeGroupAutoscalingOptions: &eg.NodeGroupAutoscalingOptions{
			ScaleDownUtilizationThreshold:    0.5,
			ScaleDownUnneededTime:            &metav1.Duration{Duration: time.Minute * 10},
			ScaleDownUnreadyTime:             &metav1.Duration{Duration: time.Minute * 10},
			ScaleDownGpuUtilizationThreshold: 0.5,
		},
	}, nil
}
