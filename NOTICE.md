Reference System Software Copyright 2021-2022 European Space Agency (ESA) https://www.esa.int

This software is distributed under the Apache Software License (ASL) v2.0, see LICENSE file or http://www.apache.org/licenses/LICENSE-2.0 for details.

Below are all the Commercial off-the-shelf (COTS) used and their respective licence:

- Cert manager :
  - Helm chart
    - Version: v1.6.1
    - License: [Apache License 2.0](https://github.com/cert-manager/cert-manager/blob/v1.6.1/LICENSE)
    - Source: https://github.com/cert-manager/cert-manager/tree/v1.6.1/deploy/charts/cert-manager
  - Container image(s)
    - quay.io/jetstack/cert-manager-cainjector:v1.6.1
      - License: [Apache License 2.0](https://github.com/cert-manager/cert-manager/blob/v1.6.1/LICENSE)

- Linkerd CNI
  - Helm chart
    - Version: 2.11.1
    - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - Source: https://github.com/linkerd/linkerd2/tree/stable-2.11.1/charts/linkerd2-cni
  - Container image(s)
    - cr.l5d.io/linkerd/cni-plugin:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)

- linkerd-control-plane
  - Helm chart
    - Version: 2.11.1
    - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - Source: https://github.com/linkerd/linkerd2/tree/main/charts/linkerd-control-plane
  - Container image(s)
    - cr.l5d.io/linkerd/policy-controller:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - cr.l5d.io/linkerd/proxy:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - cr.l5d.io/linkerd/controller:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - cr.l5d.io/linkerd/debug:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)

