// Command benchmark is used to test performance of the database
package main

import (
	"github.com/edgexfoundry/go-mod-core-contracts/clients/logger"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"log"

	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/redis"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/test"
)

func main() {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
		Port: 6379,
	}

	loggingClient := logger.NewClientStdOut("benchmark", false, models.WarnLog)
	client, err := redis.NewClient(config, loggingClient)
	if err != nil {
		log.Fatal(err)
	}
	test.BenchmarkDBFixedN(client, true)
}
