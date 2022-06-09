package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
	"k8s.io/utils/strings/slices"
)

var safescaleComplexity = map[string]map[string]int{
	"small": {
		"gateway":            1,
		"kube_control_plane": 1,
	},
	"normal": {
		"gateway":            2,
		"kube_control_plane": 2,
	},
	"large": {
		"gateway":            2,
		"kube_control_plane": 3,
	},
}

type InfraShellClient struct {
	inventoryDirPath     string
	infrastructurePath   string
	SafescaleShellClient *SafescaleShellClient
	InfraNodeGroups      map[string]*InfraNodeGroup
	NodesNodeGroupByName map[string]string
	NodesNodeGroupByID   map[string]string
	GeneratedInventory   *GeneratedInventory
}

type InfraNodeGroup struct {
	Name             string
	ActualSize       int
	TargetSize       int
	MinSize          int
	MaxSize          int
	NodesNames       []string
	NodesIDs         []string
	NodeTemplateInfo v1.Node
}

type GeneratedInventory struct {
	Cluster struct {
		Name       string `yaml:"name"`
		Complexity string `yaml:"complexity"`
		Cidr       string `yaml:"cidr"`
		Os         string `yaml:"os"`
		Nodegroups []struct {
			Name    string `yaml:"name"`
			Count   int    `yaml:"count,omitempty"`
			MinSize int    `yaml:"min_size,omitempty"`
			MaxSize int    `yaml:"max_size,omitempty"`
			Volume  struct {
				VolumeType string `yaml:"type"`
				Size       string `yaml:"size"`
			} `yaml:"volume,omitempty"`
			Sizing    string `yaml:"sizing,omitempty"`
			Kubespray struct {
				NodeLabels map[string]string `yaml:"node_labels,omitempty"`
				NodeTaints []string          `yaml:"node_taints,omitempty"`
			} `yaml:"kubespray,omitempty"`
		} `yaml:"nodegroups"`
	} `yaml:"cluster"`
}

func NewInfraShellClient(inventoryDirPath string, infrastructurePath string, sf *SafescaleShellClient) *InfraShellClient {
	isc := &InfraShellClient{
		inventoryDirPath:     inventoryDirPath,
		infrastructurePath:   infrastructurePath,
		SafescaleShellClient: sf,
		InfraNodeGroups:      make(map[string]*InfraNodeGroup, 0),
		NodesNodeGroupByName: make(map[string]string, 0),
		GeneratedInventory:   &GeneratedInventory{},
	}

	klog.V(5).Infof("Check ansible-playbook binary")
	ansibleCmd := exec.Command("ansible-playbook", "--version")
	if err := ansibleCmd.Run(); err != nil {
		klog.Fatalf("ansible-playbook not found: %v", err)
	}

	klog.V(5).Infof("Read generated_inventory_vars.yaml file")
	invFile, err := os.ReadFile(inventoryDirPath + "/group_vars/all/generated_inventory_vars.yaml")
	if err != nil {
		klog.Fatalf("could not read inventory file: %v", err)
	}

	if err = yaml.Unmarshal(invFile, isc.GeneratedInventory); err != nil {
		klog.Fatalf("error: %v", err)
	}

	if isc.GeneratedInventory.Cluster.Name == "" {
		klog.Fatalf("error: cluster name not found in inventory")
	}
	isc.readNodeGroupsFromInventory(isc.GeneratedInventory)

	klog.V(3).Infof("Syncing local nodegroups info with remote safescaled data")
	isc.updateNodesNodeGroup()

	klog.V(3).Infof("Running first update_hosts infra playbook")
	err = isc.runPlaybook("cluster.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-t", "update_hosts",
		"-e", "use_private_gateway_ip=true",
		"-e", "safescale_path=safescale",
	)
	if err != nil {
		klog.Fatalf("could not update hosts at startup: %v", err)
	}

	return isc
}

