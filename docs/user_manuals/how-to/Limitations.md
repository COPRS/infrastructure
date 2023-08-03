# Known issues and limitations

- [Known issues and limitations](#known-issues-and-limitations)
  - [1. Unable to use SCDF Undeploy/Deploy function without misconfiguration](#1-unable-to-use-scdf-undeploydeploy-function-without-misconfiguration)
    - [Issue](#issue)
    - [Workaround](#workaround)
  - [2. Failed to generate hosts.yaml](#2-failed-to-generate-hostsyaml)
    - [Issue](#issue-1)
    - [Workaround](#workaround-1)
  - [3. Impossible to add a new node to the cluster](#3-impossible-to-add-a-new-node-to-the-cluster)
    - [Issue](#issue-2)
    - [Workaround](#workaround-2)
  - [4. Asterisk (\*) in SCDF `stream-parameters.properties` causes random result](#4-asterisk--in-scdf-stream-parametersproperties-causes-random-result)
    - [Issue](#issue-3)
    - [Workaround](#workaround-3)
  - [5. Additional egress node creation fails](#5-additional-egress-node-creation-fails)
    - [Issue](#issue-4)
    - [Workaround](#workaround-4)
  - [6. Fluent Bit failed to start after a node restart](#6-fluent-bit-failed-to-start-after-a-node-restart)
    - [Issue](#issue-5)
    - [Workaround](#workaround-5)
  - [7. Manually bypass gateway for object storage traffic](#7-manually-bypass-gateway-for-object-storage-traffic)
    - [Issue](#issue-6)
    - [Workaround](#workaround-6)
  - [8. SCDF : First known tag is always used if not explicitly specified](#8-scdf--first-known-tag-is-always-used-if-not-explicitly-specified)
    - [Issue](#issue-7)
    - [Workaround](#workaround-7)

## 1. Unable to use SCDF Undeploy/Deploy function without misconfiguration

Tickets :

- [COPRS/rs-issues/#716](https://github.com/COPRS/rs-issues/issues/716)
- [spring-cloud/spring-cloud-dataflow/issues/5145](https://github.com/spring-cloud/spring-cloud-dataflow/issues/5145)

### Issue

It prevents the usage of the Undeploy/Deploy functions from SCDF for any stream that contains regular expression. For e.g. we had the following regex:

`ingestion-trigger.polling.inbox1.matchRegex=^S1.*(AUX_|AMH_|AMV_|MPL_ORB).*$`

But if you undeploy and deploy again the stream, even without editing this regex, it would become (notice the **addition** of `'` at the beginning and at the end) :

`ingestion-trigger.polling.inbox1.matchRegex='^S1.*(AUX_|AMH_|AMV_|MPL_ORB).*$'`

This causes unexpected behaviour in the java application using that regex as it is not the same anymore, thus not filtering the way it should.

### Workaround

You have to destroy the stream, edit the stream's property and redeploy the stream.

1. Simply destroy the stream in the SCDF HMI
2. Edit the file `stream-parameters.properties` of the stream
3. Zip the stream
4. Install the stream using the [how-to](/docs/user_manuals/how-to/RS%20Add-on%20-%20RS%20Core.md)

## 2. Failed to generate hosts.yaml

Ticket : [COPRS/rs-issues/#835](https://github.com/COPRS/rs-issues/issues/835)

### Issue

[The step 5 of the infrastructure's quick start](/README.md#5-generate-or-download-the-inventory-variables) might fail due to invalid `YAML` files. It's most probably because of bad indentation from the configuration done in the previous steps.

### Workaround

Fix the file `inventory/host_vars/setup/main.yaml`. The `YAML` syntax could be incorrect or a non-ascii character might be there.

1. Check the `YAML` structure

   - Install the package `yq` : <https://github.com/mikefarah/yq#install>
   - Check the syntax of `inventory/sample/host_vars/setup/main.yaml` with `yq` :
     - `yq inventory/sample/host_vars/setup/main.yaml`

   The file's content should be displayed in the terminal. If not, it means the syntax is not correct.

2. Check for non-ascii character

   `grep --color='auto' -P -n "[^\x00-\x7F]" inventory/host_vars/setup/main.yaml`

## 3. Impossible to add a new node to the cluster

Ticket : [COPRS/rs-issues/#859](https://github.com/COPRS/rs-issues/issues/859)

### Issue

Kubespray requires version `1.4.9-1` but the new nodes are deployed with a more recent version. As a result, the playbook fails and the new node is not fully deployed and configured.

### Workaround

After the step 2 *Install requirements* from the installation manual, edit the file `infrastructure/collections/kubespray/roles/container-engine/containerd/tasks/main.yml` from line 100.

Change from :

```yaml
- name: ensure containerd packages are installed
  package:
    name: "{{ containerd_package_info.pkgs }}"
    state: present
```

to :

```yaml
- name: ensure containerd packages are installed
  package:
    name: "{{ containerd_package_info.pkgs }}"
    force: true
    state: present
```

## 4. Asterisk (*) in SCDF `stream-parameters.properties` causes random result

Ticket : [COPRS/rs-issues/#902](https://github.com/COPRS/rs-issues/issues/902)

### Issue

If one parameter is declared twice in the SCDF configuration file `stream-parameters.properties` for e.g. :

```yaml
deployer.*.kubernetes.requests.memory=512Mi
deployer.message-filter.kubernetes.requests.memory=1024Mi
```

It will cause SCDF to pick randomly one or the other value for the `message-filter.kubernetes.requests.memory` parameter instead of the desired one.

### Workaround

If you need to change a parameter's value for one component, you need to explicitly set the parameter for all components.

Instead of this :

```yaml
deployer.*.kubernetes.requests.memory=512Mi
deployer.message-filter.kubernetes.requests.memory=1024Mi
```

Use this :

```yaml
deployer.app1.kubernetes.requests.memory=512Mi
deployer.app2.kubernetes.requests.memory=512Mi
deployer.app3.kubernetes.requests.memory=512Mi
deployer.message-filter.kubernetes.requests.memory=1024Mi
```

## 5. Additional egress node creation fails

Ticket : [COPRS/rs-issues/#652](https://github.com/COPRS/rs-issues/issues/652)

### Issue

When creating additional egress node (i.e. the cluster is already created) you may encounter the following issue :

```console
The task includes an option with an undefined variable. The error was: {{ hostvars[groups['gateway'][0]]['ansible_default_ipv4']['gateway'] }}: 'ansible.vars.hostvars.HostVarsVars object' has no attribute 'ansible_default_ipv4'

The error appears to be in '[...]/egress.yaml': line 38, column 7, but may be elsewhere in the file depending on the exact syntax problem.
```

### Workaround

Launch the playbook `collections/kubespray/facts.yml` first :

```Bash
ansible-playbook collections/kubespray/facts.yml -i inventory/mycluster/hosts.yaml
```

## 6. Fluent Bit failed to start after a node restart

Ticket : [COPRS/rs-issues#558](https://github.com/COPRS/rs-issues/issues/558)

### Issue

Fluent Bit might fail to start after the a node is restarted with the following error :

```console
Fluent Bit v1.9.3

* Copyright (C) 2015-2022 The Fluent Bit Authors

* Fluent Bit is a CNCF sub-project under the umbrella of Fluentd

* https://fluentbit.io



[2022/09/15 08:13:34] [error] [storage] format check failed: tail.5/1-1662139558.357308467.flb

[2022/09/15 08:13:34] [error] [storage] format check failed: tail.0/1-1662139559.215290961.flb

level=error caller=out_grafana_loki.go:94 id=1 newPlugin="unable to create queue segment in /var/log/flb-storage/loki/dquer-system: unable to load queue segment in /var/log/flb-storage/loki/dquer-system: segment file /var/log/flb-storage/loki/dquer-system/0000000000680.dque is corrupted: excess deletion records (179)"

[2022/09/15 08:13:34] [error] [lib] backend failed

[2022/09/15 08:13:34] [error] [go proxy]: plugin 'grafana-loki' failed to initialize
```

### Workaround

Connect on the node and delete the folder `/var/log/flb-storage/`.

## 7. Manually bypass gateway for object storage traffic

Ticket : [COPRS/rs-issues#647](https://github.com/COPRS/rs-issues/issues/647)

### Issue

In order to stay generic and cloud agnostic, all the outgoing traffic is routed through the gateways. It might be an issue if the gateway cannot handle the traffic from the many processing nodes to the object storage. Several timeouts can be observed across the whole cluster and cause a lot of problems.

### Workaround

On every nodes ( /!\ **except the gateways** /!\ ), add the subnet of the object storage from the Cloud provider (100.64.0.0/10 in this example) to be routed directly to the VPC gateway from the Cloud provider (192.168.0.1 in this example).

1. Add the route to `100.64.0.0/10` via `192.168.0.1` in `/etc/netplan/11-ens3-private.yaml` :

   ```console
   sudo yq -i -y  '.network.ethernets.ens3.routes += [ {"to":"100.64.0.0/10", "via":"192.168.0.1"} ]' /etc/netplan/11-ens3-private.yaml
   ```

2. Apply the netplan config :

   ```console
   sudo netplan apply
   ```

3. Verify the route is active :

   ```console
   ip route
   ```

   Search for the line `100.64.0.0/10 via 192.168.0.1 dev ens3 proto static`

## 8. SCDF : First known tag is always used if not explicitly specified

Ticket : [COPRS/rs-issues#598](https://github.com/COPRS/rs-issues/issues/598)

### Issue

In SCDF, if you do not **explicitly** specify a tag version for an application, the docker tag version will be ignored and the first known will be used. For e.g. if you have the following configuration at one time, using the tag `1.12.0-rc1` in the configuration file `stream-application-list.properties` :

```shell
processor.s1l1-preparation=docker:artifactory.coprs.esa-copernicus.eu/rs-docker/rs-core-preparation-worker:1.12.0-rc1
```

SCDF will deploy the application with the tag `1.12.0-rc1`.

**Now**, if later you update the tag version to use `1.13.2-rc1` in the configuration file `stream-application-list.properties` :

```shell
processor.s1l1-preparation=docker:artifactory.coprs.esa-copernicus.eu/rs-docker/rs-core-preparation-worker:1.13.2-rc1
```

SCDF will still deploy the version `1.12.0-rc1` instead of the desired `1.13.2-rc1`.

### Workaround

You must set the application's version in the configuration file `stream-parameters.properties` :

```bash
version.preparation-worker=1.13.2-rc1
```

**And** in the configuration file `stream-application-list.properties` :

```shell
processor.s1l1-preparation=docker:artifactory.coprs.esa-copernicus.eu/rs-docker/rs-core-preparation-worker:1.13.2-rc1
```
