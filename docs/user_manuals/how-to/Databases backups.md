# Databases backups

## Schedule databases backups

RS is deployed with a `Stash operator` that allows scheduling periodic databases backups using specific **CRDs**.

The stash documentation can be found [here](https://stash.run/docs/v2021.11.24/concepts/crds/backupconfiguration/) and [here](https://stash.run/docs/v2021.11.24/concepts/crds/repository/).

You will need at least a `Repository` set up, see an example below.

### Example S3 bucket as Repository

```yaml
apiVersion: v1 
kind: Secret 
metadata: 
  name: backup-s3-credentials
stringData:
  AWS_ACCESS_KEY_ID: {{ s3.access_key }}
  AWS_SECRET_ACCESS_KEY: {{ s3.secret_key }}
  RESTIC_PASSWORD: resticpass
---
apiVersion: stash.appscode.com/v1alpha1
kind: Repository
metadata:
  name: s3-backup
spec:
  backend:
    s3:
      endpoint: {{ s3.endpoint }}
      bucket: BACKUP_BUCKET
      region: {{ s3.region }}
    storageSecretName: backup-s3-credentials
```

### Postgresql backup

Add the following to your postgresql `values.yaml`:

```yaml
primary:
  extraVolumeMounts:
    - mountPath: /tmp/backup
      name: tmp-backup
  extraVolumes:
    - name: tmp-backup
      emptyDir: {}
  extraEnvVarsSecret: postgresql-databases-passwords
```

Deploy the following:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: BackupConfiguration
metadata:
  name: postgresql
spec:
  driver: Restic
  repository:
    name: s3-backup
  # Daily at 4 am
  schedule: "0 4 * * ?"
  paused: false
  backupHistoryLimit: 3
  target:
    alias: postgresql
    ref:
      apiVersion: apps/v1
      kind: StatefulSet
      name: postgresql-primary
    paths:
      - /tmp/backup
    volumeMounts:
      - name: tmp-backup
        mountPath: /tmp/backup
  hooks:
    preBackup:
      exec:
        command:
          - /bin/sh
          - -c
          - >
            /bin/rm -rf /tmp/backup/* 
            && export PGPASSWORD=$KEYCLOAK_DATABASE_PASSWORD && pg_dump -U keycloak keycloak > /tmp/backup/keycloak.sql
            && export PGPASSWORD=$SCDF_DATABASE_PASSWORD && pg_dump -U scdf skipper > /tmp/backup/skipper.sql
            && export PGPASSWORD=$SCDF_DATABASE_PASSWORD && pg_dump -U scdf dataflow > /tmp/backup/dataflow.sql
            && /bin/chmod 777 /tmp/backup/*
      containerName: postgresql
  runtimeSettings:
    container:
      envFrom:
        - secretRef:
            name: postgresql-databases-passwords
      resources:
        requests:
          cpu: 100m
          memory: 128M
        limits:
          cpu: 200m
          memory: 256M
    pod:
      serviceAccountName: postgresql
      automountServiceAccountToken: true
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: node-role.kubernetes.io/infra
                operator: Exists
  retentionPolicy:
    name: 'keep-last-30'
    keepLast: 30
    prune: true
```

### LDAP backup

In the `statefulset.yaml` file:

- use the `osixia/openldap-backup` image that includes backup and restore scripts instead of the `osixia/openldap` image
- merge the following values:

  ```yaml
  spec:
   template:
     spec:
       containers:
         - name: openldap
           volumeMounts:
             - mountPath: /data/backup
               name: data-backup
       volumes:
         - name: data-backup
           emptyDir: {}
  ```

Deploy the following:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: BackupConfiguration
metadata:
  name: ldap
spec:
  driver: Restic
  repository:
    name: s3-backup
  # Daily at 04:15
  schedule: "15 4 * * ?"
  paused: false
  backupHistoryLimit: 3
  target:
    alias: openldap
    ref:
      apiVersion: apps/v1
      kind: StatefulSet
      name: openldap
    paths:
      - /data/backup
    volumeMounts:
      - name: data-backup
        mountPath: /data/backup
  hooks:
    preBackup:
      exec:
        command:
          - /bin/bash
          - -c
          - >
            /bin/rm -rf /data/backup/* 
            && /sbin/slapd-backup-config 
            && /sbin/slapd-backup-data 
            && /bin/chmod 777 /data/backup/*
      containerName: openldap
  runtimeSettings:
    container:
      resources:
        requests:
          cpu: 100m
          memory: 128M
        limits:
          cpu: 200m
          memory: 256M
    pod:
      automountServiceAccountToken: true
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: node-role.kubernetes.io/infra
                operator: Exists
  retentionPolicy:
    name: 'keep-last-30'
    keepLast: 30
    prune: true
```

### Elasticsearch backup

Because the secure configuration of a `repository` used for storing backup has to be set up on cluster startup, one is configured by default in both the elasticsearch clusters, therefore you need only to configure a `SLM policy` in the *kibana* user interface by navigating to *Management - Stack Management - Data - Snapshot and Restore* and using the documentation [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html#create-slm-policy).

## Restore databases

The restore using `Stash restic driver` is done in the following steps:

- inject a sidecar in the pod to be restored (induces a pod restart)
- download the backed-up database to the shared volume mount
- manually trigger a restore in the running pod

Use the documentation [here](https://stash.run/docs/v2021.11.24/concepts/crds/restoresession/) to configure the `RestoreSession` CRDs.

### Postgresql restore

Deploy the following:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: RestoreSession
metadata:
  name: restore-latest-postgresql-backup
  namespace: database
spec:
  driver: Restic
  task:
    name: postgresql-restore
  repository:
    name: s3-backup
  target:
    alias: postgresql
    ref:
      apiVersion: apps/v1
      kind: StatefulSet
      name: postgresql-primary
    volumeMounts:
      - name: tmp-backup
        mountPath: /tmp/backup
    rules:
      - targetHosts: []
        sourceHost: ""
        paths:
          - /tmp/backup
  runtimeSettings:
    container:
      resources:
        requests:
          cpu: 100m
          memory: 128M
        limits:
          cpu: 200m
          memory: 256M
    pod:
      serviceAccountName: postgresql
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: node-role.kubernetes.io/infra
                operator: Exists
  rules:
    - snapshots: [latest]
```

Once the `RestoreSession` has the status **Success**, run the following commands to trigger a restore (change the databases credentials):

```bash
kubectl -n database wait --for condition=ready pod -l statefulset.kubernetes.io/pod-name=postgresql-primary-0
kubectl exec -n database postgresql-primary-0 -c postgresql -- /bin/sh -c "\
  export PGPASSWORD=$KEYCLOAK_DATABASE_PASSWORD && psql -U keycloak keycloak < /tmp/backup/keycloak.sql \
  && export PGPASSWORD=$SCDF_DATABASE_PASSWORD && psql -U scdf skipper < /tmp/backup/skipper.sql \
  && export PGPASSWORD=$SCDF_DATABASE_PASSWORD && psql -U scdf dataflow < /tmp/backup/dataflow.sql"
```

### LDAP restore

Deploy the following:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: RestoreSession
metadata:
  name: latest-ldap
spec:
  driver: Restic
  task:
    name: ldap-restore
  repository:
    name: s3-backup
  target:
    alias: openldap
    ref:
      apiVersion: apps/v1
      kind: StatefulSet
      name: openldap
    volumeMounts:
      - name: data-backup
        mountPath: /data/backup
    rules:
      - targetHosts: []
        sourceHost: ""
        paths:
          - /data/backup
  runtimeSettings:
    container:
      resources:
        requests:
          cpu: 100m
          memory: 128M
        limits:
          cpu: 200m
          memory: 256M
    pod:
      affinity:
        nodeAffinity:
          requiredDuringSchedulingIgnoredDuringExecution:
            nodeSelectorTerms:
            - matchExpressions:
              - key: node-role.kubernetes.io/infra
                operator: Exists
  rules:
    - snapshots: [latest]
```

Once the `RestoreSession` has the status **Success**, run the following commands to trigger a restore:

```bash
kubectl -n iam wait --for condition=ready pod -l statefulset.kubernetes.io/pod-name=openldap-0
kubectl exec -n iam openldap-0 -c openldap -- /bin/bash -c "/sbin/slapd-restore-config \$(ls /data/backup/*-config.gz | sed -e 's/\/.*\///g') \
      && /sbin/slapd-restore-data \$(ls /data/backup/*-data.gz | sed -e 's/\/.*\///g')"
```

### Elasticsearch restore

Navigate to the *Management - Stack Management - Data - Snapshot and Restore* in the *kibana* user interface and manually trigger a restore. The interface allows fine configuration of the restore settings.

**Warning**: Indices beginning with a point such as `.geoip_database` are system indices and should not be restored.
To avoid restoring these indices, in the first step of restoring the snapshot, unselect `All data streams and indices, including system indices`, then unselect the system indices.
