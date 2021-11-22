# Add a new volume to a live cluster

## Create the volume

```Bash
safescale volume create --size VOLUME_SIZE --speed DISK_TYPE VOLUME_NAME
```
**VOLUME_SIZE** in gigabyte.  
**DISK_TYPE**: default HDD

## Attach the newly created volume to an existing machine of the cluster

*The volume must be attached without being formatted for Ceph to detect it. You can do this on safescale by adding the option **--do-not-format.***

```Bash
safescale volume attach --do-no-format VOLUME_NAME MACHINE_NAME
```

## Update Ceph to integrate the volume into the Ceph pool

There is two possible way to achieve this. Either by deleting the pod Rook operator or by modifying Ceph custom resource definitions (CRD).

### Delete the Pod operator

```Bash
kubectl -n NAMESPACE delete po $(kubectl -n NAMESPACE get po -l app=rook-ceph-operator -o jsonpath='{.items[0].metadata.name}')
```
The deployment generate a new Rook operator Pod which will roll out the Ceph cluster configuration pods.

### Modify the Ceph

Updating Ceph CRD provoke a new deployment of Ceph cluster configuration pods.

```Bash
kubectl edit cephclusters.ceph.rook.io -n NAMESPACE CEPH_CLUSTER_NAME
```


