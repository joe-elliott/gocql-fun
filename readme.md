```
docker run --name some-cassandra -d -p 9042:9042 cassandra:latest
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
````