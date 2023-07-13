# Deploy Reference System platform on a custom infrastructure

## What the cluster playbook does with SafeScale

Reference System uses SafeScale to deploy its infrastructure on any compatible cloud provider. See [SafeScale](https://github.com/CS-SI/SafeScale).

The `cluster.yaml` playbook, after having created the cluster and volumes with safescale, will populate the ansible `{{ inventory_dir }}/hosts.yaml` file according to the node groups configurations, with the machines IPs and specific hosts variables.

Here is an example of a generated `hosts.yaml` file:

```yaml
all:
  hosts:
    # setup and localhost should not need to be edited
    setup:
      ansible_connection: local
      ansible_host: 127.0.0.1
    localhost:
      ansible_connection: local
      ansible_host: 127.0.0.1
# BEGIN ANSIBLE MANAGED BLOCK
    rs-demo5-node-2:
      # the IP can be a private one if you ssh to the machine through the gateway like safescale does (see the ssh-main-gateway.conf file)
      ansible_host: 192.168.3.6
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-2.pem"
      # provider_id is only used by the autoscaling components, that are not relevant if not using safescale
      provider_id: 91c54f1d-a672-4f20-8c5c-fa0b4802f0b6
    rs-demo5-node-3:
      ansible_host: 192.168.2.212
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-3.pem"
      provider_id: 36ae0efb-3158-48ad-b1ac-fcffdeb4cb77
    rs-demo5-node-5:
      ansible_host: 192.168.0.106
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-5.pem"
      provider_id: 61c1bc4c-b0ee-4e04-922f-5bd117742a8a
    rs-demo5-node-4:
      ansible_host: 192.168.2.184
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-4.pem"
      provider_id: 861775d4-525e-4117-a9ad-454bed24386f
    rs-demo5-node-1:
      ansible_host: 192.168.0.155
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-1.pem"
      provider_id: dcf54a16-d50d-48ce-9565-8d3db65246b8
    rs-demo5-node-6:
      ansible_host: 192.168.3.194
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-6.pem"
      provider_id: 65e8d3d7-07f7-44b7-9269-7c75073c7aa3
    rs-demo5-node-9:
      ansible_host: 192.168.3.62
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-9.pem"
      provider_id: 50a45f45-e5cc-4f62-9c9c-44e4e8987679
    rs-demo5-node-10:
      ansible_host: 192.168.0.113
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-10.pem"
      provider_id: 7254f51c-b9c5-4f59-8fe3-82fb2d0d4763
    rs-demo5-node-13:
      ansible_host: 192.168.1.178
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-node-13.pem"
      provider_id: 4c925ad0-2ded-43c5-882e-ceb599de003d
    rs-demo5-master-1:
      ansible_host: 192.168.0.171
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/rs-demo5-master-1.pem"
      provider_id: 183b921d-8f56-416f-83db-8b5158d41a4d
    gw-rs-demo5:
      ansible_host: 192.168.2.74
      ansible_ssh_private_key_file: "INVENTORY_DIR/artifacts/.ssh/gw-rs-demo5.pem"
      provider_id: b053b970-a945-45f3-b973-711be2b57389

gateway:
  hosts:
    gw-rs-demo5:
  vars:
    nodegroup: gateway
    kubelet_config_extra_args:
      systemReserved:
        cpu: '1'
        memory: 2Gi
    node_labels:
      node-role.kubernetes.io/gateway: ''
    node_taints:
    - node-role.kubernetes.io/gateway:NoSchedule

infra:
  hosts:
    rs-demo5-node-2:
    rs-demo5-node-5:
    rs-demo5-node-6:
    rs-demo5-node-4:
    rs-demo5-node-3:
    rs-demo5-node-1:
  vars:
    nodegroup: infra
    node_labels:
      node-role.kubernetes.io/infra: ''

kube_control_plane:
  hosts:
    rs-demo5-master-1:
  vars:
    nodegroup: kube_control_plane

# For each additional node group, write the below structure
rook_ceph:
  hosts:
    rs-demo5-node-13:
    rs-demo5-node-9:
    rs-demo5-node-10:
  vars:
    nodegroup: rook_ceph
    node_labels:
      node-role.kubernetes.io/rook_ceph: ''

kube_node:
  children:
    gateway:
    infra:
    rook_ceph:
    # Add here any additional node group

# No need to edit
etcd:
  children:
    kube_control_plane:

# No need to edit
k8s_cluster:
  children:
    kube_control_plane:
    kube_node:
# END ANSIBLE MANAGED BLOCK
```

## Build your custom inventory

See the [ansible documentation](https://docs.ansible.com/ansible/latest/getting_started/get_started_inventory.html#get-started-inventory) to understand precisely the hosts file and build your own, similar to the one above that is created by the the `cluster.yaml` playbook with the information given by SafeScale.

Configure the `ansible.cfg` according to your machines ssh configuration. (user, password, ssh private key, sudo permissions...)

> To have a functionnal filesystem on the platform, the **rook-ceph** nodes must a an additional disk with no partitions on it

## Follow the main deployment procedure

Once you have a working `hosts.yaml` file, follow the main deployment steps in the [README](../../../README.md) file, skip step **6.** and at step **7.** do not create the cluster, configure only the machine by running:

```shellsession
ansible-playbook cluster.yaml \
    -i inventory/mycluster/hosts.yaml \
    -t config
```
