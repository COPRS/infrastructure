# Add an application

## Add a package

A *package*, in *RS*, is a folder containing *apps* to be deployed when running the `apps.yaml` playbook. The package's path must be referenced in the `package_paths` variable in the `app_installer.yaml` inventory file. Use `{{ playbook_dir }}` or `{{ inventory_dir }}` as reference folder for the path.

## Create an application

Create a folder that will be used as the *release name* and *app_name* in the deployment process.
> *Note: The application installer process following the alphanumerical order. You can sort your applications prefixing each one by a number with a dash. The prefix added will not appear in the application name when deployed, its only purpose is to allow you to organize your deployments.*
*Example: 10-linkerd2*

### Deploy kubernetes resources with Kustomize

Create your `RESOURCE_TYPE-RESOURCE_NAME.yaml` file(s)

Create a `kustomize.yaml` file containing these lines:

```yaml
# kustomization.yaml

commonLabels:
  app.kubernetes.io/instance: "{{ app_name }}"

resources:
  - RESOURCE_TYPE-RESOURCE_NAME.yaml
  - ...
```

> *Note: All the files in the folder may be templated using [Jinja templating language](https://jinja.palletsprojects.com/en/3.0.x/).*

### Deploy with a Helm chart

```yaml
# kustomization.yaml

...

helmCharts:
- name: {{ app_name }}
  releaseName: {{ app_name }}
  repo: HELM_CHART_REPO_URL
  version: HELM_CHART_VERSION
  valuesFile: HELM_CHART_VALUES_FILE
  namespace: NAMESPACE
...

```

The following links provide additional documentation related to Helm integration with kustomize.

- [kustomization of a helm chart.](https://github.com/kubernetes-sigs/kustomize/blob/59c410a70af15ed330cfd5292b1a642692a7b773/examples/chart.md)
- [HelmChart structure definition.](https://github.com/kubernetes-sigs/kustomize/blob/d9435bd1b13a6764b9d271001e61837199494d1c/api/types/helmchartargs.go#L33)

### Use a private Helm repository

To pull charts from a private repository, you need to add a specific file called `helm_repository_config.yaml` containing your repository configuration at your application root path.

```yaml
# .helm_repository_config.yaml

helm_repositories:
  - name: REPOSITORY_URL
    username: "REPOSITORY_USERNAME"
    password: "REPOSITORY_PASSWORD"

```

### Supported variables

Below is the list of the variables created specifically by the `app-installer` role.

| Name | Comment |
| --- | --- |
| app_name | Name of the app folder. Used as name of your application. |
| app_dir | Root directory of your application on the bastion. It is an absolute path of this folder. |
| package_name | Name of the parent directory of your app folder. This directory is referenced as a package. |
| package_path | Absolute path of the package folder. |

All variables defined in the inventory under `inventory/mycluster/hosts_vars/setup` will be read by the `generate_inventory.yaml` playbook and used by ansible during the deployment. You may create a new file under the `apps` inventory subfolder if it does not already exists to store additional variables.

### More documentation

Reach out to the following links to help you starting your journey with Kustomize.

- [Introduction to Kustomize.](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
- [Kustomize examples on github.](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)
- [Reference docs for Kustomize’s built-in transformers and generators.](https://kubectl.docs.kubernetes.io/references/kustomize/builtins/)
