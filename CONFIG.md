# Default Configuration

> Delivered configuration corresponds to a platform that can be used without processing chains. It covers the following points :
>
> - definition of system resources required for services
> - definition of sufficient storage for services
> - high availability of services
>
> When adding an RS-ADDON / RS-CORE, the following COTS will need to be re-configured:
>
> - Storage size for Elasticsearch processing
> - Storage size for Elasticsearch security
> - Storage size for Kafka
> - Storage size for Loki
> 
> Concerning Rook Ceph cluster : a dedicated cluster should be created for an RS-ADDON if it needs one

## Prerequisite

- Reservation of 2 IPs for egress
- Platform with 21 nodes :
  - 2 Gateway (VM : 4 CPUs / 8 Go RAM)
  - 3 Master (VM : 4 CPUs / 8 Go RAM)
  - 12 Worker (VM : 4 CPUs / 16 Go RAM)
  - 2 Specific Worker for Prometheus  (VM : 8 CPUs / 32 Go RAM)
  - 2 Egress (VM : 4 CPUs / 16 Go RAM)

## Configuration of COTS

|   |   |
| - | - |
| **Apisix** | - Namespace : networking <br> ------------------------ <br> **apisix** <br> - QoS : Burstable <br> - Replicas : 3 <br> ------------------------ <br> **dashboard** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **etcd** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Persistent Volume : <br> &nbsp;&nbsp;&nbsp; - Size : 8Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> ------------------------ <br> **ingress controller** <br> - QoS : Burstable <br> - Replicas : 2 |
| **Calico** | - Namespace : kube-system <br> ------------------------ <br> **calico** <br> - QoS : Burstable <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 300m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 500M <br> &nbsp;&nbsp;&nbsp; - Request CPU : 150m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64M <br> ------------------------ <br> **kube-controller** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256M <br> &nbsp;&nbsp;&nbsp; - Request CPU : 30m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64M |
| **Cert-manager** | - Namespace : infra <br> ------------------------ <br> **cert-manager** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **cainjector** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **webhook** <br> - QoS : Burstable <br> - Replicas : 1 |
| **CoreDNS** | - Namespace : kube-system <br> ------------------------ <br> **coredns** <br> - QoS : Burstable <br> - Replicas : 2 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 170Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 70Mi <br> ------------------------ <br> **localdns** <br> - QoS: Burstable <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 170Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 70Mi <br> ------------------------ <br> **dns autoscaler** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 20m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 10Mi |
| **Elastic Operator** | - Namespace : infra <br> - QoS : Burstable <br> - Replicas : 1 <br>- Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **Elasticsearch processing** | - Namespace : database <br> - Priority Class : 900000 <br> ------------------------ <br> **coordinating** <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume : <br> &nbsp;&nbsp;&nbsp; - Size : 1Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 6Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 4Gi <br> ------------------------ <br> **data** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Persistent Volume : <br> &nbsp;&nbsp;&nbsp; - Size : 100Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 2 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 8Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 4Gi <br> ------------------------ <br> **exporter** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **Elasticsearch security**   | - Namespace : security <br> - Priority Class : 900000 <br> ------------------------ <br> **coordinating** <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume : <br> &nbsp;&nbsp;&nbsp; - Size : 1Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 400m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 6Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 4Gi <br> ------------------------ <br> **data** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Persistent Volume : <br> &nbsp;&nbsp;&nbsp; - Size : 100Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 2 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 8Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 4Gi <br> ------------------------ <br> **exporter** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **Etcd** | - Namespace : kube-system <br> - QoS : Burstable <br> - Replicas: 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 100Mi |
| **Falco** | - Namespace : security <br> - QoS : Burstable <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi |  
| **Falco Exporter** | - Namespace : security <br> - QoS : Burstable <br> - Replicas : Daemonset |  
| **Fluentbit** | - Namespace : logging <br> - QoS : Burstable <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 512Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 20m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **Fluentd** | - Namespace : logging <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 10Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 512Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 256Mi |
| **Finops Object Storage Exporter** | - Namespace : monitoring <br> - QoS : BestEffort <br> - Replicas : 1 |
| **Finops Ressources Exporter** | - Namespace : monitoring <br> - QoS : Burstable <br> - Replicas : 1 |
| **Grafana** | - Namespace : monitoring <br> ------------------------ <br> **grafana** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 8Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 200m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 256Mi <br> ----------------------------- <br>  - QoS : Burstable <br> - Replicas : 1 <br> ------ <br> **manager** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 400m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 512Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 256Mi <br> ------ <br> **kube-rbac-proxy** | 
| **Graylog** | - Namespace : security <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 10Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 1Gi | 
| **Kafka Cluster** | - Namespace : infra <br> ------------------------ <br> **operator** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 256Mi <br> ------------------------ <br> **kafka** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 200Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 4Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 250m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 2Gi <br> ------------------------ <br> - QoS : Burstable <br> - Replicas : 1 <br> ------ <br> **topic-operator** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi  <br> ------ <br> **user-operator** <br> ------ <br> **tls-sidecar** <br> ------------------------ <br> **exporter** <br>  - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi <br> ------------------------ <br> **zookeeper** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 50Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 256Mi |
| **Keycloak** | - Namespace : iam <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 250m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi |
| **Kibana processing** | - Namespace : database <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp;  - Request Memory : 1Gi |
| **Kibana security** | - Namespace : security <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp;  - Request Memory : 1Gi |
| **Kubernetes** | - Namespace : kube-system <br> ------------------------ <br> **kube-apiserver** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 250m <br> ------------------------ <br> **kube-controller-manager** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 200m <br> ------------------------ <br> **kube-scheduler** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> ------------------------ <br> **nginx-proxy** <br> - QoS : Burstable <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Request CPU : 25m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 32m | 
| **Linkerd** | - Namespace : networking <br> ------------------------ <br> **linkerd-cni** <br> - QoS : BestEffort <br> - Replicas : Daemonset <br> ------------------------ <br> **linkerd-destination** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 250Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 50Mi <br> ------------------------ <br> **linkerd-identity** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 250Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 10Mi <br> ------------------------ <br> **linkerd-proxy-injector** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 250Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 50Mi <br> ------------------------ <br> **metrics-api** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **tap** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **tap-injector** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **web** <br> - QoS : Burstable <br> - Replicas : 1 <br> ------------------------ <br> **linkerd-proxy** <br> - Replicas : Sidecar (several namespaces) <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 250Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 20Mi |
| **Loki** | - Namespace : monitoring <br> - QoS : Burstable <br> - Replicas : 1 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 100Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **MongoDB** | - Namespace : database <br> ----------------------------- <br> - QoS : Burstable <br> - Replicas : 3  <br> ------ <br> **mongodb** <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 30Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 300m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ------ <br> **metrics** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 128Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 20m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64Mi <br> ----------------------------- <br> **arbiter** <br> - QoS : Burstable <br> - Replicas : 1 | 
| **Nmap** | - Namespace : security <br> - QoS : BestEffort <br> - Replicas : CronJob |
| **OpenLDAP** | - Namespace : iam <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 8Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **PostgreSQL** | - Namespace : database <br> ----------------------------- <br> **primary** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 30Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 250m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ----------------------------- <br> **read** <br> - QoS : Burstable <br> - Replicas : 2 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 30Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 250m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi |
| **PostgreSQL Exporter** | - Namespace : database <br> ----------------------------- <br> **Keycloak Exporter** <br> - QoS : Burstable <br> - Replicas : 1 <br> ----------------------------- <br> **Spring Cloud Data Flow Exporter** <br> - QoS : Burstable <br> - Replicas : 1 <br> ----------------------------- <br> **Spring Cloud Data Flow Skipper Exporter** <br> - QoS : Burstable <br> - Replicas : 1 |
| **Prometheus Operator** | - Namespace : infra <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |
| **Prometheus Stack** | - Namespace : logging <br> - Priority Class : 1000000 <br> ----------------------------- <br> - QoS : Burstable <br> - Replicas : 1 <br> ------ <br> **alertmanager** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 128Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64Mi <br> ----- <br> **config-reloader** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 64Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64Mi <br> ----------------------------- <br> - QoS : Burstable <br> - Replicas : 2 <br> ------ <br> **prometheus** <br>- Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 150Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 4 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 20Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 2 <br> &nbsp;&nbsp;&nbsp; - Request Memory : 12Gi <br> ---- <br> **config-reloader** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 64Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 10m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64Mi <br> ----------------------------- <br> **kube-state-metrics** <br> - QoS : Burstable <br> - Replicas : 1 <br> ---------------------------- <br> **node-exporter** <br> - QoS : BestEffort <br> - Replicas : Daemonset <br> ---------------------------- <br> **thanos-sidecar** <br> - Replicas : Sidecar <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 25m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ---------------------------- <br> **thanos-compactor** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Persistent Volume :  <br> &nbsp;&nbsp;&nbsp; - Size : 100Gi <br> &nbsp;&nbsp;&nbsp; - Access Mode : ReadWriteOnce <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ---------------------------- <br> **thanos-query** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ---------------------------- <br> **thanos-storegateway** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi |
| **Rclone** | - Namespace : security <br> - QoS : BestEffort <br> - Replicas : CronJob |
| **Rook Ceph** | - Namespace : rook-ceph <br> - Priority Class : 10000000 <br> - Ceph Block Pools Replicated : 3 <br> ---------------------------- <br> **operator** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi <br> ---------------------------- <br> - QoS : Burstable <br> - Replicas : 2 <br> ------ <br> **mng** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 1Gi <br> ------<br> **watch-active** <br> ---------------------------- <br> **mon** <br> - QoS : Burstable <br> - Replicas : 3 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 1Gi <br> ---------------------------- <br> **ceph-crash** <br> - QoS : Burstable <br> - Replicas : Daemonset <br> ---------------------------- <br> **osd** <br> - QoS : Guaranteed <br> - Replicas : Daemonset <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 2Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Request Memory : 2Gi  <br> ---------------------------- <br> **rook-ceph-tools** <br> - QoS : Burstable <br> - Replicas : 1 | 
| **Spring Cloud Data Flow** | - Namespace : processing <br> ---------------------------- <br> **server** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ---------------------------- <br> **skipper** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 1 <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 1Gi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 512Mi <br> ---------------------------- <br> **prometheus-proxy** <br> - QoS : Burstable <br> - Replicas : 1 <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 500m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 256Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi |  
| **Stash** | - Namespace : infra <br> - QoS : Burstable <br> - Replicas : 1 <br> ------ <br> **operator** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 512Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 128Mi  <br> ------ <br> **pushgateway** <br> - Ressources : <br> &nbsp;&nbsp;&nbsp; - Limits CPU : 100m <br> &nbsp;&nbsp;&nbsp; - Limits Memory : 128Mi <br> &nbsp;&nbsp;&nbsp; - Request CPU : 50m <br> &nbsp;&nbsp;&nbsp; - Request Memory : 64Mi |

