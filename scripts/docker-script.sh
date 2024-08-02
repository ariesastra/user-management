#creating topic
docker exec -it <kafka_container_id> kafka-topics.sh --create --topic test_topic --bootstrap-server localhost:9092 --partitions 1 --replication-factor 1
kafka-console-consumer.sh --topic test_topic --bootstrap-server localhost:9092 --from-beginning