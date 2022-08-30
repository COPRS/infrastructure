# Using Grafana with the Grafana operator

Grafana and grafana related resources are deployed and maintainted up to date by the [grafana operator](https://github.com/grafana-operator/grafana-operator/tree/master/documentation)

The operator, every 30 minutes, loads the dashboards content from the sources defined in every *GrafanaDashboard* CRDs deployed in the cluster and syncs them with the *Grafana* instance. Therefore you should not edit dashboards that have been set up on Grafana by the operator, but work on copies of the dashboard and then export these copies to update the source referenced in the corresponding CRD.


# Deploy a dashboard using the operator

To deploy a *Grafana Dashboard* using the *Grafana Operator*, refer to the operator's documentation linked above. There is an example of a dashboard json source in the `scaler/dashboards` folder that in referenced in the `apps/autoscaling/grafana-dashboards.yaml`.
It is a good pratice to always specify the *spec.customFolderName* in the *GrafanaDashboard*.

For now, the *GrafanaDatasources* must be in the *monitoring* namespace, but the *GrafanaDashboard* resources can be in any namespace, and if you do not set a *customFolderName*, the operator will place the dashboard in a folder named according to the namespace.

# Deploy a library panel

Library panels are not managed by Grafana a independant resources but are part of a dashboard. In order to save and deploy library panels, you must first create a dashboard with a library panel, then export it and the library panel will be included in the exported JSON dashboard.

# Edit a deployed dashboard

## Duplicate the dashboard

 - Create a dedicated dashboard folder (folder name cannot be the same as a namespace in the cluster)
 - Open the dashboard
 - In settings, click "Save As..." ans save it in the previously created folder

## Edit and test the new dashboard

You can then do the modifications you need, test new panels and more without interference from the Grafana Operator on the newly duplicated dashboard.

## Export the new dashboard

 - Click on the *Share* logo ![](../../media/share_grafana.png)
 - Go to the export tab, toggle ON *Export for sharing externally*
 - Click on *View JSON* and overwrite the original dashboard source with the JSON output
 - Check that the "name" values in the "__inputs" field at the top of the JSON correspond to the information written in the corresponding *Grafana Dashboard* CRD at *spec.datasources* (DS_THANOS are usually DS_PROMETHEUS, edit one or the other to make it match)

## Upload the new version to the remote source

You may then commit and upload (or any similar operation to update the source dashboard). In less than 30 minutes, the *Grafana Operator* should update the dashboard on the *Grafana* instance.
