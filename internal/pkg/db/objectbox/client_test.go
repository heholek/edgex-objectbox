package objectbox

import (
	"testing"

	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/test"
	"github.com/edgexfoundry/edgex-go/pkg/models"
)

func TestObjectBox(t *testing.T) {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
	}
	client := NewClient(config)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}

	event := models.Event{
		Device: "my device",
	}
	objectId, err := client.AddEvent(&event)
	if err != nil {
		t.Fatalf("Could not add event: %v", err)
	}
	t.Logf("Added object ID %v", objectId)

	event.Device = "2nd device"
	objectId, err = client.AddEvent(&event)
	if err != nil {
		t.Fatalf("Could not add 2nd event: %v", err)
	}
	t.Logf("Added 2nd object ID %v", objectId)

	eventRead, err := client.EventById(string(objectId))
	if err != nil {
		t.Fatalf("Could not get 2nd event by ID: %v", err)
	}
	if event.ID != eventRead.ID || event.Device != eventRead.Device {
		t.Fatalf("Event data error: %v vs. %v", event, eventRead)
	}

}

func BenchmarkObjectBox(b *testing.B) {
	config := db.Configuration{
		DatabaseName: "benchmark-test",
	}
	client := NewClient(config)
	err := client.Connect()
	if err != nil {
		fmt.Println(err)
		return
	}
	test.BenchmarkDB(b, client)
}
