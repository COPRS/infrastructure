# Backup and restore databases

## Stash license

[Stash community](https://stash.run/docs/v2021.11.24/welcome/) is used to create and uploads backups on S3 buckets, to deploy stash, please get a license [here](https://license-issuer.appscode.com/?p=stash-community) and put it in **apps/stash/values.yaml**.

## Postgresql database

### Configure scheduled backup on S3 bucket

The documentation used to create the daily backup can be found [here](https://stash.run/docs/v2021.11.24/concepts/crds/backupconfiguration/).

Configure the schedule and the S3 bucket credentials by setting the following variables:
  - **apps/postgresql/additional_resources/BackupConfiguration/postgresql-daily.yaml**:
    - *spec.schedule*
    - *spec.retentionPolicy.keepLast*
  - **apps/postgresql/additional_resources/Repository/s3-backup.yaml**
    - *spec.backend.s3.endpoint*
    - *spec.backend.s3.bucket*
    - *spec.backend.s3.region*
  - **apps/postgresql/additional_resources/Secret/stash-s3-credentials.yaml**
    - *stringData.S3_ACCESS_KEY*
    - *stringData.S3_SECRET_KEY*

### Restore a chosen backup

The restore is done in two steps:
 - download chosen backup to the running pod (~3 minutes service interruption then ~3 minutes of downloading depending on network speed)
 - trigger a restore in the running pod (~1 minute, no service downtine using *psql*)

To restore a backup previously backed up on a S3 bucket by stash, create a `RestoreSession` based on this definition:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: RestoreSession
metadata:
  name: restore-latest-postgresql-backup
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
      name: postgresql-postgresql
    volumeMounts:
      - name: data-backup
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

Configure `spec.rules.snapshots` to choose a specific snapshot using [this documentation](https://stash.run/docs/v2021.11.24/concepts/crds/restoresession/).

After the backed-up data has been sent to the postresql pod, run the following commands to trigger a restore:
```bash
kubectl -n infra wait --for condition=ready pod -l statefulset.kubernetes.io/pod-name=postgresql-postgresql-0
kubectl exec -n infra postgresql-postgresql-0 -c postgresql -- /bin/bash -c "\
  export PGPASSWORD='keycloakpassword' && psql -U keycloak keycloak < /tmp/backup/keycloak.sql \
  && export PGPASSWORD='scdfpassword' && psql -U scdf skipper < /tmp/backup/skipper.sql \
  && export PGPASSWORD='scdfpassword' && psql -U scdf dataflow < /tmp/backup/dataflow.sql"
```


## Elasticsearch indexes

### Configure scheduled backup on S3 bucket
The documentation used to create the daily backup can be found [here](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html#create-slm-policy).

Configure the schedule and the S3 bucket credentials by setting the following variables:
  - **apps/030-elasticsearch-operator/additional_resources/ConfigMap/elasticsearch-backup-config.yaml**:
    - *data.snapshot-config.json.schedule*
    - *data.snapshot-config.json.retention.expire_after*
    - *data.s3-repository.json.settings.bucket*
  - **apps/030-elasticsearch-operator/additional_resources/Elasticsearch/elasticsearch-cluster.yaml**:
    - *spec.nodeSets[].config.s3.client.default.endpoint*
    - *spec.nodeSets[].config.s3.client.default.region*
  - **apps/030-elasticsearch-operator/additional_resources/Secret/eck-s3-credentials.yaml**:
    - *stringData.S3_ACCESS_KEY*
    - *stringData.S3_SECRET_KEY*

### Restore a chosen backup 

The restore is done in two steps:
 - download chosen backup to the running pods (no service interruption thanks to replication, then ~2 minutes of downloading depending on network speed)
 - trigger a restore in the leader running pod (~1 minute, no service downtine using *slapd*)

To restore elasticsearch indexes, access the Kibana UI using the `elastic` username and the password got by running `kubectl get secret -n infra elasticsearch-cluster-es-elastic-user -o go-template='{{.data.elastic | base64decode}}'`.
Then navigate to `Management` - `Stack Management` - `Data` - `Snapshot and Restore` to chose a backup to restore.

## LDAP database

### Configure scheduled backup on S3 bucket
The documentation used to create the daily backup can be found [here](https://stash.run/docs/v2021.11.24/concepts/crds/backupconfiguration/).

Configure the schedule and the S3 bucket credentials by setting the following variables:
  - **apps/openldap/additional_resources/BackupConfiguration/openldap-daily.yaml**:
    - *spec.schedule*
    - *spec.retentionPolicy.keepLast*
  - **apps/openldap/additional_resources/Repository/s3-backup.yaml**:
    - *spec.backend.s3.endpoint*
    - *spec.backend.s3.bucket*
    - *spec.backend.s3.region*
  - **apps/openldap/additional_resources/Secret/stash-s3-credentials.yaml**:
    - *stringData.AWS_ACCESS_KEY_ID*
    - *stringData.AWS_SECRET_ACCESS_KEY*

### Restore a chosen backup 

To restore a backup previously backed up on a S3 bucket by stash, create a `RestoreSession` based on this definition:

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

Configure `spec.rules.snapshots` to choose a specific snapshot using [this documentation](https://stash.run/docs/v2021.11.24/concepts/crds/restoresession/).

After the backed-up data has been sent to the openldap pods, run the following commands to trigger a restore:
```bash
kubectl -n security wait --for condition=ready pod -l statefulset.kubernetes.io/pod-name=openldap-0
kubectl exec -n security openldap-0 -c openldap -- /bin/bash -c "/sbin/slapd-restore-config \$(ls /data/backup/*-config.gz | sed -e 's/\/.*\///g') \
      && /sbin/slapd-restore-data \$(ls /data/backup/*-data.gz | sed -e 's/\/.*\///g')"
```
