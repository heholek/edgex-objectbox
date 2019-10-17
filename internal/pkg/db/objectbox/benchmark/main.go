// Command benchmark is used to test performance of the database
package main

import (
	"log"

	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/test"
)

func main() {
	config := db.Configuration{
		DatabaseName: "objectbox",
	}
	client, err := objectbox.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	test.BenchmarkDBFixedN(client, true)
}
