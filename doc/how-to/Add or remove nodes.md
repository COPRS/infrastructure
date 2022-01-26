# Add a node in the cluster

> LIMITATION: SafeScale does not distinguish multiple categories of new nodes. All added ones will be named  CLUSTER_NAME-node-NODE_NUMBER, whatever their future role in the cluster (master, gateway, etc.).

> LIMITATION: the ansible *hosts* file automatically generated does not support adding new nodes on the cluster.  
> Modifying ```hosts.ini``` with cluster-setup.yaml may override your changes.  
> You will have to ensure nodes are in the right groups.

First, expand the SafeScale cluster.

```Bash
safescale cluster expand \
  --os IMAGE_NAME \
  --count SUPPLEMENTARY_NODE_COUNT \
  --node-sizing cpu=CPU_COUNT,ram=MEMORY_AMOUNT,disk=STORAGE_MOUNT
```

*Do not forget to update the `WhitelistTemplateRegexp`/`BlacklistTemplateRegexp` in your SafeScale tenants file to reflect the template of the node you want to add.  
After modifying the _tenants_ file, you need to restart safescaled.  
Read [safescale documenation](https://github.com/CS-SI/SafeScale/blob/d8b98cb28c29cbbd87162b33e3a84f159a6707d9/doc/SCANNER.md#safescale-scanner) for more information related to managing templates.*

Secondly, manually update `hosts.ini` in your inventory and add the newly created nodes into their corresponding groups.

If the kind (group) of your node does not exist in the file, refer to the [new node kinds (processing, etc.)](#new_kinds) section.

If `hosts.ini`  is empty, you can update it using the following command:

```Bash
ansible-playbook cluster-setup.yaml -t hosts_update -i inventory/mycluster/hosts.ini
```

Once you set your *hosts* file, you need to launch Kubespray to integrate the newly added nodes into the Kubernetes cluster.

## <a name="worker_nodes"></a>Integrate a worker node into k8s

If your new node is not a control plane (k8s master), it is a worker node (prometheus, egress, processing, etc.).

```Bash
# Scale the cluster with all the new nodes.
ansible-playbook collections/kubespray/scale.yml -b -i inventory/mycluster/hosts.ini 

# You can add only a specific node
# First, you need to collect facts on all the nodes
ansible-playbook collections/kubespray/facts.yml -i inventory/mycluster/hosts.ini
# Then, run scale.yml only on the node you want to add
ansible-playbook collections/kubespray/scale.yml -b -i inventory/mycluster/hosts.ini --limit NODE_NAME
```

## Integrate a control-plane node into k8s

To add a new control plane node, run the following command.

```Bash
ansible-playbook collections/kubespray/cluster.yml -b -i inventory/sample/hosts.ini
```

## <a name="new_kinds"></a>New node kinds (processing, etc.)

The group for your nodes does not exist by default in the *hosts* file. You may add it yourself following the example below.

```ini
# hosts.ini

# Replace processing by the name of your new group.
[processing]
YOUR_NODE_NAME

...

# Add your newly created group to kube_node children.
# Replace processing by the name of your new group defined above.
[kube_node:children]
...
processing

```

You can add specific labels and taints to your node by adding a section in `hosts.ini` like the example below.

```ini
# hosts.ini

# Replace processing by the name of your new group defined above.
[processing:vars]
node_labels={"node-role.kubernetes.io/processing":""}

# node_labels={"LABEL_PREFIX/LABEL_KEY":"LABEL_VALUE", "LABEL2_PREFIX/LABEL2_KEY":"LABEL2_VALUE" , ...}
# if your label does not have a value, set the value to an empty string

# you can also add taints
# node_taints={"LABEL_PREFIX/LABEL_KEY:TAINT", "LABEL2_PREFIX/LABEL2_KEY:TAINT", ...}
# example: node_taints={"node-role.kubernetes.io/gateway:NoSchedule"}
```

After that, reach to [Integrate a worker node into k8s](#worker_nodes) to resume the process.

## Configure the new nodes

Once added to the Kubernetes cluster, you need to set some specific configuration on the node (firewall, DNS, public IP address, etc.).

Run the following command.

```Bash
ansible-playbook rs-setup.yaml -i inventory/cluster/hosts.ini
```

You can focus the playbook for only a specific node adding the option `--limit NODE_NAME`.

## Gateways

We consider a SafeScale gateway as a worker node into a Kubernetes cluster. Hence, follow the procedure written above about [how to add a worker node](#worker_nodes).

A gateway node requires some additional configuration:

1. Add a public IP address.
2. Assign the following additional security groups.
    - `safescale-sg_subnet_publicip.SUBNET_NAME.NETWORK_NAME`
    - `safescale-sg_subnet_gateways.SUBNET_NAME.NETWORK_NAME`

Default SUBNET_NAME and NETWORK_NAME are CLUSTER_NAME.

# Delete a node

> LIMITATION: You cannot remove the default gateways and masters deployed with the cluster. Only the worker nodes and the masters and gateway added with scaling.

> LIMITATION: SafeScale do not support to delete a specific node by name. You can only remove the n last added nodes.


While the node you want to delete is still present in the Ansible inventory, run the Kubespray playbook `remove-node.yml`.

```Bash
ansible-playbook collections/kubespray/remove-node.yml -b -i inventory/mycluster/hosts.ini -e node=NODE_NAME
```

If the node is not online, run:

```Bash
ansible-playbook collections/kubespray/remove-node.yml -b -i inventory/mycluster/hosts.ini -e reset_nodes=false -e allow_ungraceful_removal=true
``` 

Finally, you can remove the node from the cluster.

```Bash
safescale cluster shrink CLUSTER_NAME --count NUMBER_OF_NODES_TO_REMOVE
``̀

Do not forget to update your *hosts* file once the node is removed.
