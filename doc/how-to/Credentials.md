# Credentials management in RS

## The app_vars.yaml inventory

All the credentials necessary to the deployment of the different applications can be set in the `app_vars.yaml` inventory file.

On the first run of the `apps.yaml` playbook, the `app_vars.yaml` file will be templated and a new `generated_app_vars.yaml` file will be written to the `artifacts` folder. 

**The values actually used by the app-installer come from the `generated_app_vars.yaml`. You will find all the credentials there.**

This workflow prevents the app-installer from changing the credentials of the apps between the deployments on the same platform, which would cause many issues.

## Generate random credentials

Like in the example values, you can choose to generate some credentials using this ansible function:
```yaml
openldap:
  admin_user_password: "{{ lookup('password', '/dev/null length=60 chars=ascii_letters') }}"
```

## Manually set credentials

Otherwise, you can freely set passwords by hand:
```yaml
graylog:
  oidc_client_secret: "m4nuAl_s3cret_example"
```

> Note: Of course, credentials that give access to external services such as S3 storage or private registries cannot be generated and have to be set by hand, they are written in UPPER_CASE in the sample inventory

## Reuse credentials

Like in the example values, you can reuse crendentials already set up in the file. This functionnality is used is the sample inventory for the S3 keys and endpoints that are often the same accross applications:
```yaml
common:
  s3:
    endpoint: S3_ENDPOINT
    region: S3_REGION
    secret_key: S3_SECRET_KEY
    access_key: S3_ACCESS_KEY

[...]

thanos:
  s3:
    bucket: THANOS_BUCKET
    endpoint: "{{ common.s3.endpoint }}"
    region: "{{ common.s3.region }}"
    access_key: "{{ common.s3.access_key }}"
    secret_key: "{{ common.s3.secret_key }}"
```


## Retrieve credentials from a HashiCorp Vault

You can retrieve credentials from a *HashiCorp Vault* instance using the *hvac* ansible plugin:

```yaml
vault:
  url: VAULT_ENDPOINT
  token: VAULT_TOKEN
  path: VAULT_PATH # add '/data/' at the end to use kv version 2
  download_app_vars: false
  upload_backup: true

[...]

common:
  s3:
    endpoint: S3_ENDPOINT
    region: S3_REGION
    secret_key: "{{ lookup('community.hashi_vault.hashi_vault', vault.path + 's3_credentials', token=vault.token, url=vault.url)['secret_key'] }}"
    access_key: "{{ lookup('community.hashi_vault.hashi_vault', vault.path + 's3_credentials', token=vault.token, url=vault.url)['access_key'] }}"
```

> Note: The *path* variable corresponds to the *secret engine* name.

## Backup and restore all values with a HashiCorp Vault

If *vault.upload_backup* is *true*, the `apps.yaml` playbook will send a backup of the `generated_app_vars.yaml` file in JSON format to the remote secret engine.

If the *vault.download_app_vars* is *true*, the `apps.yaml` playbook will download the previously backed-up `generated_app_vars.yaml` file and write it to the artifacts folder.

> Note: All the other variables will not be taken in account, this allows an operator to deploy apps to any cluster with only the 4 *vault* variables set up, and it allows any operator with vault access to read the application specific credentials directly on the vault interface
