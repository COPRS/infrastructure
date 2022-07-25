# Credentials management in RS

## The generate_inventory.yaml playbook and inventory files

All the credentials necessary to the deployment of the different applications can be set in the different inventory files located under `{{ inventory_dir }}/host_vars/setup`.

Setup your variables by:
 - using separate files for variables corresponding to different applications, like in the sample inventory and its `{{ inventory_dir }}/host_vars/setup/apps` subfolder
 - creating new files (for example `{{ inventory_dir }}/host_vars/setup/production_env1.yaml`) with your variables
 - editing the variables given as sample in the sample inventory

On the run of the `generate_inventory.yaml` playbook, the files under `{{ inventory_dir }}/host_vars/setup` will be templated and a new `generated_inventory_vars.yaml` file will be written to the `{{ inventory_dir }}/group_vars/all` folder. 

**The values actually used by the app-installer come from the `generated_inventory_vars.yaml`. You will find all the credentials there.**

This workflow prevents the app-installer from changing the credentials of the apps between the deployments on the same platform, which would cause many issues.

## Generate random credentials

Like in the example values, you can choose to generate some credentials using this ansible function:
```yaml
# {{ inventory_dir }}/host_vars/setup/apps/openldap.yaml
openldap:
  admin_user_password: "{{ lookup('password', '/dev/null length=60 chars=ascii_letters') }}"
```

## Manually set credentials

Otherwise, you can freely set passwords by hand:
```yaml
# {{ inventory_dir }}/host_vars/setup/apps/graylog.yaml
graylog:
  oidc_client_secret: "m4nuAl_s3cret_example"
```

> Note: Of course, credentials that give access to external services such as S3 storage or private registries cannot be generated and have to be set by hand, they are written in UPPER_CASE in the sample inventory

## Reuse credentials

Like in the example values, you can reuse crendentials already set up in the inventory files. This functionnality is used in the sample inventory for the S3 keys and endpoints that are often the same accross applications:
```yaml
# {{ inventory_dir }}/host_vars/setup/main.yaml
s3:
  endpoint: S3_ENDPOINT
  region: S3_REGION
  secret_key: S3_SECRET_KEY
  access_key: S3_ACCESS_KEY
```

```yaml
# {{ inventory_dir }}/host_vars/setup/apps/thanos.yaml
thanos:
  s3:
    bucket: THANOS_BUCKET
    endpoint: "{{ common.s3.endpoint }}"
    region: "{{ common.s3.region }}"
    access_key: "{{ common.s3.access_key }}"
    secret_key: "{{ common.s3.secret_key }}"
```


## Retrieve indivindual credentials from a HashiCorp Vault

You can retrieve credentials from a *HashiCorp Vault* instance using the *hvac* ansible plugin:

```yaml
# {{ inventory_dir }}/host_vars/setup/main.yaml
vault:
  url: VAULT_ENDPOINT
  token: VAULT_TOKEN
  path: VAULT_PATH # add '/data/' after the secret engine name to use kv version 2
  download_inventory_vars: false
  upload_backup: true
  upload_existing: false

[...]


s3:
  endpoint: S3_ENDPOINT
  region: S3_REGION
  secret_key: "{{ lookup('community.hashi_vault.hashi_vault', vault.path + 'SECRET_NAME', token=vault.token, url=vault.url)['KEY_IN_SECRET'] }}"
  access_key: "{{ lookup('community.hashi_vault.hashi_vault', vault.path + 'SECRET_NAME', token=vault.token, url=vault.url)['KEY_IN_SECRET'] }}"
```


## Backup and restore all inventory variables with a HashiCorp Vault

### Options:

- **vault.upload_backup**: send a backup of the `generated_inventory_vars.yaml` file in JSON format to the remote secret engine.

- **vault.upload_existing**: upload the existing `generated_inventory_vars.yaml` file without generating it before *(useful if it has been edited by hand)*

- **vault.download_inventory_vars**: download the previously backed-up `generated_inventory_vars.yaml` file and write it to the `{{ inventory_dir }}/group_vars/all` folder.

  > Note: All the other variables will not be taken in account, this allows an operator to deploy apps to any cluster with only the few *vault* variables set up, and it allows any operator with vault access to read the application specific credentials directly on the vault web interface.
