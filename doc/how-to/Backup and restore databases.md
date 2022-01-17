# Backup and restore databases

## Postgresql database

### Configure scheduled backup on S3 bucket

The documentation used to create the daily backup can be found [here](https://stash.run/docs/v2021.11.24/concepts/crds/backupconfiguration/).

Configure the schedule and the S3 bucket credentials by setting the following variables:
  - **apps/stash/additional_resources/BackupConfiguration/postgresql-daily.yaml**:
    - *spec.schedule*
    - *spec.retentionPolicy.keepLast*

### Restore a chosen backup

*To be done*

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

To restore elasticsearch indexes, access the Kibana UI using the `elastic` username and the password got by running `kubectl get secret -n infra elasticsearch-cluster-es-elastic-user -o go-template='{{.data.elastic | base64decode}}'`.
Then navigate to `Management` - `Stack Management` - `Data` - `Snapshot and Restore` to chose a backup to restore.

## LDAP database

### Configure scheduled backup on S3 bucket
The documentation used to create the daily backup can be found [here](https://stash.run/docs/v2021.11.24/concepts/crds/backupconfiguration/).

Configure the schedule and the S3 bucket credentials by setting the following variables:
  - **apps/stash/additional_resources/BackupConfiguration/openldap-daily.yaml**:
    - *spec.schedule*
    - *spec.retentionPolicy.keepLast*
  - **apps/stash/additional_resources/Repository/s3-backup.yaml**:
    - *spec.backend.s3.endpoint*
    - *spec.backend.s3.bucket*
    - *spec.backend.s3.region*
  - **apps/stash/additional_resources/Secret/stash-s3-credentials.yaml**:
    - *stringData.AWS_ACCESS_KEY_ID*
    - *stringData.AWS_SECRET_ACCESS_KEY*

### Restore a chosen backup 

To restore a backup previously backed up on a S3 bucket by stash, create a `RestoreSession` based on this definition:

```yaml
apiVersion: stash.appscode.com/v1beta1
kind: RestoreSession
metadata:
  name: latest-ldap
  labels:
    app.kubernetes.io/instance: stash
    app.kubernetes.io/managed-by: additional_resources
spec:
  driver: Restic
  task:
    name: ldap-restore
  repository:
    name: s3-backup
  hooks:
    postRestore:
      exec:
        command:
          - /bin/sh
          - -c
          - /sbin/slapd-restore-config $(ls /data/backup/*-config.gz | sed -e 's/\/.*\///g')
      containerName: openldap
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

Configure `spec.rules.snapshots` to choose a specific snapshot, and use the documentation [here](https://stash.run/docs/v2021.11.24/concepts/crds/restoresession/).