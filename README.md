:arrow_heading_up: Go back to the [Reference System Sotfware repository](https://github.com/COPRS/reference-system-software) :arrow_heading_up:

# Installation manual

## Overview

![](./doc/img/deployment.png)

> **Integrators' machine is called BASTION in the rest of the installation manual**

## _Bastion_ requirements

- Safescale : **>= v21.11.0** (https://github.com/CS-SI/SafeScale)
- openstacksdk : **>= v0.12.0** (https://pypi.org/project/openstacksdk/)
- qemu-system : **>= v4.2.1** (https://packages.ubuntu.com/focal/qemu-kvm / https://packages.ubuntu.com/focal/qemu-system-x86)
- Packer : **>= v1.7.8** (https://github.com/hashicorp/packer)
- python3
- python3-pip
- git

## Infrastructure requirements

- A **domain name** publicly available with a wildcard **A** record.  
- A **load balancer** listening on the public IP address pointed to by the domain name.  
  Configure the load balancer to forward incoming flow toward the cluster masters.

  | Load balancer port | masters port |
  | :---: | :---: |
  | 80 | 32080 |
  | 443 | 32443 |
- A **Stash community** licence, to get [here](https://license-issuer.appscode.com/?p=stash-community) between the `kubespray.yaml` and the `apps.yaml` playbooks


## Quickstart

1. ### Get the infrastructure repository
```shellsession
git clone https://github.com/COPRS/infrastructure.git
```

2. ### Install requirements
```shellsession
cd infrastructure

git submodule update --init

python3 -m pip install --user -r collections/kubespray/requirements.txt
ansible-galaxy collection install \
    kubernetes.core \
    openstack.cloud
```

3. ### Copy the sample inventory
```shellsession
cp -rfp inventory/sample inventory/mycluster
```


4. ### Review and change the default configuration to match your needs

 - Virtual machines amount and sizing in `inventory/mycluster/host_vars/localhost/cluster_safescale.yaml`
 - Credentials, domain name, certificates (see below), S3 endpoints and buckets in `infrastructure/inventory/sample/group_vars/gateway/app_vars.yaml`
 - Packages containing the apps to be deployed in `inventory/sample/group_vars/gateway/app_installer.yaml`

5. ### If needed create an image for the machines with **Packer**
```shellsession
ansible-playbook image.yaml \
    -i inventory/mycluster/hosts.ini
```

6. ### Deploy machines with safescale
```shellsession
ansible-playbook cluster-setup.yaml \
    -i inventory/mycluster/hosts.ini
```

7. ### Install security services
```shellsession
ansible-playbook security.yaml \
    -i inventory/mycluster/hosts.ini \
    --become
```

8. ### Deploy kubernetes with Kubespray

```shellsession
# The option `--become` is required, for example writing SSL keys in /etc/,
# installing packages and interacting with various systemd daemons.
# Without --become the playbook will fail to run!

ansible-playbook collections/kubespray/cluster.yml \
    -i inventory/mycluster/hosts.ini \
    --become
```

9. ### Enable pod security policies on the cluster
```shellsession
# /!\ create first the PSP and CRB resources
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
```

10. ### Prepare the cluster for Reference System
```shellsession
ansible-playbook rs-setup.yaml \
    -i inventory/mycluster/hosts.ini
```

11. ### Deploy the apps 
```shellsession
ansible-playbook apps.yaml \
    -i inventory/mycluster/hosts.ini
```

## TLS configuration

Reference System exploits [APISIX Ingress Controller](https://apisix.apache.org/) and [Cert Manager](https://cert-manager.io/) for TLS configuration.

You need to create an [issuer](https://cert-manager.io/docs/concepts/issuer/) and a [certificate](https://cert-manager.io/docs/concepts/certificate/) for your domain name with Cert Manager.

APISIX does not work with Cert Manager for ACME HTTP01 challenges ([#781](https://github.com/apache/apisix-ingress-controller/issues/781)).  
You must use the DNS01 challenge to generate a Let's encrypt certificate. The configuration is detailled on [Cert Manager documentation](https://cert-manager.io/docs/configuration/acme/dns01).

## Dependencies

This project exploits Kubespray to deploy Kubernetes.  
The fully detailed documentation and configuration options are available on its page: [https://kubespray.io/](https://kubespray.io/#/)

## Tree view

The repository is made of the following main directories and files.
- **apps**: A package example, gathering default applications deployed with Reference System platform.
- **collections/kubespray**: folder where kubespray is integrated into the project as a git submodule.
    - `cluster.yml: The Ansible playbook to run to deploy Kubernetes or to add a master node.
    - `scale.yml`: An Ansible playbook to add a worker node.
    - `remove-node.yml`: An Ansible playbook to remove a node.
- **doc**: Here we find all the documentation describing the infrastructure deployment and maintenance operations.
- **inventory**:
  - **sample**: An Ansible inventory for a sample configuration.
      - **group_vars**:
          - **gateway**:
            - `app_installer.yaml`: The configuration of the app installer roles. It includes the list and paths of packages to install.
            - `app_vars.yaml`: Reference the different variables to set and configure to deploy the applications.
          - `k8s_cluster/kubespray.yaml`: The Kubespray configuration.
      - **host_vars/localhost**: The configuration parameters for SafeScale cluster and the OS image.
      - `hosts.ini`: The list of machines described in their respective groups.
- **roles**: The list of roles used to deploy the cluster.
- `ansible.cfg`: Ansible configuration file. It includes the ssh configuration to allow Ansible to access the machines through the gateway.
- `apps.yaml`: An Ansible playbook to deploy the applications on the platform.
- `cluster-setup.yaml`: An Ansible playbook to create the hosts, network and volumes with SafeScale.
- `delete.yaml`: An Ansible playbook to delete a SafeScale cluster and remove all the generated resources.
- `image.yaml`: An Ansible playbook to build a golden OS image.
- `rs-setup.yaml`: An Ansible playbook to configure the nodes.
- `security.yaml`: An Ansible playbook to install the security services on the nodes.

### Playbooks manual

| name | tags | utility | 
|---|---|---|
| apps.yaml | *none* |  *deploy applications*<br>Supported possible options:<br>**-e app=APP_NAME** Deploy only a specific application.<br>**-e debug=true** Keep the application resources generated for debugging.<br>**-e package=PACKAGE_NAME** Deploy only a specific package.|
| cluster-setup.yaml | *none* <br> cluster_create <br> hosts_update <br> volumes_create | *all tags below are executed* <br> create safescale cluster <br> update hosts.ini with newly created machines, fill .ssh folder with machines ssh public keys, generate ansible ssh config, update config.cfg <br> attach disks to kubernetes nodes |
| delete.yaml <br> :warning: this playbook has been developed with the only purpose of testing the project **not for production usage**| *none* <br> cleanup_generated <br> detach_volumes <br> delete_volumes <br> delete_cluster | *nothing* <br> **remove** ssh keys, added hosts in hosts.ini, ssh config file <br> detach added disks from k8s nodes <br> delete added disks from k8s nodes <br> delete safescale cluster|
| image.yaml | *none* | *make reference system golden image for k8s nodes* |
|rs-setup.yaml | *none* <br> gateway <br> | *all tags below are executed* <br> install the tools on gateways <br> configure the cluster |
| security.yaml | *none* <br> auditd <br> wazuh <br> clamav <br> openvpn <br> suricata <br> uninstall_APP_NAME| *install all security tools* <br> install auditd <br> install wazuh <br> install clamav <br> install openvpn <br> install suricata <br> uninstall the app matching APP_NAME

## Apps

Configurations proposed by default :
- **Rook Ceph** 
  - Helm chart: 
    - URL : charts.rook.io/release
    - Version : 1.7.7
    - Documentation : https://github.com/rook/rook/blob/master/Documentation/helm-operator.md
  - Images
    - Ceph
      - Registry : Docker Hub
      - Repository : ceph
      - Version : 1.7.7
    - CSI
      - Ceph
        - Registry : quay.io
        - Version : 3.4.0
      - Registrar
        - Registry : k8s.gcr.io
        - Version : 2.3.0
      - Resizer
        - Registry : k8s.gcr.io
        - Version : 1.3.0
      - Provisioner
        - Registry : k8s.gcr.io
        - Version : 3.0.0
      - Snapshotter
        - Registry : k8s.gcr.io
        - Version : 4.2.0
      - Attacher
        - Registry : k8s.gcr.io
        - Version : 3.3.0
      - Volume Replication
        - Registry : quay.io
        - Version : 0.1.0
- **Rook Ceph Cluster**
  - Helm chart:
    - URL : charts.rook.io/release
    - Version : 1.7.7
    - Documentation : https://github.com/rook/rook/blob/master/Documentation/helm-ceph-cluster.md
  - Image
    - Registry : quay.io
    - Version : 16.2.6
- **Kafka Operator**
  - Helm chart:
    - URL : strimzi.io/charts/
    - Version : 0.26.0
    - Documentation : https://github.com/strimzi/strimzi-kafka-operator/tree/main/helm-charts/helm3/strimzi-kafka-operator
  - Images
    - Registry : quay.io
    - Versions
      - Operator : 0.27.1
      - Kafka : 2.8.1
      - Zookeeper : 3.5.9
- **Elasticsearch Operator**
  - Helm chart:
    - URL : helm.elastic.co
    - Version : 1.9.0
    - Source : https://github.com/elastic/cloud-on-k8s/tree/master/deploy/eck-operator
  - Images
    - Registry : docker.elastic.co
    - Versions
      - Operator :1.9.0
      - Elasticsearch : 7.15.2
      - Kibana : 7.15.2
- **PostreSQL**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 11.0.2
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/postgresql
  - Images
    - Registry : Docker Hub
    - Repository : bitnami
    - Versions
      - PostgreSQL : 14.1.0
      - Exporter : 0.10.0
- **MongoDB**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 11.0.3
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/mongodb
  - Images
    - Registry : Docker Hub
    - Repository : bitnami
    - Versions
      - MongoDB : 5.0.6
      - Exporter : 0.30.0
- **Graylog**
  - Helm chart:
    - URL : charts.kong-z.com
    - Version : 1.9.2
    - Documentation : https://github.com/KongZ/charts/tree/main/charts/graylog
  - Image
    - Registry : Docker Hub
    - Repository : graylog
    - Version : 4.2.3
- **Spring Cloud Data Flow**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 4.1.5
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/spring-cloud-dataflow
  - Images
    - Registry : Docker Hub
    - Repository : bitnami
    - Versions
      - Server : 2.9.1
      - Skipper : 2.8.1
      - WaitForBackend : 1.22.2
      - Exporter : 1.3.0
- **Loki**
  - Helm chart:
    - URL : grafana.github.io/helm-charts
    - Version : 2.8.1
    - Documentation : https://github.com/grafana/helm-charts/tree/main/charts/loki
  - Image
    - Registry : Docker Hub
    - Repository : grafana
    - Version : 2.4.1
- **Fluentbit**
  - Helm chart:
    - URL : fluent.github.io/helm-charts
    - Version : 0.19.6
    - Documentation : https://github.com/fluent/helm-charts/tree/main/charts/fluent-bit
  - Image
    - Registry : Docker Hub
    - Repository : fluent
    - Version : 1.8.10
- **Fluentd**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 4.4.1
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/fluentd
  - Image
    - Registry : Docker Hub
    - Repository : bitnami
    - Version : 1.14.2
- **Grafana**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 7.2.2
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/grafana-operator
  - Image
    - Registry : Docker Hub
    - Repository : bitnami
    - Version : 8.2.5
- **Prometheus Stack**
  - Helm chart:
    - URL : prometheus-community.github.io/helm-charts
    - Version : 21.0.0
    - Documentation : https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack
  - Images
    - Registry : quay.io
    - Versions : 
      - AlertManager : 0.23.0
      - Node Exporter : 1.3.0
      - Operator : 0.52.1
      - Prometheus : 2.31.1
- **Thanos**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 8.1.2
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/thanos
  - Image
    - Registry : Docker Hub
    - Repository : bitnami
    - Version : 0.23.1
- **Keycloack**
  - Helm chart:
    - URL : charts.bitnami.com/bitnami
    - Version : 5.2.0
    - Documentation : https://github.com/bitnami/charts/tree/master/bitnami/keycloak
  - Image
    - Registry : Docker Hub
    - Repository : bitnami
    - Version : 15.0.2
- **OpenLDAP**
  - Helm chart: _None_
  - Image
    - Registry : Docker Hub
    - Repository : osixia
    - Version : 1.5.0
- **Falco**
  - Helm chart:
    - URL : https://github.com/falcosecurity/charts
    - Version : 1.16.2
    - Documentation : https://github.com/falcosecurity/charts
  - Image
    - Registry : Docker Hub
    - Repository : falcosecurity/falco
    - Version : 0.30.0
