package objectbox

import (
	"testing"

	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/test"
)

func BenchmarkObjectBox(b *testing.B) {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
	}
	client, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	test.BenchmarkDB(b, client)
}

func TestBenchmarkFixedNObjectBox(t *testing.T) {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
	}
	client, err := NewClient(config)
	if err != nil {
		fmt.Println(err)
		return
	}
	test.BenchmarkDBFixedN(client, true)
}
