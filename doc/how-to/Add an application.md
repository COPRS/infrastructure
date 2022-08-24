# Add an application

Your application must be in a directory referenced in your inventory by ```app_path```.  
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
keycloak
mongodb
openldap
postgresql
rook-ceph
rook-ceph-cluster
spring-cloud-dataflow
```
Hence, you have to create a folder for your app. It will contain all the files and configurations needed to run the app.

> *Note: The application installer process following the alphanumerical order. You can sort your applications prefixing each one by a number with a dash.*  
*Example: 10-linkerd2*

## Helm chart

This section covers the deployment of your application as a Helm chart.

Your repository must possess 2 files:
- **config.yaml**
- **values.yaml**

Example:
```
my-new-app/
    config.yaml
    values.yaml
```

> :warning: **the name of the folder defines the name of the release deployed.** 

The ```config.yaml``` file precise the name of the chart you deploy, its version and the namespace of the application.

```yaml
chart_ref: bitnami/kafka
chart_version: 14.2.3
namespace: infra
```

Below is the list of the supported variables.

| name | mandatory \| optional | default value | comments |
|:---|:---:|---:|--------------------|
| atomic | optional | false | If set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used |
| chart_ref | mandatory | | chart_reference on chart repository<br>path to a packaged chart<br>path to an unpacked chart directory<br>absolute URL |
| chart_repo_url | optional | | Chart repository URL where to locate the requested chart. |
| chart_version | optional | latest | Chart version to install. If this is not specified, the latest version is installed. |
| create_namespace | optional | false | Create the release namespace if not present |
| include_crds | optional | false | Include CRDs in Helm templating |
| no_hooks | optional | false | Prevent hooks from running during install. More explanations about Helm lifecycle at https://helm.sh/docs/topics/charts_hooks/ |
| namespace | optional | default | Namespace where Helm metadata for the release are stored as a secret. By convention, it is also the namespace of the installed release resources. |
| timeout | optional | 5m0s | Timeout when wait option is enabled. By default, the timeout **is set to 300 seconds**. |
| wait | optional | false | Wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment are in a ready state before marking the release as successful. Helm will wait as long as what is set with ```timeout```. |

The ```values.yaml```, is the file where you write down the parameters you apply on the chart.  
You can find examples of this file, or documentation to explain you what to write on it, in the chart documentation of the application you install.

For Bitnami, the list of possible parameters is accessible in the [bitnami repository](https://github.com/bitnami/charts/tree/master/bitnami). You can read it on the README.&#xfeff;md of the folder related to the specific chart.  

### Add the repository

You need to ensure the repository of the Helm chart you add is present in ```YOUR_INVENTORY/group_vars/gateway/app_installer.yaml```.

```yaml
helm_repositories:
  - name: REPO_NAME
    repo_url: REPO_URL
  
  ...
```
The list of supported parameters is defined below.

| parameter | required ? | description |
| --- | --- | --- |
| name | yes | The name you give to the repository.<br>You will use it as prefix of the chart reference in your app config file. |
| repo_url | yes | the URL of your repository |
| repo_username | no | username of the repository.<br>For private repositories |
| repo_password | no | password for private repositories. Must be set if ```repo_username``` is set. |

## Kustomize

This section describe the deployment of an application using kustomize.

Your repository must contain a ```kustomization.yaml``` file.

> :warning: **This kustomization.yaml must defines at least the 2 following labels for all the resources: app.kubernetes.io/instance and app.kubernetes.io/managed-by**.  
Those labels are used to link the resources together and identify them in the deployment step.  

You can use the instruction ```commonLabels``` to add labels for all the resources.
```yaml
commonLabels:
  ...
  app.kubernetes.io/instance: MY_APP_NAME
  app.kubernetes.io/managed-by: Kustomize
```

## Kustomize with a Helm chart

Follow instructions given on the two sections above.  

You need to add the file ```manifest.yaml``` as resource in your kustomization file.  It is generated from the values you provided.

```yaml
resources:
  - manifest.yaml # file generated by helm template from values.yaml
  ...
```

## Deploy custom kubernetes resources with an application

You may want to add additional kubernetes resources when deploying your application.  
To do so, place your *yaml* files inside an `additional_resources` directory in your application directory in which you can sort them by type using different subdirectories.  
In the declaration of these resources, add the following labels:
```yaml
labels:
  app.kubernetes.io/instance: MY_APP_NAME
  app.kubernetes.io/managed-by: additional_resources
```
### Deploy Grafana related custom resources

To deploy *Grafana* related resources, such as *dashboards* or *datasources*, place them as explained in the the `additional_resources` directory.
There are a few custom resources deployed by default, see [here](https://github.com/grafana-operator/grafana-operator/tree/master/deploy/examples) for more examples.

### Deploy Kafka related custom resources

To deploy Kafka related custom resources, such as *topics* or *kafka clusters*, place them as explained in the the `additional_resources` directory.
There are a few custom resources deployed by default, see [here](https://github.com/strimzi/strimzi-kafka-operator/tree/0.26.0/examples) for more examples.

### Deploy Elastic cloud related custom resources

To deploy Elastic cloud related custom resources, such as *Kibana instances* or *Elastic searc instances*, place them as explained in the the `additional_resources` directory.
There are a few custom resources deployed by default, see [here](https://www.elastic.co/guide/en/cloud-on-k8s/master/k8s-api-reference.html) for reference on these CRDs.
