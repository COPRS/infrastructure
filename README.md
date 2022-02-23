:arrow_heading_up: Go back to the [Reference System Sotfware repository](https://github.com/COPRS/reference-system-software) :arrow_heading_up:

# Installation manual

## Overview

![](./doc/img/deployment.png)

> **Integrators' machine is called BASTION in the rest of the installation manual**

## _Bastion_ requirements

- Safescale: **>= v21.11.0** (https://github.com/CS-SI/SafeScale)
- openstacksdk: **>= v0.12.0** (https://pypi.org/project/openstacksdk/)
- qemu-system: **>= v4.2.1** (https://packages.ubuntu.com/focal/qemu-kvm / https://packages.ubuntu.com/focal/qemu-system-x86)
- Packer: **>= v1.7.8** (https://github.com/hashicorp/packer)
- python3
- python3-pip
- git

## Infrastructure requirements

- A minimalist configuration : [here](./CONFIG.md)
- A **domain name** publicly available with a wildcard **A** record.  

## Dependencies

### Kubespray
This project exploits Kubespray to deploy Kubernetes.  
The fully detailed documentation and configuration options are available on its page: [https://kubespray.io/](https://kubespray.io/)

### HashiCorp Vault (optional)
This project can integrate credentials from a custom `HashiCorp Vault` instance, see the specific documentation [here](doc/how-to/Credentials.md).

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

 - Virtual machines amount and sizing, buckets names in `inventory/mycluster/group_vars/all/safescale.yaml`
 - Credentials, domain name, certificates (see below), S3 endpoints in `infrastructure/inventory/mycluster/group_vars/all/main.yaml`
 - Packages paths containing the apps to be deployed in `inventory/mycluster/group_vars/all/app_installer.yaml`

5. ### If needed create an image for the machines with `packer`
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

8. ### Deploy kubernetes with `kubespray`

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

10. ### Prepare the cluster for **Reference System**
```shellsession
ansible-playbook rs-setup.yaml \
    -i inventory/mycluster/hosts.ini
```

11. ### Generate or download the inventory variables
```shellsession
ansible-playbook generate_inventory.yaml \
    -i inventory/mycluster/hosts.ini
```

12. ### Set up the SSL certificates and the Stash license
See the documentation [here](doc/how-to/Certificates.md).

13. ### Deploy the apps 
```shellsession
ansible-playbook apps.yaml \
    -i inventory/mycluster/hosts.ini
```

## Post installation

- *NOT MANDATORY* : A **load balancer** listening on the public IP address pointed to by the domain name.  
  Configure the load balancer to forward incoming flow toward the cluster masters.

  | Load balancer port | masters port |
  | :---: | :---: |
  | 80 | 32080 |
  | 443 | 32443 |

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
            - `safescale.yaml`: Configuration of the machines, networks and buckets created by SafeScale.cluster and the OS image.
            - **apps**: One file per app deployed containing specific variables.
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

Configurations proposed by default:
- **Cert manager**
  - Helm chart:
    - Repository: charts.jetstack.io
    - Version: v1.6.1
    - Source: https://github.com/cert-manager/cert-manager/blob/master/deploy/charts/cert-manager
  - Images:
    - quay.io/jetstack/cert-manager-cainjector:v1.6.1
- **Linkerd CNI**
  - Helm chart:
    - Repository: helm.linkerd.io/stable
    - Version: 2.11.1
    - Source: https://github.com/linkerd/linkerd2/tree/main/charts/linkerd2-cni
  - Images:
    - cr.l5d.io/linkerd/cni-plugin:stable-2.11.1
- **Linkerd**
  - Helm chart:
    - Repository: helm.linkerd.io/stable
    - Version: 2.11.1
    - Source: https://github.com/linkerd/linkerd2/tree/main/charts/linkerd-control-plane
  - Images:
    - cr.l5d.io/linkerd/policy-controller:stable-2.11.1
    - cr.l5d.io/linkerd/proxy:stable-2.11.1
    - cr.l5d.io/linkerd/controller:stable-2.11.1
    - cr.l5d.io/linkerd/debug:stable-2.11.1
- **Linkerd Viz**
  - Helm chart:
    - Repository: helm.linkerd.io/stable
    - Version: 2.11.1
    - Source: https://github.com/linkerd/linkerd2/tree/main/viz/charts/linkerd-viz
  - Images:
    - cr.l5d.io/linkerd:stable-2.11.1
- **Rook Ceph** 
  - Helm chart: 
    - Repository: charts.rook.io/release
    - Version: v1.7.7
    - Source: https://github.com/rook/rook/tree/master/deploy/charts/rook-ceph
  - Images:
    - quay.io/cephcsi/cephcsi:v3.4.0
    - k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.3.0
    - k8s.gcr.io/sig-storage/csi-resizer:v1.3.0
    - k8s.gcr.io/sig-storage/csi-provisioner:v3.0.0
    - k8s.gcr.io/sig-storage/csi-snapshotter:v4.2.0
    - k8s.gcr.io/sig-storage/csi-attacher:v3.3.0
    - quay.io/csiaddons/volumereplication-operator:v0.1.0
- **Rook Ceph Cluster**
  - Helm chart:
    - Repository: charts.rook.io/release
    - Version: v1.7.7
    - Source: https://github.com/rook/rook/tree/master/deploy/charts/rook-ceph-cluster
  - Images:
    - quay.io/ceph/ceph:v16.2.6
- **ECK Operator**
  - Helm chart:
    - Repository: helm.elastic.co
    - Version: 1.9.0
    - Source: https://github.com/elastic/cloud-on-k8s/tree/master/deploy/eck-operator
  - Images:
    - docker.elastic.co/eck/eck-operator:1.9.0
    - docker.elastic.co/elasticsearch/elasticsearch:7.15.2
    - docker.elastic.co/kibana/kibana:7.15.2
    - quay.io/prometheuscommunity/elasticsearch-exporter:v1.3.0
- **Grafana Operator**
  - Helm chart: *None*
  - Images:
    - quay.io/grafana-operator/grafana-operator:v4.1.1
    - docker.io/grafana/grafana:8.3.3-ubuntu
- **Kafka Operator**
  - Helm chart:
    - Repository: strimzi.io/charts/
    - Version: 0.27.1
    - Source: https://github.com/strimzi/strimzi-kafka-operator/tree/main/helm-charts/helm3/strimzi-kafka-operator
  - Images:
    - quay.io/strimzi/operator:0.27.1
    - quay.io/strimzi/kafka:0.27.1-kafka-2.8.1
- **Prometheus Operator**
  - Helm chart:
    - Repository: prometheus-community.github.io/helm-charts
    - Version: 21.0.0
    - Source: https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-prometheus-stack
  - Images:
    - quay.io/prometheus-operator/prometheus-operator:v0.52.1-amd64
    - quay.io/prometheus-operator/prometheus-config-reloader:v0.52.1-amd64
    - quay.io/prometheus/node-exporter:v1.3.0
    - quay.io/prometheus/alertmanager:v0.23.0
    - quay.io/prometheus/prometheus:v2.31.1
- **Stash Operator**
  - Helm chart:
    - Repository: charts.appscode.com/stable
    - Version: v0.17.0
    - Source: https://github.com/stashed/installer/tree/master/charts/stash-community
  - Images:
    - docker.io/appscode/stash:v0.17.0
- **Fluent-bit**
  - Helm chart:
    - Repository: fluent.github.io/helm-charts
    - Version: 0.19.6
    - Source: https://github.com/fluent/helm-charts/tree/main/charts/fluent-bit
  - Images:
    - docker.io/fluent/fluent-bit:1.8.10
- **MongoDB**
  - Helm chart:
    - Repository: charts.bitnami.com/bitnami
    - Version: 11.0.3
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/mongodb
  - Images:
    - docker.io/bitnami/mongodb:5.0.6-debian-10-r14
    - docker.io/bitnami/mongodb-exporter:0.30.0-debian-10-r58
- **OpenLDAP**
  - Helm chart: *None*
  - Images:
    - docker.io/osixia/openldap:1.5.0
- **PostreSQL**
  - Helm chart:
    - Repository: charts.bitnami.com/bitnami
    - Version: 11.0.2
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/postgresql
  - Images:
    - docker.io/bitnami/postgresql:14.1.0-debian-10-r80
    - quay.io/prometheuscommunity/postgres-exporter:v0.10.0
- **Thanos**
  - Helm chart:
    - Repository: charts.bitnami.com/bitnami
    - Version: 8.1.2
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/thanos
  - Images:
    - docker.io/bitnami/thanos:0.23.1-scratch-r3
- **Fluentd**
  - Helm chart:
    - Repository: charts.bitnami.com/bitnami
    - Version: 4.4.1
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/fluentd
  - Images:
    - docker.io/bitnami/fluentd:1.14.2-debian-10-r23
- **Loki**
  - Helm chart:
    - Repository: grafana.github.io/helm-charts
    - Version: 2.8.1
    - Source: https://github.com/grafana/helm-charts/tree/main/charts/loki
  - Images:
    - docker.io/grafana/loki:2.4.1
- **Apisix**
  - Helm chart:
    - Repository: charts.apiseven.com
    - Version: 0.7.3
    - Source: https://github.com/apache/apisix-helm-chart/tree/master/charts/apisix
  - Images:
    - docker.io/apache/apisix:2.10.0-alpine
    - docker.io/apache/apisix-dashboard:2.10.1-alpine
    - docker.io/apache/apisix-ingress-controller:1.3.0
    - docker.io/bitnami/etcd:3.4.16-debian-10-r14
- **Falco**
  - Helm chart:
    - Repository: falcosecurity.github.io/charts
    - Version: 1.16.2
    - Source: https://github.com/falcosecurity/charts/tree/master/falco
  - Images:
    - docker.io/falcosecurity/falco:0.30.0
    - docker.io/falcosecurity/falco-exporter:0.6.0
- **FinOps object storage exporter**
  - Helm chart:
    - Repository: artifactory.coprs.esa-copernicus.eu/artifactory/rs-helm
    - Version: 1.0.0
    - Source: *Private*
  - Images:
    - artifactory.coprs.esa-copernicus.eu/cs-docker/finops-object-storage-exporter:release-0.3.0
- **FinOps resources exporter**
  - Helm chart:
    - Repository: artifactory.coprs.esa-copernicus.eu/artifactory/rs-helm
    - Version: 1.0.0
    - Source: *Private*
  - Images:
    - artifactory.coprs.esa-copernicus.eu/cs-docker/finops-resources-exporter:release-0.3.0
- **Graylog**
  - Helm chart:
    - Repository: charts.kong-z.com
    - Version: 1.9.2
    - Source: https://github.com/KongZ/charts/tree/main/charts/graylog
  - Images:
    - docker.io/graylog/graylog:4.2.3-1
- **Keycloack**
  - Helm chart:
    - Repository: codecentric.github.io/helm-charts
    - Version: 16.0.5
    - Source: https://github.com/codecentric/helm-charts/tree/master/charts/keycloak
  - Images:
    - docker.io/jboss/keycloak:15.0.2
- **Spring Cloud Data Flow**
  - Helm chart:
    - Repository: charts.bitnami.com/bitnami
    - Version: 4.1.5
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/spring-cloud-dataflow
  - Images:
    - docker.io/bitnami/spring-cloud-dataflow:2.9.1-debian-10-r7
    - docker.io/bitnami/spring-cloud-skipper:2.8.1-debian-10-r6
    - docker.io/bitnami/prometheus-rsocket-proxy:1.3.0-debian-10-r334