func (i *InfraShellClient) IncreaseNodeGroupSize(nodeGroup string, count int) (err error) {
	i.InfraNodeGroups[nodeGroup].TargetSize += count

	klog.V(3).Infof("Running safescale infra expand playbook")
	err = i.runPlaybook("cluster.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-t", "expand,update_hosts",
		"-e", "use_private_gateway_ip=true",
		"-e", "expand_nodegroup_name="+nodeGroup,
		"-e", "expand_count="+strconv.Itoa(count),
		"-e", "safescale_path=safescale",
	)
	if err != nil {
		klog.V(1).Infof("could not expand safescale cluster: %v", err)
		return err
	}

	// klog.V(3).Infof("Running kubespray facts playbook")
	// err = i.runPlaybook("collections/kubespray/facts.yml",
	// 	"-i", inventory_path+"/hosts.yaml",
	// )
	// if err != nil {
	// 	klog.V(1).Infof("could not run kubespray facts playbook: %v", err)
	// 	return err
	// }

	klog.V(3).Infof("Getting last created node from safescaled")
	ngn, err := i.SafescaleShellClient.GetNodeGroupNodesNames(nodeGroup)
	if err != nil {
		klog.V(5).Infof("could not get nodes for nodegroup: %v", err)
		return err
	}
	sort.Strings(ngn)
	createdNodes := ngn[len(ngn)-count:]
	klog.V(5).Infof("Created %v", createdNodes)

	klog.V(3).Infof("Running cluster config-machines playbook")
	err = i.runPlaybook("cluster.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-t", "config",
		"--limit="+strings.Join(createdNodes, ","),
	)
	if err != nil {
		klog.V(1).Infof("could not configure newly created machines: %v", err)
		return err
	}

	klog.V(3).Infof("Running security playbook")
	err = i.runPlaybook("security.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-b",
		"--limit="+strings.Join(createdNodes, ","),
	)
	if err != nil {
		klog.V(1).Infof("could not run security playbook: %v", err)
		return err
	}

	klog.V(3).Infof("Running kubespray scale playbook")
	err = i.runPlaybook("collections/kubespray/cluster.yml",
		"-i", inventory_path+"/hosts.yaml",
		"-b",
		"--limit="+strings.Join(createdNodes, ","),
	)
	if err != nil {
		klog.V(1).Infof("could not run kubespray scale playbook: %v", err)
		return err
	}

	klog.V(3).Infof("Running providerids playbook")
	err = i.runPlaybook("cluster.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-t", "providerids",
		"-e", "safescale_path=safescale",
	)
	if err != nil {
		klog.V(1).Infof("could not run provider ids update: %v", err)
		return err
	}

	klog.V(1).Infof("Finished scaling up %s by %d node(s)", nodeGroup, count)
	i.InfraNodeGroups[nodeGroup].ActualSize += count
	return nil
}

func (i *InfraShellClient) RemoveNodesFromNodeGroup(nodeGroup string, nodes []string) (err error) {
	i.InfraNodeGroups[nodeGroup].TargetSize -= len(nodes)

	if len(nodes) < 1 {
		return fmt.Errorf("nodes to delete cannot be empty")
	}

	for _, node := range nodes {
		klog.V(3).Infof("Running kubespray remove-node playbook for %s", node)
		err = i.runPlaybook("collections/kubespray/remove-node.yml",
			"-i", inventory_path+"/hosts.yaml",
			"-e", "skip_confirmation=yes",
			"-e", "reset_nodes=false",
			"-e", "allow_ungraceful_removal=true",
			"-e", "node="+node,
			"-b",
		)
		if err != nil {
			klog.V(1).Infof("could not run kubespray remove-node playbook for %s: %v", node, err)
			// do not return if failed at this point
		}
	}

	klog.V(3).Infof("Running infra delete nodes playbook")
	err = i.runPlaybook("cluster.yaml",
		"-i", inventory_path+"/hosts.yaml",
		"-t", "shrink,update_hosts",
		"-e", "use_private_gateway_ip=true",
		"-e", "nodes_to_delete="+strings.Join(nodes, ","),
		"-e", "safescale_path=safescale",
	)
	if err != nil {
		klog.V(1).Infof("could not run delete nodes playbook: %v", err)
		return err
	}

	klog.V(1).Infof("Finished deleting %v ", nodes)
	i.InfraNodeGroups[nodeGroup].ActualSize -= len(nodes)
	return nil
}

func (i *InfraShellClient) readNodeGroupsFromInventory(inv *GeneratedInventory) error {
	for _, nodegroup := range inv.Cluster.Nodegroups {

		cpu, ram, err := ParseSafeScaleSizing(nodegroup.Sizing)
		if err != nil {
			klog.V(1).Infof("could not parse sizing: %v", err)
			return err
		}
		apods := resource.Quantity{}
		apods.Add(resource.MustParse("110"))

		taints := make([]v1.Taint, 0)
		for _, taint := range nodegroup.Kubespray.NodeTaints {
			key_value := strings.Split(taint, ":")

			taints = append(taints, v1.Taint{
				Key:    key_value[0],
				Effect: v1.TaintEffect(key_value[1]),
			})
		}

		ngni, err := i.SafescaleShellClient.GetNodeGroupNodesIDs(nodegroup.Name)
		if err != nil {
			klog.V(1).Infof("could not get nodes ids for nodegroup: %v", err)
			return err
		}
		ngnn, err := i.SafescaleShellClient.GetNodeGroupNodesNames(nodegroup.Name)
		if err != nil {
			klog.V(1).Infof("could not get nodes names for nodegroup: %v", err)
			return err
		}

		minSize := nodegroup.MinSize
		maxSize := nodegroup.MaxSize
		if slices.Contains([]string{"kube_control_plane", "gateway"}, nodegroup.Name) {
			minSize = safescaleComplexity[strings.ToLower(i.GeneratedInventory.Cluster.Complexity)][nodegroup.Name]
			maxSize = minSize
		}

		i.InfraNodeGroups[nodegroup.Name] = &InfraNodeGroup{
			Name:       nodegroup.Name,
			ActualSize: len(ngni),
			TargetSize: len(ngni),
			MinSize:    minSize,
			MaxSize:    maxSize,
			NodesIDs:   ngni,
			NodesNames: ngnn,
			NodeTemplateInfo: v1.Node{
				Status: v1.NodeStatus{
					Capacity: v1.ResourceList{
						v1.ResourcePods:   apods.DeepCopy(),
						v1.ResourceCPU:    cpu.DeepCopy(),
						v1.ResourceMemory: ram.DeepCopy(),
					},
					Allocatable: v1.ResourceList{
						v1.ResourcePods:   apods.DeepCopy(),
						v1.ResourceCPU:    cpu.DeepCopy(),
						v1.ResourceMemory: ram.DeepCopy(),
					},
				},
				Spec: v1.NodeSpec{
					Taints: taints,
				},
				ObjectMeta: metav1.ObjectMeta{
					Labels: nodegroup.Kubespray.NodeLabels,
				},
			},
		}
		klog.V(5).Infof("infraNodeGroup %s: %v:", nodegroup.Name, i.InfraNodeGroups[nodegroup.Name])
	}
	return nil
}

