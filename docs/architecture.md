# Architecture

![Architecture](images/System_Architecture.jpg)

A state machine executor is the element to run job. To make it easy to scale out, we add message queue into architecture. The high throughput of system is guaranteed by Elasticsearch service.

### Message System
NATS

### Log System
Elastic Search

### Error Handle Strategy
Rebuild the log system

### Deploy
DockerFile

### Serialization and Deserialization
Json format is chosen for easy to read and develop reason. Next stage json will be replaced by Protocol Buffers to acquire better performance.