## Exposed services

| Exposed service | URL subdomain | URL subpath |
| --------------- | ------------- | ----------- |
| apisix | apisix | /* |
| kube-apiserver | kube | /* |
| linkerd | linkerd | /* |
| grafana | monitoring | /* |
| prometheus | monitoring | /prometheus |
| thanos | monitoring | /thanos |
| kibana processing | processing | /kibana |
| spring cloud dataflow | processing | /* |
| kibana security | security | /kibana |
| graylog | security | /* |
| keycloack | iam | /* |

## Predefined groups and users

### Groups

| name | Services |
| ---- | -------- |
| networking | linkerd |
| operator | grafana (admin) \| prometheus \| kibana processing \| spring cloud dataflow |
| {{ keycloak.realm.name \| lower }}-admin | keycloack (admin realm console) |
| {{ keycloak.realm.name \| lower }}-user | keycloack (user console) |
| security | graylog \| kibana security |
| sudo |  (sudoer for ssh) |

### Users

{{ keycloak.realm.name | lower }}-admin => ALL GROUPS

## ETL

| Source (Logs) | Topics (Kafka) | Consumer | Nb consumer | Destination |
| ------------- | -------------- | -------- | ----------- | ----------- |
| /var/log/containers/\*.log (excluded : /var/log/containers/\*fluent\*.log) <br> **All logs** | fluentbit.processing | fluentd | 2 | Loki |
| /var/log/containers/\*.log (excluded : /var/log/containers/\*fluent\*.log) <br> **Only log JSON contains : `header.type: REPORT`** | fluentbit.trace | fluentd | 2 | Elasticsearch Processing |
| /var/log/syslog | fluentbit.system | fluentd <br> graylog | 2 <br> 2 | Loki <br> Elasticsearch Security |
| /var/log/containers/\*_kube-system_\*.log | fluentbit.docker_security | graylog | 2 | Elasticsearch Security |
| /var/log/containers/keycloak-?_iam_keycloak-\*.log | fluentbit.keycloak | graylog | 2 | Elasticsearch Security |
| /var/log/containers/apisix-\*_networking_apisix-\*.log | fluentbit.ingress | graylog | 2 | Elasticsearch Security |
| /var/ossec/logs/alerts/alerts.json | fluentbit.wazuh | graylog | 2 | Elasticsearch Security |
| /var/log/audit_\*.log | fluentbit.auditd | graylog | 2 | Elasticsearch Security |
| /var/log/containers/falco-?????_security_falco\*.log | fluentbit.falco | graylog | 2 | Elasticsearch Security |
| /var/log/containers/nmap-job-\*.log | fluentbit.scans | graylog | 2 | Elasticsearch Security |

## Kafka Topics

| Topic | Partition | Replication Factor | Segment Size (bytes) | Retention (time in ms) | Retention (size in bytes) | Cleanup Policy | Min Insync Replicas | Unclean Leader Election Enabled |
| -- | :--: | :--: | :--: | :--: | :--: | :--: | :--: | :--: |
| fluentbit.processing | 6 | 3 | 715827882 | | 2863311530 | delete | 2 | false |
| fluentbit.trace | 6 | 3 | 89478485 | 604800000 | | delete | 2 | false |
| fluentbit.system | 6 | 3 | 89478485 | | 357913941 | delete | 2 | false |
| fluentbit.docker_security | 6 | 3 | 178956970 | | 715827882 | delete | 2 | false |
| fluentbit.wazuh | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |
| fluentbit.auditd | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |
| fluentbit.falco | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |
| fluentbit.scans | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |
| fluentbit.ingress | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |
| fluentbit.keycloak | 6 | 3 | 1789956970 | | 715827882 | delete | 2 | false |

## Logs and metrics retention

| Destination | Retention | Infos |
| ----------- | --------- | ----- |
| Loki | 1460h (~60d) | ||
| Prometheus | 2d | metrics are saved in S3 via Thanos |
| Thanos S3 bucket | 30d | 5m resolution |
| Thanos S3 bucket | 10 years | 1h resolution |
| Elasticsearch processing | lifetime retention | | 
| Elasticsearch security | 6 months | |
