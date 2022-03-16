# RS add-ons managements

## Deploy an RS add-on compliant with the COPRS ICD

 - Place the zip file in any directory on the bastion
 - Run the `deploy-rs-addon.yaml` playbook with the following variables:
   - **stream_name**: name given by *Spring Cloud Dataflow* to the created stream
   - **rs_addon_url**: direct download url of the zip file

For example:
```shellsession
ansible-playbook deploy-rs-addon.yaml \
    -i inventory/mycluster/hosts.ini \
    -e rs_addon_url=https://artifactory.coprs.esa-copernicus.eu/artifactory/demo-zip/demo-rs-addon.zip \
    -e stream_name=example-stream-name
```
