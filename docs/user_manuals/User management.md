# Manage users and groups using keycloak

## Default configuration

Find the default configuration of the groups and roles here: [user_manuals/config.md](../config.md#groups)

## Add a user

Using the [keycloak documentation](https://www.keycloak.org/documentation.html) and the groups and roles present on the platform, you can create more groups and users to fit your needs. Use the console at *https://iam.{{ platform_domain_name }}/auth/admin/{{ keycloak.realm.name }}/console* using a user in the *admin* group.


> Note: For your user to be able to log into a machine using ssh, enter manually a `uidNumber` and a `homeDirectory` in the `Attributes` tab of the user management

> Note: Grafana needs the user to have an **email**, even a mock one

