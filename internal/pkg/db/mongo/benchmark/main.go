// Command benchmark is used to test performance of the database
package main

import (
	"log"

	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/mongo"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/test"
)

func main() {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
		Host:         "0.0.0.0",
		Port:         27017,
		Timeout:      1000,
	}
	client, err := mongo.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	test.BenchmarkDBFixedN(client, true)
}
