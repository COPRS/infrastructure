:arrow_heading_up: Go back to the [Reference System Sotfware repository](https://github.com/COPRS/reference-system-software) :arrow_heading_up:

# Infrastructure - Installation manual

## Overview

![deployment](./docs/media/deployment.png)

> **Integrators' machine is called BASTION in the rest of the installation manual**

## _Bastion_ requirements

- Safescale: **>= v22.11.6** (<https://github.com/CS-SI/SafeScale>)
- openstacksdk: **>= v0.12.0** (<https://pypi.org/project/openstacksdk/>)
- qemu-system: **>= v4.2.1** (<https://packages.ubuntu.com/focal/qemu-kvm> / <https://packages.ubuntu.com/focal/qemu-system-x86>)
- Packer: **>= v1.7.8** (<https://github.com/hashicorp/packer>)
- python3
- python3-pip
- git
- jq
- cloud-image-utils

## Infrastructure requirements

- Default configuration : [here](./CONFIG.md)
- A **domain name** publicly available with a wildcard **A** record.  

## Dependencies

### Kubespray

This project exploits Kubespray to deploy Kubernetes.  
The fully detailed documentation and configuration options are available on its page: [https://kubespray.io/](https://kubespray.io/)

### HashiCorp Vault (optional)

This project can integrate credentials from a custom `HashiCorp Vault` instance, see the specific documentation [here](./docs/user_manuals/how-to/Credentials.md).

## Quickstart

### 1. Get the infrastructure repository

```shellsession
git clone https://github.com/COPRS/infrastructure.git
```

### 2. Install requirements

```shellsession
cd infrastructure

git submodule update --init

python3 -m pip install --user \
  pyOpenSSL ecdsa \
  -r collections/kubespray/requirements.txt

ansible-galaxy collection install \
    kubernetes.core \
    openstack.cloud
```

### 3. Copy the sample inventory

```shellsession
cp -rfp inventory/sample inventory/mycluster
```

### 4. Review and change the default configuration to match your needs

- Node groups and S3 buckets in `inventory/mycluster/host_vars/setup/safescale.yaml`
- Credentials, domain name, the stash license, S3 endpoints in `infrastructure/inventory/mycluster/host_vars/setup/main.yaml`
- Packages paths containing the apps to be deployed in `inventory/mycluster/host_vars/setup/app_installer.yaml`

### 5. Generate or download the inventory variables

```shellsession
ansible-playbook generate_inventory.yaml \
    -i inventory/mycluster/hosts.yaml
```

### 6. If needed create an image for the machines with `packer`

```shellsession
ansible-playbook image.yaml \
    -i inventory/mycluster/hosts.yaml
```

### 7. Create and configure machines

```shellsession
ansible-playbook cluster.yaml \
    -i inventory/mycluster/hosts.yaml
```

### 8. Install security services

```shellsession
ansible-playbook security.yaml \
    -i inventory/mycluster/hosts.yaml \
    --become
```

### 9. Deploy kubernetes with `kubespray`

```shellsession
# The option `--become` is required, for example writing SSL keys in /etc/,
# installing packages and interacting with various systemd daemons.
# Without --become the playbook will fail to run!

ansible-playbook collections/kubespray/cluster.yml \
    -i inventory/mycluster/hosts.yaml \
    --become
```

### 10. Enable pod security policies (PSP) on the cluster

```shellsession
# /!\ create first the PSP and ClusterRoleBinding resources
# before enabling the admission plugin

ansible-playbook collections/kubespray/upgrade-cluster.yml \
    -i inventory/mycluster/hosts.yaml \
    --tags cluster-roles \
    -e podsecuritypolicy_enabled=true \
    --become

ansible-playbook collections/kubespray/upgrade-cluster.yml \
    -i inventory/mycluster/hosts.yaml \
    --tags master \
    -e podsecuritypolicy_enabled=true \
    --become
```

### 11. Setup RS specifics

```shellsession
ansible-playbook rs-setup.yaml \
    -i inventory/mycluster/hosts.yaml
```

### 12. Add the providerID spec to the nodes for the autoscaling

```shellsession
ansible-playbook cluster.yaml -i inventory/mycluster/hosts.yaml -t providerids
```

### 13. Deploy the apps

```shellsession
ansible-playbook apps.yaml \
    -i inventory/mycluster/hosts.yaml
```

## Post installation

- User's Manual : [here](./docs/user_manuals/README.md)
- _NOT MANDATORY_ : A **load balancer** listening on the public IP address pointed to by the domain name.  
  Configure the load balancer to forward incoming flow toward the cluster masters.

  | Load balancer port | masters port | protocol |
  | :---: | :---: | :--: |
  | 80 | 32080 | TCP |
  | 443 | 32443 | TCP |

> For **health check** : `https://node-ip:6443/readyz` <br>
> `node-ip` : private ip of each master

- You may disable access to Keycloak master realm. From Apisix interface: open Route tab, search for `iam_keycloak_keycloak-superadmin` and click on `Offline`.

## Build

In order to build the [autoscaling component](./docs/user_manuals/how-to/Cluster%20scaling.md#autoscaling) check from source, first clone the GitHub repository :

### 1. Build the cluster-autoscaler

```Bash
NAME_IMAGE="cluster-autoscaler" ;
CA_TAG="1.22.3" ;
REGISTRY_BASE="<CHANGE_ME>" ;
PROJECT="<CHANGE_ME>" ;

git clone https://github.com/COPRS/infrastructure.git

mkdir scaler/build ;
wget https://github.com/kubernetes/autoscaler/archive/refs/tags/cluster-autoscaler-${CA_TAG}.tar.gz -O scaler/build/ca.tar.gz ;

cd scaler/build ;
tar -xzf ca.tar.gz autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler ;

cd scaler/build/autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler ;
make build-arch-amd64 ;

cd scaler/build/autoscaler-cluster-autoscaler-${CA_TAG}/cluster-autoscaler ;
docker build -t ${REGISTRY_BASE}/${PROJECT}/${NAME_IMAGE}:${CA_TAG} -f Dockerfile.amd64 . ;
```

### 2. Build the rs-infra-autoscaler

```Bash
NAME_IMAGE="rs-infra-scaler" ;
SAFESCALE_TAG="v22.11.6" ;
SCALER_TAG="1.6.0" ;
REGISTRY_BASE="<CHANGE_ME>" ;
PROJECT="<CHANGE_ME>" ;


git clone https://github.com/COPRS/infrastructure.git

mkdir scaler/ansible_resources ;
cp -r *.yaml ansible.cfg roles inventory scaler/ansible_resources ;

cd scaler ;
docker build --build-arg SAFESCALE_TAG=${SAFESCALE_TAG} -t ${REGISTRY_BASE}/${PROJECT}/${NAME_IMAGE}:${SCALER_TAG} -f Dockerfile .
```

### 3. Build the safescale daemon

```Bash
NAME_IMAGE="safescaled" ;
SAFESCALE_TAG="v22.11.6" ;
REGISTRY_BASE="<CHANGE_ME>" ;
PROJECT="<CHANGE_ME>" ;

git clone https://github.com/COPRS/infrastructure.git

docker build --build-arg SAFESCALE_TAG=${SAFESCALE_TAG} -t ${REGISTRY_BASE}/${PROJECT}/${NAME_IMAGE}:${SAFESCALE_TAG} -f scaler/safescaled.Dockerfile .
```

## Known issues and limitations

The project and this component has few known issues and limitations. Head over the [limitation](./docs/user_manuals/how-to/Limitations.md) page to learn more about it.

## Tree view

The repository is made of the following main directories and files.

- **apps**: A package example, gathering default applications deployed with Reference System platform.
- **collections/kubespray**: folder where kubespray is integrated into the project as a git submodule.
  - `cluster.yml`: The Ansible playbook to run to deploy Kubernetes or to add a master node.
  - `scale.yml`: An Ansible playbook to add a worker node.
  - `remove-node.yml`: An Ansible playbook to remove a node.
- **doc**: Here we find all the documentation describing the infrastructure deployment and maintenance operations.
- **inventory**:
  - **sample**: An Ansible inventory for a sample configuration.
    - **group_vars**:
      - **all**:
        - `app_installer.yaml`: The configuration of the app installer roles. It includes the list and paths of packages to install.
        - `main.yaml`: Mandatory platform configuration.
        - `kubespray.yaml`: The Kubespray configuration.
        - `safescale.yaml`: Configuration of the machines, networks and buckets created by SafeScale, and more.
        - **apps**: One file per app deployed containing specific variables.
    - `hosts.yaml`: The list of machines described in their respective groups, this file is managed by the `cluster.yaml` playbook.
- **roles**: The list of roles used to deploy the cluster.
- `ansible.cfg`: Ansible configuration file. It includes the ssh configuration to allow Ansible to access the machines through the gateway.
- `apps.yaml`: An Ansible playbook to deploy the applications on the platform.
- `cluster.yaml`: An Ansible playbook to manage the safescale cluster and its machines
- `delete.yaml`: An Ansible playbook to delete a SafeScale cluster and remove all the generated resources.
- `generate_inventory.yaml`: An Ansible playbook to generate inventory vars.
- `image.yaml`: An Ansible playbook to build a golden OS image.
- `security.yaml`: An Ansible playbook to install the security services on the nodes.

### Playbooks manual

| name | tags | utility |
|---|---|---|
| apps.yaml | _none_ |  _deploy applications_<br>Supported possible options:<br>**-e app=APP_NAME** Deploy only a specific application.<br>**-e debug=true** Keep the application resources generated for debugging.<br>**-e package=PACKAGE_NAME** Deploy only a specific package.|
| cluster.yaml | _none_  <br> create_cluster <br> config <br> gateway <br> update_hosts <br> providerids <br> | _all tags below are executed_ <br> create safescale cluster <br> configure cluster machines <br> configure gateways <br> update hosts.yaml with newly created machines, fill .ssh folder with machines ssh private keys <br> write providerID spec to kube nodes |
| delete.yaml <br> :warning: this playbook has been developed with the only purpose of testing the project **not for production usage**| _none_ <br> cleanup_generated <br> detach_volumes <br> delete_volumes <br> delete_cluster | _nothing_ <br> **remove** ssh keys, added hosts in hosts.yaml, ssh config file <br> detach added disks from k8s nodes <br> delete added disks from k8s nodes <br> delete safescale cluster|
| generate_inventory.yaml | _none_ | _Generate/download/upload inventory vars in group_vars/all_ |
| image.yaml | _none_ | _make reference system golden image for k8s nodes_ |
| security.yaml | _none_ <br> auditd <br> wazuh <br> clamav <br> openvpn <br> suricata <br> uninstall_APP_NAME| _install all security tools_ <br> install auditd <br> install wazuh <br> install clamav <br> install openvpn <br> install suricata <br> uninstall the app matching APP_NAME |

# Copyright and license

The Reference System Software as a whole is distributed under the Apache License, version 2.0. A copy of this license is available in the [LICENSE](LICENSE) file. Reference System Software depends on third-party components and code snippets released under their own license (obviously, all compatible with the one of the Reference System Software). These dependencies are listed in the [NOTICE](NOTICE.md) file.
