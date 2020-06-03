```
docker run --name some-cassandra -d -p 9042:9042 cassandra:latest
docker run -it --rm cassandra cqlsh 192.168.1.126
```

run cql in main.go

```
go run ./main.go
```