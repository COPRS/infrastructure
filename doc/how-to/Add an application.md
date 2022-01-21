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
commonLabels:
  ...
  app.kubernetes.io/instance: MY_APP_NAME
  app.kubernetes.io/managed-by: Kustomize
```

Reach out to the following links to help you starting your journey with Kustomize.

- [Introduction to Kustomize.](https://kubectl.docs.kubernetes.io/guides/introduction/kustomize/)
- [Kustomize examples on github.](https://github.com/kubernetes-sigs/kustomize/tree/master/examples)
- [Reference docs for Kustomize’s built-in transformers and generators.](https://kubectl.docs.kubernetes.io/references/kustomize/builtins/)

### Deploy Grafana related custom resources

There are a few custom resources deployed by default, see [here](https://github.com/grafana-operator/grafana-operator/tree/master/deploy/examples) for more examples.

### Deploy Kafka related custom resources

There are a few custom resources deployed by default, see [here](https://github.com/strimzi/strimzi-kafka-operator/tree/0.26.0/examples) for more examples.

### Deploy Elastic cloud related custom resources

There are a few custom resources deployed by default, see [here](https://www.elastic.co/guide/en/cloud-on-k8s/master/k8s-api-reference.html) for reference on these CRDs.
