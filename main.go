/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
package main

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"

	"github.com/gocql/gocql"
)

func main() {
	// connect to the cluster
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "example"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()
	defer session.Close()

	// insert a tweet
	if err := session.Query(`INSERT INTO tweet (timeline, id, text) VALUES (?, ?, ?)`,
		"me", gocql.TimeUUID(), "hello world").Exec(); err != nil {
		log.Fatal(err)
	}

	var id gocql.UUID
	var text string

	/* Search for a specific set of records whose 'timeline' column matches
	 * the value 'me'. The secondary index that we created earlier will be
	 * used for optimizing the search */
	if err := session.Query(`SELECT id, text FROM tweet WHERE timeline = ? LIMIT 1`,
		"me").Consistency(gocql.One).Scan(&id, &text); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Tweet:", id, text)

	things := 100
	//wg := sync.WaitGroup{}
	errs := int64(0)

	for i := 0; i < things; i++ {
		//	wg.Add(1)
		go func() {
			for {
				// list all tweets
				iter := session.Query(`SELECT id, text FROM tweet WHERE timeline = ?`, "me").Iter()
				/*for iter99.Scan(&id, &text) {
					fmt.Println("Tweet:", id, text)
				}*/
				if err := iter.Close(); err != nil {
					//log.Println(err)
					atomic.AddInt64(&errs, 1)
				}

				time.Sleep(100 * time.Millisecond)
			}
		}()
	}

	for {
		time.Sleep(5 * time.Second)
		log.Println(atomic.LoadInt64(&errs))
	}

	//wg.Wait()
}
