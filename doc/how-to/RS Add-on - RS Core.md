# RS Add-on / RS Core

## Deploy

> Compliant with the COPRS ICD

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

## Uninstall

> **Use IHM**

### Destroy stream

![Destroy a stream](../img/destroy_stream.png)

If the stream was deployed, it is undeployed before the stream definition is deleted.

> Often, you want to stop a stream but retain the name and definition for future use. In this case, you **undeploy** a stream

### Unregister applications

> NOT MANDATORY

When you destroy a stream, you can also unregister applications of this stream.

![Unregister applications](../img/unregister_applications.png)

> Not unregister applications when only undeploy a stream

### Additionnal resources

Not necessary