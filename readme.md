```
docker run --name cassandra -d -p 9042:9042 cassandra:3.11.5
docker run -it --rm cassandra cqlsh 192.168.1.126
```

do this crap:
```
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
```


```
go run ./main.go
```

can repro but then gocql recovers from:
```
sudo iptables -I OUTPUT -p tcp -d 127.0.0.1 --dport 9042 -j DROP
sudo iptables -D OUTPUT 1
```

repros, but is this fair?:
```
docker update cassandra --cpus 0.1
```