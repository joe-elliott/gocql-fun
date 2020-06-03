```
docker run --name some-cassandra -d -p 9042:9042 cassandra:latest
docker run -it --rm cassandra cqlsh 192.168.1.126
```

run cql in main.go

```
go run ./main.go
```

can repro but then gocql recovers from:
```
sudo iptables -I OUTPUT -p tcp -d 127.0.0.1 --dport 9042 -j DROP
sudo iptables -D OUTPUT 1
````