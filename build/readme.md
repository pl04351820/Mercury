## Run server with container

- Start NATs Server
```
docker build . -t nats
docker run nats
```

- Start es Server
```
docker build . -t es
docker run -p 9200:9200 -p 9300:9300 -t es
```