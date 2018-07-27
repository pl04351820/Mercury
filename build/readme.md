## Run server with container

- Start NATs Server
```
docker build . -t nats
docker run -p 4222:4222 nats
```

- Start es Server
```
docker build . -t es
docker run -p 9200:9200 -p 9300:9300 es
```