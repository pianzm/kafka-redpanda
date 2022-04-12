# kafka-redpanda
Compare Benchmark Kafka and Redpanda

Testing Kafka performance compare dengan Redpanda:

```shell
➜  kafka-redpanda git:(main) go run main.go run-consumer -p kafka
Produce 1000000 records in 2.905461s
Consume 1000000 records in 587.410708ms
➜  kafka-redpanda git:(main) go run main.go run-consumer -p panda
Produce 1000000 records in 1.2659305s
Consume 1000000 records in 519.069625ms
```


Masing-masing running di docker dengan 1 broker.

Summary:
Advantage Kafka
High throughput
Low latency
Resistance to node/machine failure within a cluster
Data persist on disk
Scalability
Distributed
Disadvantages of Kafka
Only listen specific topic defined
Require Zookeeper

Advantage of Redpanda
Implement all Kafka API
Single binary to deploy
Any library built for Kafka are inter-operable for Redpanda
Up to 6x faster Kafka transactions

For references:
https://redpanda.com/platform/
