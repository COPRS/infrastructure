# Manage users and groups using keycloak

## Default configuration

Find the default configuration of the groups and roles here: [user_manuals/config.md](../config.md#groups)

## Add a user

Using the [keycloak documentation](https://www.keycloak.org/documentation.html) and the groups and roles present on the platform, you can create more groups and users to fit your needs. Use the console at `https://iam.{{ platform_domain_name }}/auth/admin/{{ keycloak.realm.name }}/console` using a user in the *admin* group.


> Note: For your user to be able to log into a machine using ssh, enter manually a `uidNumber` and a `homeDirectory` in the `Attributes` tab of the user management

> Note: Grafana needs the user to have an **email**, even a mock one

## Grafana's exception

Grafana natively integrates Oauth and its users and roles, but one role cannot yet be configured via Oauth: the `Server Admin`role (different than a regular admin, see https://grafana.com/docs/grafana/latest/permissions/#grafana-server-admin-role).
This role allows the user to manage plugins or organisations. This permission cannot be given to a Grafana admin through Oauth yet (see Grafana issue 29211), it is therefore not possible to have a user that would be a `Server Admin` based on its Keycloak roles.

Run these two commands on an operator's machine to give to user with id `2` (that would be the admin user created by default on the platform) the `Server Admin` permissions:
```bash
PASSWORD=$(kubectl -n monitoring get secrets grafana-admin-credentials --template={{.data.GF_SECURITY_ADMIN_PASSWORD}} | base64 -d)
kubectl exec -n monitoring deploy/grafana-deployment -c grafana -- /bin/bash -c "curl -X PUT -H 'Content-Type: application/json' -d '{\"isGrafanaAdmin\": true}' http://admin:$PASSWORD@localhost:3000/api/admin/users/2/permissions"
```
