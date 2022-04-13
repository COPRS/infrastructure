# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/) and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

> Content of release :
> - **Added** for new features.
> - **Changed** for changes in existing functionality.
> - **Deprecated** for soon-to-be removed features.
> - **Removed** for now removed features.
> - **Fixed** for any bug fixes.
> - **Security** in case of vulnerabilities.

## [0.6.0-rc1] - 2022-04-13
### Changed
- [#301 - Data stored into LOKI to be improved](https://github.com/COPRS/SCRUM-Tickets/issues/301)
### Security
- [#294 - [SECURITY] Clamav - Deploy the daemon on RS Nodes accessible from internet](https://github.com/COPRS/SCRUM-Tickets/issues/294)

## [0.5.0-rc1] - 2022-03-18
### Added
- [#221 - Software component to deploy a RS addon / RS core on RS Platform](https://github.com/COPRS/SCRUM-Tickets/issues/221)
- [#232 - Procedure to uninstall a RS-core / RS-add-on](https://github.com/COPRS/SCRUM-Tickets/issues/232)
- [#265 - [SCALER][IMPLEMENTATION] horizontal scalling of pods](https://github.com/COPRS/SCRUM-Tickets/issues/265)
- [#267 - [SECURITY][MONITOR][OPENLDAP] Provide monitoring for Openldap in Grafana](https://github.com/COPRS/SCRUM-Tickets/issues/267)
- [#271 - [SECURITY] Connecting kubectl with Keycloak (oAuth2)](https://github.com/COPRS/SCRUM-Tickets/issues/271)

## [0.4.0-rc3] - 2022-03-07
### Fixed
- [#287 - [BUG] [Infra] Apparmor prevents node-exporter from exporting utilisation metrics](https://github.com/COPRS/SCRUM-Tickets/issues/287)

## [0.4.0-rc2] - 2022-02-22
### Added
- Add CONFIG.md
  
## [0.4.0-rc1] - 2022-02-21
### Changed
- [#234 - V1 / Technical debt](https://github.com/COPRS/SCRUM-Tickets/issues/234)
### Fixed
- [#177 - [BUG][Documentation] Installation manual of the infrastructure deployment (infrastructure readme)](https://github.com/COPRS/SCRUM-Tickets/issues/177)

## [0.3.0-rc3] - 2022-02-18
### Fixed
- [#274 - [BUG] [Infra] Processing traces are not forwarded to elasticsearch](https://github.com/COPRS/SCRUM-Tickets/issues/274)
- [#281 - [BUG] Several Grafana actions are forbidden](https://github.com/COPRS/SCRUM-Tickets/issues/281)

## [0.3.0-rc2] - 2022-02-04
### Fixed
- [#251 - [BUG] [Infra] Elasticsearch app sample configuration has no coordinating node](https://github.com/COPRS/SCRUM-Tickets/issues/251)

## [0.3.0-rc1] - 2021-12-15
### Added
- [#102 - [Tradeoff] Using ISTIO ?](https://github.com/COPRS/SCRUM-Tickets/issues/102)
- [#125 - Backup and restore databases - ELASTICSEARCH](https://github.com/COPRS/SCRUM-Tickets/issues/125)
- [#126 - SECURITY: Monitor specific scenarios through the SIEM](https://github.com/COPRS/SCRUM-Tickets/issues/126)
- [#175 - FinOPS: monitor & control system costs - RESOURCES](https://github.com/COPRS/SCRUM-Tickets/issues/175)
- [#178 - Compliance to "Non Functional Requirements"](https://github.com/COPRS/SCRUM-Tickets/issues/178)
- [#179 - Security: deploy, configure FALCO and link logs to the SIEM](https://github.com/COPRS/SCRUM-Tickets/issues/179)
- [#185 - Backup and Restore databases - POSTGRESQL](https://github.com/COPRS/SCRUM-Tickets/issues/185)
- [#186 - Backup and Restore databases - LDAP](https://github.com/COPRS/SCRUM-Tickets/issues/186)
- [#187 - FinOPS: monitor & control system costs - STORAGE](https://github.com/COPRS/SCRUM-Tickets/issues/187)
- [#188 - FinOPS: monitor & control system costs - NETWORK](https://github.com/COPRS/SCRUM-Tickets/issues/188)
- [#189 - Ingress Controller](https://github.com/COPRS/SCRUM-Tickets/issues/189)
  
## [0.2.0-rc1] - 2021-12-15
### Added
- [#79 - SECURITY: live demonstration of FALCO](https://github.com/COPRS/SCRUM-Tickets/issues/79)
- [#117 - MONITORING: Process new format of traces](https://github.com/COPRS/SCRUM-Tickets/issues/117)
- [#119 - MONITORING: deploy LOKI](https://github.com/COPRS/SCRUM-Tickets/issues/119)
- [#120 - Security : List of trusted sources](https://github.com/COPRS/SCRUM-Tickets/issues/120)
- [#121 - INFRA/SECURITY: Deploy the new IAM solution](https://github.com/COPRS/SCRUM-Tickets/issues/121)
- [#122 - SECURITY: Deploy the secret management solution](https://github.com/COPRS/SCRUM-Tickets/issues/122)
- [#123 - INFRA: Provide a base image to use for future developments](https://github.com/COPRS/SCRUM-Tickets/issues/123)
- [#124 - SECURITY: Provide a security clearance after deployment](https://github.com/COPRS/SCRUM-Tickets/issues/124)
- [#145 - MONITORING: Deploy the metrics, log and trace chain](https://github.com/COPRS/SCRUM-Tickets/issues/145)

### Fixed
- [#154 - [Infra] Second master fails to join cluster during Kubernetes deployment](https://github.com/COPRS/SCRUM-Tickets/issues/154)

## [0.1.0-rc1] - 2021-11-17
### Added
- [#53 - SECURITY: Tradeoff on a new IAM solution](https://github.com/COPRS/SCRUM-Tickets/issues/53)
- [#55 - SECURITY: Tradeoff on secret management](https://github.com/COPRS/SCRUM-Tickets/issues/55)
- [#77 - INFRA: Create and deploy a new template image of Operating System](https://github.com/COPRS/SCRUM-Tickets/issues/77)
- [#83 - Create Ansible playbooks to deploy Kubernetes](https://github.com/COPRS/SCRUM-Tickets/issues/83)
- [#84 - Deploy the latest version of databases](https://github.com/COPRS/SCRUM-Tickets/issues/84)
- [#87 - Create an Ansible playbook to manage my Operating System services](https://github.com/COPRS/SCRUM-Tickets/issues/87)
- [#89 - Deploy Sprint Cloud Data Flow](https://github.com/COPRS/SCRUM-Tickets/issues/89)
- [#90 - Deploy the latest version of Kafka](https://github.com/COPRS/SCRUM-Tickets/issues/90)
- [#91 - Deploy Rook/CephFS into the cluster](https://github.com/COPRS/SCRUM-Tickets/issues/91)
- [#97 - Deploy infrastructure](https://github.com/COPRS/SCRUM-Tickets/issues/97)
