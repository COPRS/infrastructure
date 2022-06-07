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

func (s *SafescaleShellClient) GetNodeGroups() ([]string, error) {
	cmd := exec.Command("safescale", "tag", "list")
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	cmd.Wait()
	klog.V(5).Infof("Tag list response: %v", string(out))
	if err != nil {
		klog.V(1).Infof("could not run safescale tag list: %v", err)
		return nil, err
	}

	status := gjson.GetBytes(out, "status")
	if status.String() != "success" {
		klog.V(1).Infof("could not get tags: %v", err)
		return nil, err
	}

	tags := gjson.GetBytes(out, "result.#.name").Array()
	klog.V(5).Infof("Parsed tags: %v", tags)

	var nodeGroups []string
	for _, tag := range tags {
		nodeGroups = append(nodeGroups, tag.Str)
	}
	return nodeGroups, nil
}

func (s *SafescaleShellClient) GetNodeGroupNodesIDs(nodeGroup string) ([]string, error) {
	cmd := exec.Command("safescale", "tag", "inspect", nodeGroup)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	cmd.Wait()
	klog.V(5).Infof("Tag inspect response: %v", string(out))
	if err != nil {
		klog.V(1).Infof("could not run safescale tag inspect: %v", err)
		return nil, err
	}

	status := gjson.GetBytes(out, "status")
	if status.String() != "success" {
		klog.V(1).Infof("could not inspect tag: %v", err)
		return nil, err
	}

	tagNodes := gjson.GetBytes(out, "result.hosts.#.id").Array()
	klog.V(5).Infof("Parsed nodes ids: %v", tagNodes)

	ngn := make([]string, 0)

	for _, node := range tagNodes {
		ngn = append(ngn, node.Str)
	}
	return ngn, nil
}

func (s *SafescaleShellClient) GetNodeGroupNodesNames(nodeGroup string) ([]string, error) {
	cmd := exec.Command("safescale", "tag", "inspect", nodeGroup)
	cmd.Env = os.Environ()
	out, err := cmd.CombinedOutput()
	cmd.Wait()
	klog.V(5).Infof("Tag inspect response: %v", string(out))
	if err != nil {
		klog.V(1).Infof("could not run safescale tag inspect: %v", err)
		return nil, err
	}

	status := gjson.GetBytes(out, "status")
	if status.String() != "success" {
		klog.V(1).Infof("could not inspect tag: %v", err)
		return nil, err
	}

	tagNodes := gjson.GetBytes(out, "result.hosts.#.name").Array()
	klog.V(5).Infof("Parsed nodes names: %v", tagNodes)

	ngn := make([]string, 0)

	for _, node := range tagNodes {
		ngn = append(ngn, node.Str)
	}
	return ngn, nil
}
