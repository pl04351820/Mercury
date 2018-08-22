# Architecture

![Architecture](images/System_Architecture.jpg)

A state machine executor is the element to run job. To make it easy to scale out, we add message queue into architecture. The high throughput of system is guaranteed by Elasticsearch service.

Term | Explanation
 :---: |  :---:
NATS  | Cloud-Native message system. Commit-Log Paradigm and supports for Request Reply
Log Service | <center> ElasticSearch Cluster Service  </center>
Projector | Constructor and Reconstruct FSM
API Proxy |  Invoke Serverless Function from Third Party Service 
