// Command benchmark is used to test performance of the database
package main

import (
	"log"

	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/redis"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/test"
)

func main() {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
		Port: 6379,
	}
	client, err := redis.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	test.BenchmarkDBFixedN(client, true)
}
