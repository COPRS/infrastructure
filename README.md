# Installation manual

## Overview
![](./doc/img/deployment.png)

> **Integrators' machine is called BASTION in the rest of the installation manual**

## _Bastion_ requirements

- Safescale : **>= v21.05.0-rc3** (https://github.com/CS-SI/SafeScale)
- openstacksdk : **>= v0.12.0** (https://pypi.org/project/openstacksdk/)
- qemu-system : **>= v4.2.1** (https://packages.ubuntu.com/focal/qemu-kvm / https://packages.ubuntu.com/focal/qemu-system-x86)
- Packer : **>= v1.7.8** (https://github.com/hashicorp/packer)

## Quick start

```Bash
## ON BASTION

# get the infrastructure repository
git clone -b 0.1.0-rci https://github.com/COPRS/infrastructure.git

cd infrastructure

# install requirements
git submodule update --init
python3 -m pip install --user -r platform/collections/kubespray/requirements.txt
ansible-galaxy collection install \
    kubernetes.core \
    openstack.cloud

# access platform repository
cd platform

# Copy ``inventory/sample`` as ``inventory/mycluster``
cp -rfp inventory/sample inventory/mycluster

# Review and change paramters under ``inventory/mycluster/group_vars`` or ``inventory/mycluster/host_vars``
cat inventory/mycluster/host_vars/localhost/cluster.yaml
cat inventory/mycluster/host_vars/localhost/image.yaml
cat inventory/mycluster/group_vars/all/kubespray.yaml
cat inventory/mycluster/group_vars/bastion/apps.yaml

# If needed create an image for the machines with Packer
ansible-playbook playbooks/image.yaml \
    -i inventory/mycluster/hosts.ini

# Deploy machines with safescale
ansible-playbook playbooks/cluster-setup.yaml \
    -i inventory/mycluster/hosts.ini

# Install security services
ansible-playbook playbooks/security.yaml \
    -i inventory/mycluster/hosts.ini \
    --become

# Deploy kubernetes with Kubespray - run the playbook as root
# The option `--become` is required, as for example writing SSL keys in /etc/,
# installing packages and interacting with various systemd daemons.
# Without --become the playbook will fail to run!
ansible-playbook collections/kubespray/cluster.yml \
    -i inventory/mycluster/hosts.ini \
    --become

# Enable pod security policies on the cluster
# /!\ you first need to create the psp and crb resources
# before enabling the admission plugin
ansible-playbook collections/kubespray/upgrade-cluster.yml \
    -i inventory/mycluster/hosts.ini \
    --tags cluster-roles \
    -e podsecuritypolicy_enabled=true \
    --become

ansible-playbook collections/kubespray/upgrade-cluster.yml \
    -i inventory/mycluster/hosts.ini \
    --tags master \
    -e podsecuritypolicy_enabled=true \
    --become

# Configure kubernetes and deploy apps
ansible-playbook playbooks/rs-setup.yaml \
    -i inventory/mycluster/hosts.ini
```

## Dependencies

This project exploit Kubespray to deploy Kubernetes.  
The fully detailled documentation and configuration options are available on its page: [https://kubespray.io/](https://kubespray.io/#/)

## Tree view
The repository is made of the following main directories.
- apps
- doc
- platform
### Apps
This folder gather the configuration of the applications deployed on the platform.  
Each application has its own folder inside apps with the values of the Helm chart, the kustomization files and the patches related.  
The application's directory can be splitted by environment with subfolders like dev, prod, etc.
### Doc
Here we find all the documentation describing the infrastructure deployment and maintenance operations.
### Platform
This directory concentrate what is required to deploy the infrastructure with Ansible.
- **collections/kubespray**: folder where kubespray is integrated in the project as a git submodule.
    - `cluster.yaml` the Ansible playbook to run to deploy Kubernetes.
- **inventory**: 
    - **sample**: Ansible inventory for sample configuration.
        - **group_vars**:
            - `all/kubespray.yaml`: kubespray configuration.
            - `bastion/app_installer.yaml`: application installer configuration: helm version, repositories, applications directory paths
        - **host_vars/localhost**: safescale cluster configuration and image configuration.
        - `hosts.ini`: host inventory.
- **playbooks**: list of Ansible playbooks to run to deploy the platform.
    - `clean.yaml`: remove the generated files by the different playbooks, delete the cluster and remove the volumes.
    - `cluster-setup.yaml`: deploy the network, the machines and the volumes with safescale.
    - `image.yaml`: build the image used to create the machines.
    - `rs-setup.yaml`: prepare the necessary resources for the platform and deploy the applications present in apps.
    - `security.yaml`: deploy the security services.
- **roles**: list of roles used to deploy the cluster.
    - **security**: roles describing the installation of the different security tools.
- `ansible.cfg`: Ansible configuration file. It includes the ssh configuration to allow Ansible to access the machines through the gateway.

## Playbooks

| name | tags | utility | 
|---|---|---|
| cluster-setup.yaml | *none* <br> cluster_create <br> hosts_update <br> volumes_create | *all tags below are executed* <br> create safescale cluster <br> update hosts.ini with newly created machines, fill .ssh folder with machines ssh public keys, generate ansible ssh config, update config.cfg <br> attach disks to kubernetes nodes |
| delete.yaml <br> :warning: this playbook has been developed with the only purpose of testing the project **not for production usage**| *none* <br> cleanup_generated <br> detach_volumes <br> delete_volumes <br> delete_cluster | *nothing* <br> **remove** ssh keys, added hosts in hosts.ini, ssh config file <br> detach added disks from k8s nodes <br> delete added disks from k8s nodes <br> delete safescale cluster|
| image.yaml | *none* | *make reference system golden image for k8s nodes* |
|rs-setup.yaml | *none* <br> gateway <br> apps | *all tags below are executed* <br> install tools on gateways <br> deploy applications (adding -e app_name=APP_NAME deploy only the app matching APP_NAME) |
| security.yaml | *none* <br> auditd <br> wazuh <br> clamav <br> openvpn <br> suricata <br> uninstall_APP_NAME| *install all securty tools* <br> install auditd <br> install wazuh <br> install clamav <br> install openvpn <br> install suricata <br> uninstall the app matching APP_NAME

## Apps

To configure apps, refer to the following Helm Charts :
- Rook Ceph : https://github.com/rook/rook/blob/master/Documentation/helm-operator.md
- Rook Ceph Cluster : https://github.com/rook/rook/blob/master/Documentation/helm-ceph-cluster.md
- Kafka : https://github.com/bitnami/charts/tree/master/bitnami/kafka
- PostreSQL : https://github.com/bitnami/charts/tree/master/bitnami/postgresql
- Elasticsearch : https://github.com/bitnami/charts/tree/master/bitnami/elasticsearch
- MongoDB : https://github.com/bitnami/charts/tree/master/bitnami/mongodb
- Graylog : https://github.com/KongZ/charts/tree/main/charts/graylog
- Spring Cloud Data Flow : https://github.com/bitnami/charts/tree/master/bitnami/spring-cloud-dataflow
