# Optimization for well-known zones and/or internal-only domains

Edit the file `inventory/mycluster/host_vars/setup/kubespray.yaml` to add `coredns_external_zones` and `nodelocaldns_external_zones` with the cloud provider internal network.

For e.g. with Orange flexible engine, add the following snippet at the end of the file to use the Cloud provider's DNS to resolve the internal networks for the object storage I/O:
```yaml
coredns_external_zones:
- zones:
  - prod-cloud-ocb.orange-business.com
  nameservers:
  - 100.125.0.41
  - 100.126.0.41
nodelocaldns_external_zones:
- zones:
  - prod-cloud-ocb.orange-business.com
  nameservers:
  - 100.125.0.41
  - 100.126.0.41
```