- Linkerd Viz
  - Helm chart
    - Version: 2.11.1
    - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)
    - Source: https://github.com/linkerd/linkerd2/tree/stable-2.11.1/viz/charts/linkerd-viz
  - Container image(s)
    - cr.l5d.io/linkerd:stable-2.11.1
      - License: [Apache License 2.0](https://github.com/linkerd/linkerd2/blob/stable-2.11.1/LICENSE)

- Rook Ceph
  - Helm chart
    - Version: v1.9.4
    - License: [Apache License 2.0](https://github.com/rook/rook/blob/v1.9.4/LICENSE)
    - Source: https://github.com/rook/rook/tree/v1.9.4/deploy/charts/rook-ceph
  - Container image(s)
    - quay.io/cephcsi/cephcsi:v3.6.1
      - License: [Apache License 2.0](https://github.com/ceph/ceph-csi/blob/v3.6.1/LICENSE)
    - k8s.gcr.io/sig-storage/csi-node-driver-registrar:v2.5.0
      - License: [Apache License 2.0](https://github.com/kubernetes-csi/node-driver-registrar/blob/v2.5.0/LICENSE)
    - k8s.gcr.io/sig-storage/csi-resizer:v1.4.0
      - License: [Apache License 2.0](https://github.com/kubernetes-csi/external-resizer/blob/v1.4.0/LICENSE)
    - k8s.gcr.io/sig-storage/csi-provisioner:v3.1.0
      - License: [Apache License 2.0](https://github.com/kubernetes-csi/external-provisioner/blob/v3.1.0/LICENSE)
    - k8s.gcr.io/sig-storage/csi-snapshotter:v5.0.1
      - License: [Apache License 2.0](https://github.com/kubernetes-csi/external-snapshotter/blob/v5.0.1/LICENSE)
    - k8s.gcr.io/sig-storage/csi-attacher:v3.4.0
      - License: [Apache License 2.0](https://github.com/kubernetes-csi/external-attacher/blob/v3.4.0/LICENSE)
    - quay.io/csiaddons/volumereplication-operator:v0.3.0
      - License: [Apache License 2.0](https://github.com/csi-addons/volume-replication-operator/blob/v0.3.0/LICENSE)

- ECK Operator
  - Helm chart
    - Version: 1.9.0
    - License: [Elastic License 2.0](https://github.com/elastic/cloud-on-k8s/blob/1.9.0/LICENSE.txt)
    - Source: https://github.com/elastic/cloud-on-k8s/tree/1.9.0/deploy/eck-operator
  - Container image(s)
    - docker.elastic.co/eck/eck-operator:1.9.0
      - License: [Elastic License 2.0](https://github.com/elastic/cloud-on-k8s/blob/1.9.0/LICENSE.txt)
    - docker.elastic.co/elasticsearch/elasticsearch:7.15.2
      - License: [Elastic License 2.0](https://github.com/elastic/elasticsearch/blob/v7.15.2/LICENSE.txt)
    - docker.elastic.co/kibana/kibana:7.15.2
      - License: [Elastic License 2.0](https://github.com/elastic/kibana/blob/v7.15.2/LICENSE.txt)
    - quay.io/prometheuscommunity/elasticsearch-exporter:v1.3.0
      - License: [Apache License 2.0](https://github.com/prometheus-community/elasticsearch_exporter/blob/v1.3.0/LICENSE)

- Grafana Operator
  - Helm chart: *None*
    - Version: 1.1.0
    - License: [Apache License 2.0](https://github.com/COPRS/infrastructure/blob/1.1.0-rc3/LICENSE)
  - Container image(s)
    - quay.io/grafana-operator/grafana-operator:v4.5.0
      - License: [Apache License 2.0](https://github.com/grafana-operator/grafana-operator/blob/v4.5.0/LICENSE)
    - docker.io/grafana/grafana:9.0.2-ubuntu
      - License: [GNU Affero General Public License v3.0](https://github.com/grafana/grafana/blob/v9.0.2/LICENSE)

- Kafka Operator
  - Helm chart:
    - Version: 0.27.1
    - Licence: [Apache License 2.0](https://github.com/strimzi/strimzi-kafka-operator/blob/0.27.1/LICENSE)
    - Source: https://github.com/strimzi/strimzi-kafka-operator/tree/0.27.1/helm-charts/helm3/strimzi-kafka-operator
  - Container image(s)
    - quay.io/strimzi/operator:0.27.1
      - License: [Apache License 2.0](https://github.com/strimzi/strimzi-kafka-operator/blob/0.27.1/LICENSE)
    - quay.io/strimzi/kafka:0.27.1-kafka-2.8.1
      - License: [Apache License 2.0](https://github.com/strimzi/strimzi-kafka-operator/blob/0.27.1/LICENSE)

- Prometheus Operator
  - Helm chart:
    - Version: 21.0.0
    - Licence: [Apache License 2.0](https://github.com/prometheus-community/helm-charts/blob/kube-prometheus-stack-21.0.0/LICENSE)
    - Source: https://github.com/prometheus-community/helm-charts/tree/kube-prometheus-stack-21.0.0/charts/kube-prometheus-stack
  - Container image(s)
    - quay.io/prometheus-operator/prometheus-operator:v0.52.1-amd64
      - License: [Apache License 2.0](https://github.com/prometheus-operator/prometheus-operator/blob/v0.52.1/LICENSE)
    - quay.io/prometheus-operator/prometheus-config-reloader:v0.52.1-amd64
      - License: [Apache License 2.0](https://github.com/prometheus-operator/prometheus-operator/blob/v0.52.1/LICENSE)
    - quay.io/prometheus/node-exporter:v1.3.0
      - License: [Apache License 2.0](https://github.com/prometheus/node_exporter/blob/v1.3.0/LICENSE)
    - quay.io/prometheus/alertmanager:v0.23.0
      - License: [Apache License 2.0](https://github.com/prometheus/alertmanager/blob/v0.23.0/LICENSE)
    - quay.io/prometheus/prometheus:v2.31.1
      - License: [Apache License 2.0](https://github.com/prometheus/prometheus/blob/v2.31.1/LICENSE)

- Stash Operator
  - Helm chart:
    - Version: v0.17.0
    - Licence: [AppsCode Community License 1.0.0](https://github.com/stashed/installer/blob/master/LICENSE.md)
    - Source: https://github.com/stashed/installer/tree/master/charts/stash-community
  - Container image(s)
    - docker.io/appscode/stash:v0.17.0
      - License: [AppsCode Community License 1.0.0](https://github.com/stashed/installer/blob/master/LICENSE.md)

- Fluent-bit 
  - Helm chart:
    - Version: 0.19.6
    - Licence: [Apache License 2.0](https://github.com/fluent/helm-charts/blob/fluent-bit-0.19.6/LICENSE)
    - Source: https://github.com/fluent/helm-charts/tree/fluent-bit-0.19.6/charts/fluent-bit
  - Container image(s)
    - docker.io/fluent/fluent-bit:1.9.3
      - License: [Apache License 2.0](https://github.com/fluent/fluent-bit/blob/v1.9.3/LICENSE)

- MongoDB
  - Helm chart:
    - Version: 11.0.3
    - Licence: [Apache License 2.0](https://github.com/bitnami/charts/blob/master/LICENSE.md)
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/mongodb
  - Container image(s)
    - docker.io/bitnami/mongodb:5.0.6-debian-10-r14
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)
    - docker.io/bitnami/mongodb-exporter:0.30.0-debian-10-r58
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)

- OpenLDAP
  - Helm chart: *None*
  - Container image(s)
    - docker.io/osixia/openldap:1.5.0
      - License: [MIT License](https://github.com/osixia/docker-openldap/blob/v1.5.0/LICENSE)

- PostreSQL
  - Helm chart:
    - Version: 11.0.2
    - Licence: [Apache License 2.0](https://github.com/bitnami/charts/blob/master/LICENSE.md)
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/postgresql
  - Container image(s)
    - docker.io/bitnami/postgresql:14.1.0-debian-10-r80
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)
    - quay.io/prometheuscommunity/postgres-exporter:v0.10.0
      - License: [Apache License 2.0](https://github.com/prometheus-community/postgres_exporter/blob/v0.10.0/LICENSE)

- Thanos
  - Helm chart:
    - Version: 8.3.0
    - Licence: [Apache License 2.0](https://github.com/bitnami/charts/blob/master/LICENSE.md)
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/thanos
  - Container image(s)
    - docker.io/bitnami/thanos:0.23.1-scratch-r3
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)

- Fluentd
  - Helm chart:
    - Version: 4.5.2
    - Licence: [Apache License 2.0](https://github.com/bitnami/charts/blob/master/LICENSE.md)
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/fluentd
  - Container image(s)
    - docker.io/bitnami/fluentd:1.14.2-debian-10-r23
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)

- Loki
  - Helm chart:
    - Version: 0.48.1
    - Licence: [Apache License 2.0](https://github.com/grafana/helm-charts/blob/loki-distributed-0.48.1/LICENSE)
    - Source: https://github.com/grafana/helm-charts/tree/loki-distributed-0.48.1/charts/loki-distributed
  - Container image(s)
    - docker.io/grafana/loki:2.5.0
      - License: [GNU Affero General Public License v3.0](https://github.com/grafana/loki/blob/v2.5.0/LICENSE)

- Apisix
  - Helm chart:
    - Version: 0.10.0
    - Licence: [Apache License 2.0](https://github.com/apache/apisix/blob/master/LICENSE)
    - Source: https://github.com/apache/apisix-helm-chart/tree/apisix-0.10.0/charts/apisix
  - Container image(s)
    - docker.io/apache/apisix:2.14.1-alpine
      - License: [Apache License 2.0](https://github.com/apache/apisix-docker/blob/master/LICENSE)
    - docker.io/apache/apisix-dashboard:2.11-alpine
      - License: [Apache License 2.0](https://github.com/apache/apisix-docker/blob/master/LICENSE)
    - docker.io/apache/apisix-ingress-controller:1.4.1
      - License: [Apache License 2.0](https://github.com/apache/apisix-docker/blob/master/LICENSE)
    - docker.io/bitnami/etcd:3.5.4-debian-11-r9
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)

- Falco
  - Helm chart:
    - Version: 1.16.2
    - Licence: [Apache License 2.0](https://github.com/falcosecurity/charts/blob/falco-1.16.2/LICENSE)
    - Source: https://github.com/falcosecurity/charts/tree/falco-1.16.2/falco
  - Container image(s)
    - docker.io/falcosecurity/falco:0.30.0
      - License: [Apache License 2.0](https://github.com/falcosecurity/falco/blob/0.30.0/COPYING)
    - docker.io/falcosecurity/falco-exporter:0.6.0
      - License: [Apache License 2.0](https://github.com/falcosecurity/falco-exporter/blob/v0.6.0/LICENSE)

- FinOps object storage exporter
  - Helm chart:
    - Version: 1.0.0
    - Licence: [Apache License 2.0](https://github.com/COPRS/monitoring/blob/1.0.0/finops/resources-exporter/LICENSE)
    - Source: https://github.com/COPRS/monitoring/tree/1.0.0/finops/resources-exporter/helm
  - Container image(s)
    - artifactory.coprs.esa-copernicus.eu/cs-docker/finops-object-storage-exporter:release-0.3.0
      - License: [Apache License 2.0](https://github.com/COPRS/monitoring/blob/1.0.0/finops/resources-exporter/LICENSE)

- FinOps resources exporter
  - Helm chart:
    - Version: 1.0.0
    - Licence: [Apache License 2.0](https://github.com/COPRS/monitoring/blob/1.0.0/finops/object-storage-exporter/LICENSE)
    - Source: https://github.com/COPRS/monitoring/tree/1.0.0/finops/object-storage-exporter/helm
  - Container image(s)
    - artifactory.coprs.esa-copernicus.eu/cs-docker/finops-resources-exporter:release-0.3.0
      - License: [Apache License 2.0](https://github.com/COPRS/monitoring/blob/1.0.0/finops/object-storage-exporter/LICENSE)

- Graylog
  - Helm chart:
    - Version: 1.9.2
    - Licence: [Apache License 2.0](https://github.com/KongZ/charts/blob/graylog-1.9.2/LICENSE)
    - Source: https://github.com/KongZ/charts/tree/graylog-1.9.2
  - Container image(s)
    - docker.io/graylog/graylog:4.3.3-1
      - License: [Apache License 2.0](https://github.com/Graylog2/graylog-docker/blob/4.3.3-1/LICENSE)

- Keycloack
  - Helm chart:
    - Version: 16.0.5
    - Licence: [Apache License 2.0](https://github.com/codecentric/helm-charts/blob/keycloak-16.0.5/LICENSE)
    - Source: https://github.com/codecentric/helm-charts/tree/keycloak-16.0.5/charts/keycloak
  - Container image(s)
    - docker.io/jboss/keycloak:15.0.2
      - License: [Apache License 2.0](https://github.com/keycloak/keycloak-containers/blob/15.0.2/License.html)

- Spring Cloud Data Flow
  - Helm chart:
    - Version: 7.0.1
    - Licence: [Apache License 2.0](https://github.com/bitnami/charts/blob/master/LICENSE.md)
    - Source: https://github.com/bitnami/charts/tree/master/bitnami/spring-cloud-dataflow
  - Container image(s)
    - docker.io/bitnami/spring-cloud-dataflow:2.9.4-debian-10-r7
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)
    - docker.io/bitnami/spring-cloud-skipper:2.8.4-debian-10-r6
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)
    - docker.io/bitnami/prometheus-rsocket-proxy:1.3.0-debian-10-r334
      - License: [Apache License 2.0](https://github.com/bitnami/containers/blob/main/LICENSE.md)

- Keda
  - Helm chart:
    - Version: 2.6.2
    - Licence: [Apache License 2.0](https://github.com/kedacore/charts/blob/v2.6.2/LICENSE)
    - Source: https://github.com/kedacore/charts/tree/v2.6.2/keda
  - Container image(s)
    - ghcr.io/kedacore/keda:2.6.1
      - License: [Apache License 2.0](https://github.com/kedacore/keda/blob/v2.6.1/LICENSE)
    - ghcr.io/kedacore/keda-metrics-apiserver:2.6.1
      - License: [Apache License 2.0](https://github.com/kedacore/keda/blob/v2.6.1/LICENSE)

- Prometheus blackbox exporter
  - Helm chart:
    - Version: 5.8.2
    - Licence: [Apache License 2.0](https://github.com/prometheus-community/helm-charts/blob/prometheus-blackbox-exporter-5.8.2/LICENSE)
    - Source: https://github.com/prometheus-community/helm-charts/tree/prometheus-blackbox-exporter-5.8.2/charts/prometheus-blackbox-exporter
  - Container image(s)
    - docker.io/prom/blackbox-exporter:v0.20.0
      - License: [Apache License 2.0](https://github.com/prometheus/blackbox_exporter/blob/v0.20.0/LICENSE)

- Autoscaling
  - Helm chart: *None*
  - Container image(s)
    - artifactory.coprs.esa-copernicus.eu/rs-docker/safescaled:v22.06.1
      - License: [Apache License 2.0](LICENSE)
    - artifactory.coprs.esa-copernicus.eu/rs-docker/cluster-autoscaler:1.22.3
      - License: [Apache License 2.0](LICENSE)
    - artifactory.coprs.esa-copernicus.eu/rs-docker/rs-infra-scaler:1.1.0
      - License: [Apache License 2.0](LICENSE)

- Kubelet rubber stamp
  - Helm chart: *None*
  - Container image(s)
    - docker.io/digitalocean/kubelet-rubber-stamp:v0.3.1-do.2
      - License: [Apache License 2.0](https://github.com/digitalocean/kubelet-rubber-stamp/blob/v0.3.1-do.2/LICENSE)

- Calico
  - Helm chart: *None*
  - Container image(s)
    - quay.io/calico/cni:v3.19.2
      - License: [Apache License 2.0](https://github.com/projectcalico/calico/blob/master/cni-plugin/LICENSE)

- ClamAV
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 0.103.5
  - License: [GNU General Public License v2.0](https://github.com/Cisco-Talos/clamav/blob/clamav-0.103.5/COPYING)

- CoreDNS
  - Helm chart: *None*
  - Container image(s)
    - k8s.gcr.io/coredns/coredns:1.8.0
      - License: [Apache License 2.0](https://github.com/coredns/coredns/blob/v1.8.0/LICENSE)

- Kubernetes
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: v1.22.2
  - License: [Apache License 2.0](https://github.com/kubernetes/kubernetes/blob/v1.22.2/LICENSE)

- Suricata
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 6.0.4
  - License: [GNU General Public License v2.0](https://github.com/OISF/suricata/blob/suricata-6.0.4/LICENSE)

- Wazuh
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 4.2.5-1
  - License: [GNU General Public License v2.0](https://github.com/wazuh/wazuh/blob/v4.2.5/LICENSE)

- Zookeeper
  - Helm chart: *None*
  - Container image(s)
    - k8s.gcr.io/kubernetes-zookeeper
      - Version: 3.5.9
      - License: [Apache License 2.0](https://github.com/apache/zookeeper/blob/release-3.5.9/LICENSE.txt)

- Safescale
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: v21.11.0
  - License: [Apache License 2.0](https://github.com/CS-SI/SafeScale/blob/v21.11.0/LICENSE)

- Rclone
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 1.57.0
  - License: [MIT License](https://github.com/rclone/rclone/blob/v1.57.0/COPYING)

- OpenVPN
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 2.4.7
  - License: [GNU General Public License v2.0](https://github.com/OpenVPN/openvpn/blob/v2.4.7/COPYRIGHT.GPL)

- Nmap
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 7.92-r2
  - License: [Nmap Public Source License Version 0.94](https://github.com/nmap/nmap/blob/master/LICENSE)
  
- Containerd
  - Helm chart: *None*
  - Container image(s): *None*
  - Version: 1.4.9-1
  - License: [Apache License 2.0](https://github.com/containerd/containerd/blob/v1.4.9/LICENSE)
