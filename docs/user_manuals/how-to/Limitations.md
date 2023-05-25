# Known issues and limitations

## 1. Unable to use SCDF Undeploy/Deploy function without misconfiguration

### Issue

The issue was first raised in our project [COPRS/rs-issues/issues/716](https://github.com/COPRS/rs-issues/issues/716), and later on in the SCDF project itself [spring-cloud/spring-cloud-dataflow/issues/5145](https://github.com/spring-cloud/spring-cloud-dataflow/issues/5145).

It prevents the usage of the Undeploy/Deploy functions from SCDF for any stream that contains regular expression. For e.g. we had the following regex:

`ingestion-trigger.polling.inbox1.matchRegex=^S1.*(AUX_|AMH_|AMV_|MPL_ORB).*$`

But if you undeploy and deploy again the stream, even without editing this regex, it would become (notice the **addition** of `'` at the beginning and at the end) :

`ingestion-trigger.polling.inbox1.matchRegex='^S1.*(AUX_|AMH_|AMV_|MPL_ORB).*$'`

This causes unexpected bahavior in the java application using that regex as it is not the same anymore, thus not filtering the way it should.

### Workaround

You have to destroy the stream, edit the stream's property and redeploy the stream.

1. Simply destroy the stream in the SCDF HMI
2. Edit the file `stream-parameters.properties` of the stream
3. Zip the stream
4. Install the stream using the [how-to](/docs/user_manuals/how-to/RS%20Add-on%20-%20RS%20Core.md)

## 2. Failed to generate hosts.yaml

### Issue

[The step 5 of the insfrastucture's quickstart](/README.md#5-generate-or-download-the-inventory-variables) might fail due to invalid `YAML` files. It's most probably because of bad indentation from the configuration done in the previous steps.

The issue is described in ticket [COPRS/rs-issues/issues/835](https://github.com/COPRS/rs-issues/issues/835)
### Workaround

Fix the file `inventory/host_vars/setup/main.yaml`. The `YAML` syntax could be incorrect or a non-ascii character might be there.

1. Check the `YAML` structure

   - Install the package `yq` : <https://github.com/mikefarah/yq#install>
   - Check the syntax of `inventory/host_vars/setup/main.yaml` with `yq` :
     - `yq inventory/host_vars/setup/main.yaml`

   The file's content should be displayed in the terminal. If not, it means the syntax is not correct.

2. Check for non-ascii character

   `grep --color='auto' -P -n "[^\x00-\x7F]" inventory/host_vars/setup/main.yaml`
