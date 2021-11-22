# Purpose
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

# To Improve
Suricata => choose dynamically the interface to work on 
Wazuh => Find a way to handle wazuh master ip and pass it to agent during install
Wazuh => handle wazuh password with a vault 

## AuditD
**Scope: All**

Auditd is the userspace component to the Linux Auditing System. It's responsible for writing audit records to the disk.
Customs rules can be added in ```/platform/roles/security/auditd/files```

The version of this cots can be edited in ```/platform/roles/security/auditd/default/main.yaml``` 
-  ```auditd_version``` change to the wished version 

## ClamAv
**Scope: gateways and masters**

ClamAv is an antivirus controlled by systemd.  
Rules bases are reloaded each days.  
The version of this cots can be edited in ```/platform/roles/security/clamav/default/main.yaml``` 
-  ```clamav_version``` change to the wished version 

## Wazuh
**Scope: The Manager is installed only on the first master node and agent that are installed on all remaining nodes.**
Rules bases are reloaded each days. 

Some parameters needs to be edited in ```/platform/roles/security/wazuh/default/main.yaml``` 
- ```wazuh_repository``` the repositoy version to use
- ```wazuh_version``` the package version to install 
- ```wazuh_registration_file_path``` the path to autd.pass file that contains the password that allow to join the wazuh master
- ```wazuh_registration_password``` the wazuh master password to join cluster surveillance
- ```wazuh_ossec_conf``` the path to ossec.conf file
- ```wazuh_ossec_agent_conf```  the path to agent.conf file 
- ```wazuh_manager_ip``` the wazuh master ip

## Suricata
**Scope: Gateway**

Suricata is a NIDS that will reload rules every day to try to detect attack. 
Some parameters needs to be edited in ```/platform/roles/security/suricata/default/main.yaml``` 

- ```suricata_version``` is the version of the package to be installed
- ```suricata_repo_version``` is the version of the repository package to be installed
- ```suricata_iface``` is the interface to monitor
- ```suricata_local_cidr``` is the cidr to monitor
- ```suricata_path ``` is the install path

## OpenVPN 
**Scope: Gateway**

OpenVpn client installed on gateway.
The version of this cots can be edited in ```/platform/roles/security/openvpn/default/main.yaml``` 
-  ```openvpn_version``` change to the wished version 

# Update
To update, create a specific yaml file with COTS to uninstall
then update conf file and rereun install script on thoses nodes. 