from kafka import KafkaProducer
from kafka.errors import KafkaError

brokerUri = "kafka-cluster-kafka-bootstrap.infra.svc.cluster.local:9092"
topic = "demo.keda"

producer = KafkaProducer(bootstrap_servers=[brokerUri])

for x in range(100):
    msg = "msg-" + str(x)
    future = producer.send(topic, msg.encode('utf-8'))

    try:
        record_metadata = future.get(timeout=10)
    except KafkaError:
        log.exception()
        exit

    print(msg)
