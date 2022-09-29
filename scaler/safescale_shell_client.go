package main

import (
	"os"
	"os/exec"

	"github.com/tidwall/gjson"
	"k8s.io/klog/v2"
)

type SafescaleShellClient struct {
	tenant string
}

func NewSafescaleShellClient(tenant string) *SafescaleShellClient {
	ssc := SafescaleShellClient{
		tenant: tenant,
	}
	ssc.setTenant(tenant)
	return &ssc
}

func (s *SafescaleShellClient) setTenant(tenant string) error {
	klog.V(5).Infof("Setting safescale tenant to %s", tenant)
	safescaleCmd := exec.Command("safescale", "tenant", "set", tenant)
	safescaleCmd.Env = os.Environ()
	_, err := safescaleCmd.CombinedOutput()
	safescaleCmd.Wait()
	if err != nil {
		klog.Fatalf("could not set tenant %s on safescaled: %v", tenant, err)
	}
	klog.V(3).Infof("Succesfully set safescale tenant to %s", tenant)
	return nil
}

func (s *SafescaleShellClient) GetNodeGroupNodesIDs(clusterName string, nodeGroup string) ([]string, error) {
	cmd := exec.Command("safescale", "label", "inspect", clusterName+"-nodegroup")
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	cmd.Wait()
	klog.V(5).Infof("Label inspect response: %v", string(out))
	if err != nil {
		klog.V(1).Infof("could not run safescale label inspect: %v", err)
		return nil, err
	}

	status := gjson.GetBytes(out, "status")
	if status.String() != "success" {
		klog.V(1).Infof("could not inspect label: %v", err)
		return nil, err
	}

	nodeGroupNodes := gjson.GetBytes(out, "result.hosts.#(value==\""+nodeGroup+"\")#.id").Array()
	klog.V(5).Infof("Parsed nodes ids: %v", nodeGroupNodes)

	ngn := make([]string, 0)

	for _, node := range nodeGroupNodes {
		ngn = append(ngn, node.Str)
	}
	return ngn, nil
}

func (s *SafescaleShellClient) GetNodeGroupNodesNames(clusterName string, nodeGroup string) ([]string, error) {
	cmd := exec.Command("safescale", "label", "inspect", clusterName+"-nodegroup")
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	cmd.Wait()
	klog.V(5).Infof("Label inspect response: %v", string(out))
	if err != nil {
		klog.V(1).Infof("could not run safescale label inspect: %v", err)
		return nil, err
	}

	status := gjson.GetBytes(out, "status")
	if status.String() != "success" {
		klog.V(1).Infof("could not inspect label: %v", err)
		return nil, err
	}

	nodeGroupNodes := gjson.GetBytes(out, "result.hosts.#(value==\""+nodeGroup+"\")#.name").Array()
	klog.V(5).Infof("Parsed nodes names: %v", nodeGroupNodes)

	ngn := make([]string, 0)

	for _, node := range nodeGroupNodes {
		ngn = append(ngn, node.Str)
	}
	return ngn, nil
}
