# Move ceph cluster between nodes

A key feature of the 0.9.0 release, is the node group dedicated to the Ceph Cluster. If you have a cluster running version 0.8.0 or below, you may want to follow this procedure to move you filesystem to the new infrastructure repartition.

The ceph commands can be run from the `ceph-tools` container.


## 1. Backup sensitive data

This procedure is complex, and it case of failure, you may lose all your data, so please backup any sensitive data beforehand. see [Databases backups.md]('./Databases backups.md')

## 2. Create the rook_ceph node group

Following the documentation [./Cluster configuration.md]('./Cluster configuration.md') and [./Cluster scaling.md]('./Cluster scaling.md'), create the `rook_ceph` node group with at least three nodes ans the amount of storage you need.

## 3. Allow rook-ceph pods to run on the new nodes

Edit the `rook-ceph` CephCluster's placement affinity spec to allow the rook-ceph pods to run on the new nodes, as well as the old ones. Then restart the `rook-ceph-operator` pod and the new OSD should be created.

## 4. Out the old OSD from the Ceph Cluster

Following the [documentation](https://docs.ceph.com/en/latest/rados/operations/add-or-rm-osds/), mark the old OSDs *out* of the ceph cluster. Ceph should then start moving the data from the old OSD to the new ones. This step may take a few hours depending on the amount of data and the replication.

Once all the data has been transfered, on the *Ceph dashboard* 100% of the objects should be *Healthy*.

## 5. Prevent ceph pods from running on the old storage nodes

Edit again the `rook-ceph` CephCluster's placement affinity spec to allow the rook-ceph pods to run **only** on the new nodes.

## 6. Definitively delete the old OSDs

Delete the old OSDs kubernetes deployments.

## 7. Move the monitors one by one

One by one, delete the monitors following the [documentation](https://docs.ceph.com/en/latest/rados/operations/add-or-rm-mons/) as well as the corresponding `ceph-mon` deployments. New ones should be automatically created one by one on the new nodes.

**The Ceph Cluster should now be in a *HEALTHY* state. The migration is complete.**