func (i *InfraShellClient) updateNodesNodeGroup() error {
	nngbn := make(map[string]string, 0)
	nngbi := make(map[string]string, 0)

	for nodeGroupName := range i.InfraNodeGroups {
		klog.V(5).Infof("infraNodeGroup %s", nodeGroupName)
		ngnn, err := i.SafescaleShellClient.GetNodeGroupNodesNames(nodeGroupName)
		if err != nil {
			klog.V(1).Infof("could not get nodes names for nodegroup: %v", err)
			return err
		}
		ngni, err := i.SafescaleShellClient.GetNodeGroupNodesIDs(nodeGroupName)
		if err != nil {
			klog.V(1).Infof("could not get nodes ids for nodegroup: %v", err)
			return err
		}
		i.InfraNodeGroups[nodeGroupName].NodesIDs = ngni
		i.InfraNodeGroups[nodeGroupName].NodesNames = ngnn
		for _, nodeName := range ngnn {
			nngbn[nodeName] = nodeGroupName
		}
		for _, nodeID := range ngni {
			nngbn[nodeID] = nodeGroupName
		}
	}

	i.NodesNodeGroupByName = nngbn
	i.NodesNodeGroupByID = nngbi
	return nil
}

func (i *InfraShellClient) runPlaybook(args ...string) error {
	cmd := exec.Command("ansible-playbook", args...)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Env = os.Environ()
	cmd.Dir = i.infrastructurePath
	klog.V(3).Infof("running: %v", cmd.Args)
	err := cmd.Run()
	cmd.Wait()
	klog.V(5).Infof("playbook output: %v", stdout.String())
	if err != nil {
		klog.V(1).Infof("could not run playbook with args %v: %v", cmd.Args, err)
		klog.V(3).Infof("stderr: %s", stderr.String())
		return err
	}
	klog.V(3).Infof("finished with success running: %v", cmd.Args)
	return nil
}

func ParseSafeScaleSizing(sizing string) (*resource.Quantity, *resource.Quantity, error) {
	sizing = strings.TrimSpace(sizing)

	var cpu, memory string
	var err error

	for _, resource := range strings.Split(sizing, ",") {
		resource = strings.TrimSpace(resource)
		if strings.HasPrefix(resource, "cpu") {
			cpu, err = ParseSafeScaleSizingValue(resource[4:])
			if err != nil {
				return nil, nil, err
			}
		} else if strings.HasPrefix(resource, "ram") {
			memory, err = ParseSafeScaleSizingValue(resource[4:])
			if err != nil {
				return nil, nil, err
			}
		}
	}

	parsedCpu := resource.MustParse(cpu)
	parsedRam := resource.MustParse(memory + "Gi")
	return &parsedCpu, &parsedRam, nil
}

func ParseSafeScaleSizingValue(sizing string) (string, error) {
	sizing = strings.TrimSpace(sizing)

	if len(sizing) == 0 {
		return "", fmt.Errorf("could not parse empty sizing")
	}
	if len(sizing) == 1 {
		return sizing, nil
	}
	r := regexp.MustCompile(`\[(\d+)-(\d+)\]`)
	res := r.FindStringSubmatch(sizing)

	if len(res) == 0 {
		return "", fmt.Errorf("could not parse sizing: %v", sizing)
	}
	min, err := strconv.Atoi(res[1])
	if err != nil {
		return "", err
	}
	max, err := strconv.Atoi(res[2])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", int((max+min)/2)), nil
}
