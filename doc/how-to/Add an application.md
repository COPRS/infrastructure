# Add an application

## Register your application

Your application must be in a directory referenced in your inventory by ```app_paths```.  
You can create several folder depending on your needs.

**For example**:

```yaml
# platform/inventory/sample/group_vars/gateway/app_installer.yaml

...

app_paths:
  - "{{ playbook_dir }}/../../apps"
  - "/home/foo/myproject/myapps"        # new directory added
```

Each folder contains a list of applications you will deploy on the platform.

For example: 
```Bash
$ ls -A1 ./apps
elasticsearch
graylog
kafka
```
Hence, you have to create a folder for your app. It will contain all the files and configurations needed to run the app.

> *Note: The application installer process following the alphanumerical order. You can sort your applications prefixing each one by a number with a dash.*  
> *Note: The prefix added will not appear in the application name when deployed. Its only purpose is to allow you to organize your deployments.* 
*Example: 10-linkerd2*

## Configure your application

The app_installer role exploit [Kustomize](https://kustomize.io/) to deploy your applications.

Each application folder must contains a ```kustomization.yaml``` file.
Each resource defined must include the label ```app.kubernetes.io/instance: YOUR_APP_NAME```.

You can use the instruction ```commonLabels``` to add labels simultaneously on all the resources.

```yaml
# kustomization.yaml

commonLabels:
  ...
  app.kubernetes.io/instance: {{ app_name }}
```

*The files listed in your application folder are parsed by ansible. Therefore, you can exploit ansible templating and your inventory to configure it.*

Reach out to the following links to help you starting your journey with Kustomize.

- [Introduction to Kustomize.](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
- [Kustomize examples on github.](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)
- [Reference docs for Kustomize’s built-in transformers and generators.](https://kubectl.docs.kubernetes.io/references/kustomize/builtins/)

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

...

```

The following links provide additional documentation related to Helm integration with kustomize.

- [kustomization of a helm chart.](https://github.com/kubernetes-sigs/kustomize/blob/59c410a70af15ed330cfd5292b1a642692a7b773/examples/chart.md)
- [HelmChart structure definition.](https://github.com/kubernetes-sigs/kustomize/blob/d9435bd1b13a6764b9d271001e61837199494d1c/api/types/helmchartargs.go#L33)

### Supported variables

Below is the list of the variables created specificaly by the app installer role.

| Name | Comment |
| --- | --- |
| app_name | Name of the app folder. Used as name of your application. |
| app_dir | Root directory of your application on the bastion. It is an absolute path of this folder. |
| package_name | Name of the parent directory of your app folder. This directory is referenced as a package. |
| package_path | Absolute path of the package folder. |

All variables defined in the inventory can also be exploited in the applications.

### Deploy Grafana related custom resources

There are a few custom resources deployed by default, see [here](https://github.com/grafana-operator/grafana-operator/tree/master/deploy/examples) for more examples.

### Deploy Kafka related custom resources

There are a few custom resources deployed by default, see [here](https://github.com/strimzi/strimzi-kafka-operator/tree/0.26.0/examples) for more examples.

### Deploy Elastic cloud related custom resources

There are a few custom resources deployed by default, see [here](https://www.elastic.co/guide/en/cloud-on-k8s/master/k8s-api-reference.html) for reference on these CRDs.
