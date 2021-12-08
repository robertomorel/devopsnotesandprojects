# Go - Kafka - Example

## Apache Kafka

### Commands
First of all, run the docker-compose `docker-compose up`, than run `docker exec -it kafka_kafka_1 bash`

#### Topics

- Available options for topics
  - `kafka-topics`
- Creating topic called 'test' on the server 'localhost:9092' with 3 partitions
  - `kafka-topics --create --topic=test --bootstrap-server=localhost:9092 --partitions=3`
- Listing all the topic (many are created automatically)
  - `kafka-topics --list --bootstrap-server=localhost:9092`  
- Describing a topic  
  - `kafka-topics --bootstrap-server=localhost:9092 --topic=test --describe` 

#### Consumer

- Available options for consumers
  - `kafka-console-consumer`
- Setting a consumer which is hearing the topic
  - `kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test` 
- Setting a consumer which is hearing the topic and can receive all the messages from creation of the topic 
  - `kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test --from-beginning`  
- Setting a consumer in a group called 'x', which is hearing the topic 
  - `kafka-console-consumer --bootstrap-server=localhost:9092 --topic=test group=x` 
- Available options for consumer groups  
  - `kafka-consumer-groups`  
- Describing some consumer group
  - `kafka-consumer-groups --bootstrap-server=localhost:9092 group=x --describe`        

#### Producer

- Available options for producers
  - `kafka-console-producer`
- Setting a producer which can send messages to the topic
  - `kafka-console-producer --bootstrap-server=localhost:9092 --topic=test`
  - After that, we can send any text message we want. e.g.


## Initiating Go-Lang in the project
Inside the docker container, run: `go mod init github.com/robertomorel/gokafka`

## How to run
1. `docker-compose up -d`

2. To check if everything is running: `docker-compose ps`

3. Entering in the container: `docker exec -it gokafka bash` 

4. Setting a consumer up in another terminal: `docker exec -it gokafka_kafka_1 bash`

 > - "Ctrl + l" you clean the terminal 