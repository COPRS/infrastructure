# Create and mount a Ceph shared filesystem

## Create a PVC

To create a shared filesystem on the platform that multiple pods will read or write into, configure and apply the following persistent volume:
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: shared-processing-fs
spec:
  accessModes:
    - ReadWriteMany
  volumeMode: Filesystem
  resources:
    requests:
      storage: 80Gi
  storageClassName: ceph-fs
```

## Create a pod and mount the PVC

Then create a pod that mounts the shared filesystem:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: shared-fs-pod
spec:
  containers:
    - name: shared-fs-pod
      image: d3fk/s3cmd
      volumeMounts:
      - mountPath: /mnt/shared-fs
        name: shared-fs
      command:
        - /bin/sh
        - -c
        - "apk add tar && sleep 1d"
  volumes:
    - name: shared-fs
      persistentVolumeClaim:
        claimName: shared-processing-fs
```

# Write files from s3 bucket

Get a shell in the running container:

```shell
kubectl exec -it shared-fs-pod -c shared-fs-pod -- /bin/sh
```

Read the documentation (here)[https://s3tools.org/s3cmd-howto] on how to use `s3cmd`, some details may depend on the S3 location or bucket structure.
For example, run these commands to copy a file named **file.zip** from the root of an S3 bucket named **test-s3-bucket** to the shared filesystem:

```shell
# Configure your S3 credentials (zone, secret key, endpoint...)
s3cmd --configure

# Sync the content of the bucket to the mounted file system
s3cmd sync s3://test-s3-bucket/ /mnt/shared-fs
```

# Send files from a local computer

Send the local file **/run/media/USB/file.zip** to the pod with the following command:

```shell
kubectl cp --retries=100 /run/media/USB/file.zip shared-fs-pod:/mnt/shared-fs/file.zip -c shared-fs-pod
```
