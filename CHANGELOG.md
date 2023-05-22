# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

> Content of release :
>
>- **Added** for new features.
>- **Changed** for changes in existing functionality.
>- **Deprecated** for soon-to-be removed features.
>- **Removed** for now removed features.
>- **Fixed** for any bug fixes.
>- **Security** in case of vulnerabilities.

## [1.5.0-rc1] - 2023-01-04

### Changed

- Assign GrafanaAdmin role instead of Admin
- Allow deployment of custom remote and unsigned Grafana plugins
- Update Grafana operator to 4.8.0 and Grafana to 9.3.0 (both latest)
- Update SCDF to 2.9.6
- Update fluent-bit to 2.0.6

### Fixed

- [#607 - \[BUG\] Missing traces on Elastic Search](https://github.com/COPRS/rs-issues/issues/607)
- [#691 - \[BUG\] \[Infra\] SCDF dashboard access unsuccessful - apisix route](https://github.com/COPRS/rs-issues/issues/691)

## [1.4.0-rc1] - 2022-11-24

### Fixed

- [#551 - \[BUG\] rs-addon deployment - Error when updating existing additional resource](https://github.com/COPRS/rs-issues/issues/687)
- [#687 - \[BUG\] \[OPS\] Kafka information are not reacheable with promotheus](https://github.com/COPRS/rs-issues/issues/687)
- [#691 - \[BUG\] \[Infra\] SCDF dashboard access unsuccessful - apisix route](https://github.com/COPRS/rs-issues/issues/691)

## [1.3.0-rc1] - 2022-10-26

### Added

- Deploy Reference System on OVH cloud and local cluster

### Changed

- Improvements :
  - Prevent node autoscaling concurrence for better reliability
  - Reduce autoscaler's calls to safescale, greatly improving its performance
  - Merge falco and falco-exporter apps for better clarity
  - Deploy rs_addon additional resources using kubectl server side for better homogeneity
  - Limit app-installer's ansible templating to the app's root dir to allow the placement of any kind of file in the app's subfolders

### Fixed

- [#561 - \[BUG\] \[Infra\] Fluentbit - fluentd fails when topic_to_loki_regex setting is empty](https://github.com/COPRS/rs-issues/issues/561)

## [1.2.0-rc1] - 2022-09-28

### Added

- [#540 - \[STORAGE\] Procedure to copy file and directories on the RS shared disk](https://github.com/COPRS/rs-issues/issues/540)
- Documentation and Notice.md
  
## [1.1.0-rc3] - 2022-09-13

### Fixed

- Set scdf stream deployment namespace at playbook launch

## [1.1.0-rc2] - 2022-09-05

### Fixed

- [#341 - \[BUG\] \[Infra\] Keycloak configuration prevents Grafana logout](https://github.com/COPRS/rs-issues/issues/341)

## [1.1.0-rc1] - 2022-08-31

### Added

- Docs skeleton

### Changed

- [#504 - Allow SCDF to deploy in multiple namespaces](https://github.com/COPRS/rs-issues/issues/504)

### Fixed

- [#295 - \[BUG\] \[Infra\] Graylog - Falco dashboard not showing events](https://github.com/COPRS/rs-issues/issues/295)
- [#371 - \[BUG\] Grafana User Preferences removed](https://github.com/COPRS/rs-issues/issues/371)
- [#390 - \[BUG\] \[Infra\] Suricata log management on node - no log rotation implemented](https://github.com/COPRS/rs-issues/issues/390)
- [#443 - \[BUG\] \[Infra\] rs-addon deployment: scdf script cannot deploy streams if nodelocaldns is not set as gateway dns](https://github.com/COPRS/rs-issues/issues/443)
- [#452 - \[BUG\] \[Infra\] \[Deploy\] - Error when deploying on a tenant without any existing cluster](https://github.com/COPRS/rs-issues/issues/452)
- [#455 - \[BUG\] \[OPS\] log from lifecycle-eviction-job were not well managed by fluentbit](https://github.com/COPRS/rs-issues/issues/455)
- [#456 - \[BUG\] \[Infra\] rs-setup: missing dns configuration for the pip modules](https://github.com/COPRS/rs-issues/issues/456)
- [#464 - \[BUG\] \[Infra\] Kubernetes - Wrong DNS address is deployed in resolv.conf](https://github.com/COPRS/rs-issues/issues/464)
- [#494 - \[BUG\] \[Infra\] Prometheus pods stuck in init state](https://github.com/COPRS/rs-issues/issues/494)
- [#516 - \[BUG\] \[Infra\] prometheus: prometheus node-exporter are not deployed on rook-ceph nodes](https://github.com/COPRS/rs-issues/issues/516)

## [0.10.0-rc1] - 2022-08-03

### Fixed

- [#377 - \[BUG\] \[GRAFANA\] Every 10 minutes, Grafana dashboards are overwritten by github dashboards.](https://github.com/COPRS/rs-issues/issues/377)
- [#378 - \[BUG\] \[GRAFANA\] No backup for library panel](https://github.com/COPRS/rs-issues/issues/378)

## [0.9.0-rc1] - 2022-07-06

### Changed

- [#322 - \[Infra\] \[Doc\] Specify which files to input s3 and vault credentials](https://github.com/COPRS/rs-issues/issues/322)
- [#343 - Grafana improvement : alert by email feature](https://github.com/COPRS/rs-issues/issues/343)

### Fixed

- [#318 - \[BUG\] \[IVV\] \[Infra\] Requests to exposed services are blocked after 5min of inactivity](https://github.com/COPRS/rs-issues/issues/318)
- [#361 - \[BUG\] \[Infra\] Vault variable not optional - generated_inventory_vars.yaml not created](https://github.com/COPRS/rs-issues/issues/361)
- [#362 - \[BUG\] \[Documentation\] \[Infra\] README.md: Missing requirements for deployment](https://github.com/COPRS/rs-issues/issues/362)
- [#365 - \[BUG\] \[Infra\] Suricata version 0.6.4-0ubuntu1 is obsolete](https://github.com/COPRS/rs-issues/issues/365)
- [#367 - \[BUG\] \[Documentation\] \[Infra\] rs-addon playbook cannot use local zip files](https://github.com/COPRS/rs-issues/issues/367)
- [#379 - \[BUG\] \[HMI\] Unavailability of graylog](https://github.com/COPRS/rs-issues/issues/379)
- [#414 - \[BUG\] \[Infra\] gateways' ips are swapped in the hosts.yaml file](https://github.com/COPRS/rs-issues/issues/414)
- [#415 - \[BUG\] \[Documentation\] Cluster scaling misses a step in the deployment of the autoscaler on an existing cluster](https://github.com/COPRS/rs-issues/issues/415)
- [#416 - \[BUG\] \[Infra\] Auto-scaler: wrong node is initialized after expansion](https://github.com/COPRS/rs-issues/issues/416)
- [#417 - \[BUG\] \[Infra\] Auto-scaler: nodegroup expansion fails at security playbook](https://github.com/COPRS/rs-issues/issues/417)
- [#419 - \[BUG\] \[Infra\] Auto-scaler limitation: autoscaler should not be used for infra nodes with volumes](https://github.com/COPRS/rs-issues/issues/419)
- [#420 - \[BUG\] \[Infra\] falco priority class is named fluentbit-priority](https://github.com/COPRS/rs-issues/issues/420)
- [#421 - \[BUG\] \[Infra\] Obsolete charts: fluentd and thanos charts versions are no longer available](https://github.com/COPRS/rs-issues/issues/421)
- [#424 - \[BUG\] \[Infra\] Falco pods are OOM killed](https://github.com/COPRS/rs-issues/issues/424)
- [#440 - \[BUG\] \[Infra\] prometheus-blackbox-exporter: wrong chart version](https://github.com/COPRS/rs-issues/issues/440)

## [0.8.0-rc2] - 2022-06-14

### Fixed

- [#416 - \[BUG\] \[Infra\] Auto-scaler: wrong node is initialized after expansion](https://github.com/COPRS/rs-issues/issues/416)

## [0.8.0-rc1] - 2022-06-10

### Added

- [#342 - Update component to monitor End Point (Blackbox exporter)](https://github.com/COPRS/rs-issues/issues/342)
- [#357 - \[SCALER\] \[IMPLEMENTATION IN COPRS\] Implementation in RS Infrastructure](https://github.com/COPRS/rs-issues/issues/357)

### Changed

- [#393 - ICD update for RS Core Sentinel-2 chains](https://github.com/COPRS/rs-issues/issues/393)

### Fixed

- [#277 - \[BUG\] \[OPS\] Unable to create TLS certificates due to missing credentials](https://github.com/COPRS/rs-issues/issues/277)
- [#363 - \[BUG\] \[Infra\] hosts.ini template is not correctly formatted](https://github.com/COPRS/rs-issues/issues/363)
- [#374 - \[BUG\] \[Infra\] Keda: Missing CRD)](https://github.com/COPRS/rs-issues/issues/374)
- [#394 - \[BUG\] \[Infra\] RS-Addon : rs-addon deployment fails when looking for additional resources](https://github.com/COPRS/rs-issues/issues/394)

## [0.7.0-rc1] - 2022-05-11

### Changed

- [#347 - Update software component to deploy a RS addon / RS core on RS Platform with several DSL lines.](https://github.com/COPRS/rs-issues/issues/347)

### Fixed

- [#268 - \[BUG\] \[SECURITY\] \[IAM\] Remove double authentication for Graylog (for now, there is double authentication)](https://github.com/COPRS/rs-issues/issues/268)

## [0.6.0-rc1] - 2022-04-13

### Changed

- [#301 - Data stored into LOKI to be improved](https://github.com/COPRS/rs-issues/issues/301)

### Security

- [#294 - \[SECURITY\] Clamav - Deploy the daemon on RS Nodes accessible from internet](https://github.com/COPRS/rs-issues/issues/294)

## [0.5.0-rc1] - 2022-03-18

### Added

- [#221 - Software component to deploy a RS addon / RS core on RS Platform](https://github.com/COPRS/rs-issues/issues/221)
- [#232 - Procedure to uninstall a RS-core / RS-add-on](https://github.com/COPRS/rs-issues/issues/232)
- [#265 - \[SCALER\] \[IMPLEMENTATION\] horizontal scalling of pods](https://github.com/COPRS/rs-issues/issues/265)
- [#267 - \[SECURITY\] \[MONITOR\] \[OPENLDAP\] Provide monitoring for Openldap in Grafana](https://github.com/COPRS/rs-issues/issues/267)
- [#271 - \[SECURITY\] Connecting kubectl with Keycloak (oAuth2)](https://github.com/COPRS/rs-issues/issues/271)

## [0.4.0-rc3] - 2022-03-07

### Fixed

- [#287 - \[BUG\] \[Infra\] Apparmor prevents node-exporter from exporting utilisation metrics](https://github.com/COPRS/rs-issues/issues/287)

## [0.4.0-rc2] - 2022-02-22

### Added

- Add CONFIG.md
  
## [0.4.0-rc1] - 2022-02-21

### Changed

- [#234 - V1 / Technical debt](https://github.com/COPRS/rs-issues/issues/234)

### Fixed

- [#177 - \[BUG\] \[Documentation\] Installation manual of the infrastructure deployment (infrastructure readme)](https://github.com/COPRS/rs-issues/issues/177)

## [0.3.0-rc3] - 2022-02-18

### Fixed

- [#274 - \[BUG\] \[Infra\] Processing traces are not forwarded to elasticsearch](https://github.com/COPRS/rs-issues/issues/274)
- [#281 - \[BUG\] Several Grafana actions are forbidden](https://github.com/COPRS/rs-issues/issues/281)

## [0.3.0-rc2] - 2022-02-04

### Fixed

- [#251 - \[BUG\] \[Infra\] Elasticsearch app sample configuration has no coordinating node](https://github.com/COPRS/rs-issues/issues/251)

## [0.3.0-rc1] - 2021-12-15

### Added

- [#102 - \[Tradeoff\] Using ISTIO ?](https://github.com/COPRS/rs-issues/issues/102)
- [#125 - Backup and restore databases - ELASTICSEARCH](https://github.com/COPRS/rs-issues/issues/125)
- [#126 - SECURITY: Monitor specific scenarios through the SIEM](https://github.com/COPRS/rs-issues/issues/126)
- [#175 - FinOPS: monitor & control system costs - RESOURCES](https://github.com/COPRS/rs-issues/issues/175)
- [#178 - Compliance to "Non Functional Requirements"](https://github.com/COPRS/rs-issues/issues/178)
- [#179 - Security: deploy, configure FALCO and link logs to the SIEM](https://github.com/COPRS/rs-issues/issues/179)
- [#185 - Backup and Restore databases - POSTGRESQL](https://github.com/COPRS/rs-issues/issues/185)
- [#186 - Backup and Restore databases - LDAP](https://github.com/COPRS/rs-issues/issues/186)
- [#187 - FinOPS: monitor & control system costs - STORAGE](https://github.com/COPRS/rs-issues/issues/187)
- [#188 - FinOPS: monitor & control system costs - NETWORK](https://github.com/COPRS/rs-issues/issues/188)
- [#189 - Ingress Controller](https://github.com/COPRS/rs-issues/issues/189)
  
## [0.2.0-rc1] - 2021-12-15

### Added

- [#79 - SECURITY: live demonstration of FALCO](https://github.com/COPRS/rs-issues/issues/79)
- [#117 - MONITORING: Process new format of traces](https://github.com/COPRS/rs-issues/issues/117)
- [#119 - MONITORING: deploy LOKI](https://github.com/COPRS/rs-issues/issues/119)
- [#120 - Security : List of trusted sources](https://github.com/COPRS/rs-issues/issues/120)
- [#121 - INFRA/SECURITY: Deploy the new IAM solution](https://github.com/COPRS/rs-issues/issues/121)
- [#122 - SECURITY: Deploy the secret management solution](https://github.com/COPRS/rs-issues/issues/122)
- [#123 - INFRA: Provide a base image to use for future developments](https://github.com/COPRS/rs-issues/issues/123)
- [#124 - SECURITY: Provide a security clearance after deployment](https://github.com/COPRS/rs-issues/issues/124)
- [#145 - MONITORING: Deploy the metrics, log and trace chain](https://github.com/COPRS/rs-issues/issues/145)

### Fixed

- [#154 - \[Infra\] Second master fails to join cluster during Kubernetes deployment](https://github.com/COPRS/rs-issues/issues/154)

## [0.1.0-rc1] - 2021-11-17

### Added

- [#53 - SECURITY: Tradeoff on a new IAM solution](https://github.com/COPRS/rs-issues/issues/53)
- [#55 - SECURITY: Tradeoff on secret management](https://github.com/COPRS/rs-issues/issues/55)
- [#77 - INFRA: Create and deploy a new template image of Operating System](https://github.com/COPRS/rs-issues/issues/77)
- [#83 - Create Ansible playbooks to deploy Kubernetes](https://github.com/COPRS/rs-issues/issues/83)
- [#84 - Deploy the latest version of databases](https://github.com/COPRS/rs-issues/issues/84)
- [#87 - Create an Ansible playbook to manage my Operating System services](https://github.com/COPRS/rs-issues/issues/87)
- [#89 - Deploy Sprint Cloud Data Flow](https://github.com/COPRS/rs-issues/issues/89)
- [#90 - Deploy the latest version of Kafka](https://github.com/COPRS/rs-issues/issues/90)
- [#91 - Deploy Rook/CephFS into the cluster](https://github.com/COPRS/rs-issues/issues/91)
- [#97 - Deploy infrastructure](https://github.com/COPRS/rs-issues/issues/97)
