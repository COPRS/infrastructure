// Copyright 2023 CS Group
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

type NodeInfo struct {
	Name string
	Id   string
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

func (s *SafescaleShellClient) GetNodeGroupsNodes(clusterName string) (map[string][]NodeInfo, error) {
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

	nodeGroupsNodes := make(map[string][]NodeInfo, 0)

	hosts := gjson.GetBytes(out, "result.hosts").Array()

	for _, host := range hosts {
		nodeGroup := host.Get("value").Str
		id := host.Get("id").Str
		name := host.Get("name").Str
		nodeInfo := NodeInfo{
			Id:   id,
			Name: name,
		}

		nodeGroupsNodes[nodeGroup] = append(nodeGroupsNodes[nodeGroup], nodeInfo)

	}

	klog.V(5).Infof("Parsed node groups: %v", nodeGroupsNodes)

	return nodeGroupsNodes, nil
}
