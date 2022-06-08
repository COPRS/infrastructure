# Configure the `safescale.yaml` inventory file

## Cluster wide parameters

 - name : name of the cluster
 - complexity : SafeScale cluster complexity, see below for details
 - cidr : CIDR of the cluster subnet, you should not need to change it
 - os : OS image to use for the machines in the cluster, should reference the one previously created with the `image.yaml` playbook

## Node group specification

For each node groups, the following fields are availables:
 - name : name of the node group, in valid ansible variable format (no "-" for example, see [here](https://docs.ansible.com/ansible/latest/user_guide/playbooks_variables.html#creating-valid-variable-names))
 - sizing : virtual machine sizing, in the SafeScale format, see [here](https://github.com/CS-SI/SafeScale/blob/master/doc/USAGE.md#safescale_sizing)
 - min_size : minimum amount of nodes in the node group, will be used at cluster creation
 - max_size : maximum amount of nodes in the node group
 - volume : definition of the volume to be added to each node in the node group
   - type : type of volume (SSD, HDD or COLD)
   - size : size in Go
 - kubespray : kubespray specific variables for this node group, find the documentation [here](https://github.com/kubernetes-sigs/kubespray/blob/master/docs/vars.md) and some examples below

This specification will be used by the cluster-autoscaler to run its scale up simulations.

## Mandatory node groups specifics

The node groups `kube_control_plane`, `gateway` and `infra` are mandatory and cannot have volumes.

The amount of nodes in the `kube_control_plane` and `gateway` node groups are tied the SafeScale cluster complexity, meaning:
 - small: 1 gateway, 1 kube_control_plane, 1 infra
 - normal: 2 gateways, 2 kube_control_plane, 3 infra
 - large: 2 gateways, 3 kube_control_plane, 7 infra

The minimum size if the below defined `infra` node group must be superior to the size of the `infra` node group defined by the SafeScale complexity above.

Here is an example of a valid basic cluster configuration:
```yaml
cluster:
  name: rs-dev
  complexity: small
  cidr: 192.168.0.0/16
  os: "csc-rs-ubuntu"
  nodegroups:
    - name: kube_control_plane
      sizing: "cpu=4,ram=[8-10],disk=80"
    - name: gateway
      sizing: "cpu=2,ram=[4-5],disk=20"
      kubespray:
        node_labels:
          node-role.kubernetes.io/gateway: ''
        node_taints:
          - node-role.kubernetes.io/gateway:NoSchedule
        kubelet_config_extra_args:
          systemReserved:
            cpu: "1"
            memory: "2Gi"
    - name: infra
      min_size: 2
      max_size: 5
      sizing: "cpu=8,ram=[14-18],disk=40"
      kubespray:
        node_labels: 
          node-role.kubernetes.io/infra: ''
```

## S3 buckets

SafeScale will create for you the S3 buckets specified in the `buckets` array, this example creates the four buckets that are used by default in Reference System:
```yaml
buckets:
  - "{{ cluster.name }}-elasticsearch-processing"
  - "{{ cluster.name }}-elasticsearch-security"
  - "{{ cluster.name }}-thanos"
  - "{{ cluster.name }}-loki"
```