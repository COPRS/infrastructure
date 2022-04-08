## Grafana components backup

### Dashboards

The grafana operator allows for grafana dashboards to be saved in configuration files as `GrafanaDahsboard` resources.

To create such a configuration file, follow the [grafana operator documentation](https://github.com/grafana-operator/grafana-operator/tree/master/deploy/examples/dashboards), which provides many examples of `GrafanaDashboard` configuration files.

Once the created the configuration file has been created:
- put it in the same directory as the app it concerns. For instance, the `myapp` dashboard should be at `apps/myapp/mygrafanadashboard.yaml`.
- add the name of the configuration file in the `kustomization.yaml` file under the "resources" as so:

```yaml
# apps/myapp/kustomization.yaml
namespace: mynamespace

commonLabels:
  app.kubernetes.io/instance: "{{ app_name }}"

helmCharts:
- name: myapp
  releaseName: "{{ app_name }}"
  repo: myrepo
  version: 1
  valuesFile: values.yaml
  namespace: mynamespace

resources:
  - mygrafanadashboard.yaml
```

On the next deployment, the dashboard will be deployed during the deployment of `myapp`.

### Datasources

Just as dashboards, the grafana operator allows for data sources to be saved in configuration files. 

The grafana operator documentation provides [some examples of datasources](https://github.com/grafana-operator/grafana-operator/tree/master/deploy/examples/datasources). If you need to create data sources for elasticsearch, please follow this example:

```yaml
apiVersion: integreatly.org/v1alpha1
kind: GrafanaDataSource
metadata:
  name: myelasticsearch-datasource
  labels:
    app.kubernetes.io/instance: myelasticsearch
spec:
  name: myelasticsearch-datasource
  datasources:
    - name: myelasticsearch-myindex
      type: elasticsearch
      access: Server
      database: "myindex"
      url: http://myelasticsearch-server:9200/
      isDefault: false
      version: 1
      editable: true
      jsonData:
        tlsSkipVerify: true
        timeField: "myTimeField"
        esVersion: 7.15.2
        logMessageField: "myMessageField"
```

Once, the datasource file has been created:
- put it in the same directory as the app it concerns. For instance, the `myapp` datasource should be at `apps/myapp/mygrafanadatasource.yaml`.
- add the name of the configuration file in the `kustomization.yaml` file under the "resources", and a json patch, as so:

```yaml
# apps/myapp/kustomization.yaml
namespace: mynamespace

commonLabels:
  app.kubernetes.io/instance: "{{ app_name }}"

helmCharts:
- name: myapp
  releaseName: "{{ app_name }}"
  repo: myrepo
  version: 1
  valuesFile: values.yaml
  namespace: mynamespace

resources:
  - mygrafanadatasource.yaml

# Put grafana datasource in the right namespace
patchesJson6902:
  - target:
      group: integreatly.org
      version: v1alpha1
      kind: GrafanaDataSource
      name: myelasticsearch-datasource
    patch: |-
      - op: replace
        path: /metadata/namespace
        value: monitoring
```
On the next deployment, the datasource will be deployed during the deployment of `myapp`.

## Keycloak

To download a realm, in the administration console :
- go to `Export`
- choose the data you wish to export
- download the realm and place it under `apps/keycloak/custom-realm.json`

On the next deployment, the realm will be imported in keycloak.

**Remark**: The current realm management policy is to overwrite configuration at each new deployment of keycloak. If you wish to change the policy, in `apps/keycloak/values.yaml`, line 80 change `OVERWRITE_EXISTING` to `IGNORE_EXISTING`.

## Graylog

Graylog components such as dashboards or streams can be saved in content packs.

To create a content pack:
- in the Graylog HMI, on the top menu, select `System` -> `Content Packs`
- Click on the top right button `Create a content pack`
- Select any resource you wish to save
- Fill out the other mandatory fields
- Download the content pack

**Warning**: Some resources in the content pack may appear twice in the content pack. It will cause no error, however, on the restore of the content pack, the resource will appear twice in the HMI.

To deploy this content pack, follow the instrictions in [the Graylog content pack  manual](Graylog%20content%20packs.md).