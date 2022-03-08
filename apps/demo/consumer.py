from kafka import KafkaConsumer
import time

delay = 1
consumerGroup = "demo"
topic = "demo.keda"
brokerUri = "kafka-cluster-kafka-bootstrap.infra.svc.cluster.local:9092"

consumer = KafkaConsumer(topic,
                         group_id=consumerGroup,
                         bootstrap_servers=[brokerUri])
for message in consumer:
    print ("%s:%d:%d: key=%s value=%s" % (message.topic, message.partition,
                                          message.offset, message.key,
                                          message.value))
    time.sleep(delay)
