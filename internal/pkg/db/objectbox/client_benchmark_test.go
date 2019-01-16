package objectbox

import (
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/test"
	"sync"
	"testing"

	contract "github.com/edgexfoundry/edgex-go/pkg/models"
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

// this breaks edgex test suite (`make test` or `go test ./...`)
// if it needs to be run manually, make it a command
//func TestBenchmarkFixedNObjectBox(t *testing.T) {
//	config := db.Configuration{
//		DatabaseName: "benchmark-test",
//	}
//	client, err := NewClient(config)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	test.BenchmarkDBFixedN(client, true)
//}

var mutex sync.Mutex

func doNoDefer() (result []contract.Event, err error) {
	mutex.Lock()

	result = []contract.Event{}

	mutex.Unlock()
	return
}
func doDefer() ([]contract.Event, error) {
	mutex.Lock()
	defer mutex.Unlock()

	return []contract.Event{}, nil
}

func BenchmarkDeferYes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doDefer()
	}
}
func BenchmarkDeferNo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doNoDefer()
	}
}
