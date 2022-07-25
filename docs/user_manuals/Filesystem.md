## Add a new volume to a live cluster

### Create the volume

```Bash
safescale volume create --size VOLUME_SIZE --speed DISK_TYPE VOLUME_NAME
```
**VOLUME_SIZE** in gigabyte.  
**DISK_TYPE**: default HDD

### Attach the newly created volume to an existing machine of the cluster

*The volume must be attached without being formatted nor mounted for Ceph to detect it. You can do this on safescale by adding the options **--do-not-format** and **--do-not-mount***. 

```Bash
safescale volume attach --do-not-format --do-not-mount VOLUME_NAME MACHINE_NAME
```

### Update Ceph to integrate the volume into the Ceph pool

There is two possible way to achieve this. Either by deleting the pod Rook operator or by modifying Ceph custom resource definitions (CRD).

#### Delete the Pod operator

```Bash
kubectl -n NAMESPACE delete po $(kubectl -n NAMESPACE get po -l app=rook-ceph-operator -o jsonpath='{.items[0].metadata.name}')
```
The deployment generate a new Rook operator Pod which will roll out the Ceph cluster configuration pods.

#### Modify the CRD

Updating Ceph CRD provoke a new deployment of Ceph cluster configuration pods.

```Bash
kubectl edit cephclusters.ceph.rook.io -n NAMESPACE CEPH_CLUSTER_NAME
```

## Known errors

### PVC stuck on pending

On cluster startup PVC may need several minutes before bounding with a PV. It is the time for the Ceph cluster to be fully operational. Meanwhile the other app relying on persistent storage are stuck in pending mode with theirs PVC.

If you are still worried, you can check that the OSD pods are running and ready.

```Bash
kubectl get po -n NAMESPACE -l app=rook-ceph-osd
```

You can also check that the Ceph cluster is in a healthy status.

```Bash
kubectl -n NAMESPACE get cephclusters.ceph.rook.io
```

If the cluster is not in ```HEALTH_OK```, you can get some explaining informations with the following command.

```Bash
kubectl -n NAMESPACE get cephclusters.ceph.rook.io CLUSTER_NAME -ojsonpath='{.status.ceph.details}' | jq
```

### An operation with the give volume ID already exists

The full error message is:  

```failed to provision volume with StorageClass "ceph-block": rpc error: code = Aborted desc = an operation with the given Volume ID pvc-xxx already exists```

Delete the pod provisioner in charge of the block storage

```Bash
kubectl get po -n NAMESPACE -l app=csi-rbdplugin-provisioner
```
