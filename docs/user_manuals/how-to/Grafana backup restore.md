# Backup and restore Grafana items

## 1. Back Grafana items

### 1.1. Get the Grafana service ip and port with

```Bash
kubectl -n monitoring get service grafana-service -o jsonpath="http://{.spec.clusterIP}:{.spec.ports[].port}"
```

This value will be set instead of `<REPLACE_ME_URL>`

### 1.2. Get the Grafana admin username

```Bash
kubectl -n monitoring get secret grafana-admin-credentials -o jsonpath='{.data.GF_SECURITY_ADMIN_USER}' | base64 -d
```

This value will be set instead of `<REPLACE_ME_USERNAME>`

### 1.3. Get the Grafana admin password

```Bash
kubectl -n monitoring get secret grafana-admin-credentials -o jsonpath='{.data.GF_SECURITY_ADMIN_PASSWORD}' | base64 -d
```

This value will be set instead of `<REPLACE_ME_PASSWORD>`

### 1.4. Create a new docker volume

`docker volume create gdg`

### 1.5. Run once the tool to initialize the folder's structure

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 ctx list`

### 1.6. Replace the configuration (Step 1)

#### 1.6.1. Edit the configuration file

`sudo vi /var/lib/docker/volumes/gdg/_data/config/importer.yml`

#### 1.6.2. Replace the content with the following template

```YAML
context_name: grafana
contexts:
    grafana:
        adminenabled: true
        organization: ""
        output_path: ""
        password: <REPLACE_ME_PASSWORD>
        storage: ""
        token: ""
        url: <REPLACE_ME_URL>
        user_name: <REPLACE_ME_USERNAME>
        watched:
            - '<REPLACE_ME_FOLDER1>'
            - '<REPLACE_ME_FOLDER2>'
global:
    debug: true
    ignore_ssl_errors: false
json_logs: false
loglevel: debug
storage_engine:
    any_label:
        access_key_id: ""
        bucket_name: ""
        cloud_type: s3
        kind: cloud
        secret_key: ""
```

#### 1.6.3. Replace the value

- `<REPLACE_ME_PASSWORD>`
- `<REPLACE_ME_URL>`
- `<REPLACE_ME_USERNAME>`

Use value from the previous steps.

### 1.7. Retrieve the list of the folder

In order to backup the dashboards, you must specify in which folder are located the dashboards.

This values will be set instead of `<REPLACE_ME_FOLDER1>` and `REPLACE_ME_FOLDER2` and so on ...

#### 1.7.1. Option 1 : gdg tool

Run the follow command to list all the folders :

```Bash
docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 folder list
INFO[0000] Listing Folders for context: 'grafana'       
┌────┬───────────┬─────────────────────────┐
│ ID │ UID       │ TITLE                   │
├────┼───────────┼─────────────────────────┤
│ 41 │ OM22BDB4z │ 0. Home (ENG)           │
│ 43 │ RFh2BvB4z │ 1. Infrastructure (ENG) │
│ 47 │ 6P22BDfVz │ 2. Service (ENG)        │
│ 57 │ GyT2fvBVz │ 3. Processing (ENG)     │
│ 66 │ 8XA2BvfVk │ 4. Performance (ENG)    │
│ 72 │ r_43BDf4z │ 5. Analysis (ENG)       │
│ 81 │ 5_S3fvB4z │ 6. OPS Manager (ENG)    │
│  1 │ Jk6kLiDVk │ autoscaling             │
│  3 │ EaezLiv4z │ database                │
│ 32 │ FK108iDVk │ iam                     │
│ 11 │ FCgmLmD4z │ infra                   │
│ 13 │ MGkmYiD4k │ logging                 │
│  9 │ RBgiYmvVz │ monitoring              │
│ 19 │ GHmiYmvVk │ networking              │
│  5 │ A3ezYmv4z │ processing              │
│  7 │ LHgiLmvVk │ rook-ceph               │
│ 17 │ ujziYiDVk │ security                │
└────┴───────────┴─────────────────────────┘
```

Copy the list of the folders that are in the column `TITLE`.

#### 1.7.2. Option 2 : Grafana GUI

Go to your Grafana instance web interface and copy the fodler's name that is next to the folder icon.

### 1.8. Replace the configuration (Step 2)

#### 1.8.1. Edit the configuration file

`sudo vi /var/lib/docker/volumes/gdg/_data/config/importer.yml`

#### 1.8.2. Replace the value

- `<REPLACE_ME_FOLDER1>`
- `<REPLACE_ME_FOLDER1>`
- and so on ...

### 1.9. Backup from Grafana to locally into the gdg volume

Depending on what you want to backup, execute the following commands.

#### 1.9.1. Backup folders into gdg volume

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 folders import`

#### 1.9.2. Backup datasources into gdg volume

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 datasources import`

#### 1.9.3. Backup dashboard into gdg volume

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 dashboards import`

### 1.10. Copy the backups

Copy the backup you just made somewhere, for e.g. in `/home/safescale/grafana_backup/`.

#### 1.10.1. Copy folders' backup

`sudo cp -r /var/lib/docker/volumes/gdg/_data/folders /home/safescale/grafana_backup/`

#### 1.10.2. Copy datasources's backup

`sudo cp -r /var/lib/docker/volumes/gdg/_data/datasources /home/safescale/grafana_backup/`

#### 1.10.3. Copy dashboard' backup

`sudo cp -r /var/lib/docker/volumes/gdg/_data/dashboards /home/safescale/grafana_backup/`

### 1.11. Delete the gdg volume

Because the password is written in clear text in the configuration, it shall be deleted when you are done.

`docker volume rm gdg`

## 2. Restore Grafana items

### 2.1. Repeat steps from 1.1 (Get the Grafana service ip and port with) to 1.8.2 (Replace the value)

### 2.2. Copy the export to the volume

For e.g. if you have your backup in `/home/safescale/grafana_backup/`.

#### 2.2.1 Folders

`sudo cp -r /home/safescale/grafana_backup/folders /var/lib/docker/volumes/gdg/_data/`

#### 2.2.2 Datasources

`sudo cp -r /home/safescale/grafana_backup/datasources /var/lib/docker/volumes/gdg/_data/`

#### 2.2.3 Dashboards

`sudo cp -r /home/safescale/grafana_backup/dashboards /var/lib/docker/volumes/gdg/_data/`

### 2.3. Export from locally in the gdg volume into Grafana

Depending on what you want to restore, execute the following commands.

#### 2.3.1. Import folders in Grafana

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 folders export`

#### 2.3.2. Import datasources in Grafana

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 datasources export`

#### 2.3.3. Import dashboard in Grafana

`docker run -it --rm --mount source=gdg,target=/app/ ghcr.io/esnet/gdg:0.4 dashboards export`

### 2.4. Delete the gdg volume

Because the password is written in clear text in the configuration, it shall be deleted when you are done.

`docker volume rm gdg`
