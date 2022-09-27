# Set up grafana SMTP alerting

All this documentation is based on the official [Grafana Documentation](https://grafana.com/docs/grafana), please refer to it for more specific use cases.
By default, an image renderer is deployed alongside Grafana, and will add images. 

## Set up SMTP credentials in the ansible variables

In order for Grafana to send emails, set up SMTP credentials in the ansible variables, either directly in the `{{ inventory_dir }}/group_vars/all/generated_inventory_vars.yaml` file or before the inventory generation in the `{{ inventory_dir }}/hosts_vars/setup/apps/grafana.yaml` like:

```yaml
grafana:
  smtp:
    enabled: true
    host: SMTP_HOST
    user: SMTP_USER
    password: SMTP_PASSWORD
    from_address: FROM_ADDRESS
    from_name: FROM_NAME
```

These settings will be written in the grafana configuration on deployment.


## Set up actual alerting

In order to have a functionnal alerting mecanism, you will have to configure the following Grafana components:

 - ### Alert rules

Set evaluation criteria that determines whether an alert instance will fire. An alert rule consists of one or more queries and expressions, a condition, the frequency of evaluation, and optionally, the duration over which the condition is met.

Grafana managed alerts support multi-dimensional alerting, which means that each alert rule can create multiple alert instances. This is exceptionally powerful if you are observing multiple series in a single expression.

Once an alert rule has been created, they go through various states and transitions. The state and health of alert rules help you understand several key status indicators about your alerts.

If you want a capture of a dashboard embed in you notification, do not forget to set up the **Dashboard UID** and the **Panel ID** parameters here.

 - ### Labels

Match an alert rule and its instances to notification policies and silences. They can also be used to group your alerts by severity.

- ### Notification policies

Set where, when, and how the alerts get routed. Each notification policy specifies a set of label matchers to indicate which alerts they are responsible for. A notification policy has a contact point assigned to it that consists of one or more notifiers.

 - ### Contact points

Define how your contacts are notified when an alert fires, an email list for example.
