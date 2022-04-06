# Logs monitoring and retention

## Store kafka topics content in Loki

Fluentd can be used to read the content of some *Kafka* topics, and send it to *Loki* for monitoring and retention.

To achieve this, set a regex that matches the topic(s) name(s) to be read and stored in the **topic_to_loki_regex** deployment var in the **fluentd.yaml** inventory file.

The content of the topic(s) will be available in *Loki* using the label(s) `{kafka_topic="TOPIC_NAME"}`.


## Monitor system logs

System logs are accessible in *Loki* using the labels `SYSLOG_IDENTIFIER` and `HOSTNAME`.
The log content is in `MESSAGE` field.

## Monitor applications logs

Applications logs are accessible in *Loki* using the labels `kubernetes_container_image`, `kubernetes_container_name`, `kubernetes_host`, `kubernetes_namespace_name` and `kubernetes_pod_name`.
The applicative log is in the `log` field.
