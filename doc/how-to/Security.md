# Managing security COTS on the cluster's nodes

The purpose of this document is to explain how to install and uninstall 
security COTS on nodes.

There is 5 COTS that are currently deployed in order to provide sucurity to the 
infrastructure. 

- AuditD
- ClamAv
- Wazuh
- Suricata
- OpenVPN

According to the purpose of each COTS, the node where the installation is perfomed change.

## Future improvements
Suricata: choose dynamically the interface to work on 
Wazuh: handle wazuh password with a vault 

## AuditD
**Scope: All**

Auditd is the userspace component to the Linux Auditing System. It's responsible for writing audit records to the disk.
The COTS can be configured by updating the file : ```infrastructure/platform/roles/security/auditd/defaults/main.yml```

| Name | Function | Required |
|------|----------|----------|
|  auditd_version    |    Version of auditd to use      |   Yes   |
|  rules   |   List of rules to add to auditd each rule is composed of two attributes the name and the rule description      |     No     |

In order to make easy the management of auditd logs, Laurel can be deployed.
This tool is plugged to audisp and allow to merge auditd message types in one unique JSON.
The documentation and exemple of configuration are available here : https://github.com/threathunters-io/laurel

This plugin allow you to set laurel.conf through `laurel_conf`
The audisp plugin can be configured by setting `laurel_audisp` property.
Finally the release to be installed can be set by `using laurel_url`.
NB if one property is set, all properties must be set too otherwise the installation will fail. 

## ClamAv
**Scope: gateways and masters**
The COTS can be configured by updating the file : ```infrastructure/platform/roles/security/clamav/defaults/main.yml`

ClamAv is an antivirus controlled by systemd.  
Rules bases are reloaded each days.  
The version of this cots can be edited in ```/platform/roles/security/clamav/default/main.yaml```
Two crons are created by default, you can change their name or hour of execution by updating the cots configuration.

| Name | Function | Required |
|------|----------|----------|
|  clamav_version   | Clamav version  |    Yes      |     
|  cron_update |  Object cron to update clamav rules each day | No  |
|  cron_scan |  Object cron to run the clamav scan each day | No  |

Be aware that the clamav cron use a lot of resources and can impact the stability of the plateform.
## Wazuh
**Scope: The Manager is installed only on the first master node and agent that are installed on all remaining nodes.** 
The COTS can be configured by updating the file : ```infrastructure/platform/roles/security/wazuh/defaults/main.yml```
We strongly advise to use the lastest version of Wazuh as the version 3 and 4 are not the same. Using version 3, the playbook may not work.

| Name | Function | Required |
|------|----------|----------|
| wazuh_repository    |    Repository of wazuh (can change according version)     |   Yes   |
| wazuh_version   | Wazuh version  |    Yes      |     
| wazuh_registration_password  |  The wazuh master password to join cluster surveillance  |  Yes |
| unused_rules  |  List of rules to not use  |  No |
| unused_decoders  |  List of decoders to not use  |  No |
| agent_conf  |  XML content describing agent configuration  |  Yes |
| log_level_output  |  Level log output that must appears in alerts.json (from 1 to 15)  |  No |
| ossec_conf  | Set Custom ossec.conf for master  |  No |


## Suricata
**Scope: Gateway**

Suricata is a NIDS that will reload rules every day to try to detect attack. 
The COTS can be configured by updating the file : ```infrastructure/platform/roles/security/suricata/defaults/main.yml```
As the Wazuh playbook, we strongly advise you to use the provided version in parameters as the install may not work otherwise.

| Name | Function | Required |
|------|----------|----------|
| suricata_version    |  The version of the package to be installed  |   Yes   |
| suricata_repo_version   | The version of the repository package to be installed  |    Yes      |     
| suricata_iface  |  The interface to monitor  |  Yes |
| suricata_local_cidr  |  Cidr to monitor  |  Yes |
| ignored_rules  |  List of rules to ignore  |  No |
| rotateIpReputation  |  Object that describe the ip reputation management  |  No |
| rotate_rules  |  Object that describe the rule rotation  |  No |
| categories | Object that describe the categories for ip reputation | No |

The purpose of the object `rotateIpReputation` is to create a bash file that is going to be executed once a day to update the ip reputation of suricata.
rotateIpReputation object is optional, it is composed of 3 attributes:
- The name that should end by .sh.
- The content, it is the actions that want to perform to update the ip reputation of suricata in bash.
- The cron paramater, as the name and the hour of execution each day.

The same object structure is available by updating `rotate_rules` object, here the purpose is to describe the bash file that is going to update the rules each day.

`categories` Object is composed of two attributes: 
- a name fir the file.
- the content of the file following `id,short name,description` format. 
It is used to manage categories for ip reputation. 

## OpenVPN 
**Scope: Gateway**

OpenVpn client installed on gateway.
The COTS can be configured by updating the file : ```infrastructure/platform/roles/security/openvpn/defaults/main.yml```

| Name | Function | Required |
|------|----------|----------|
| openvpn_version    |   OpenVPN version     |   Yes   |
| conf_files   | List of objects that describe configuration files for openvpn  |    No      |     

The `conf_files` is a list of file with th following structure : 
- A `name`.
- The `content` of the file. 
It can be used to add certificate and client configuration for the VPN client.
