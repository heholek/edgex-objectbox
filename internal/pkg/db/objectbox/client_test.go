package objectbox

import (
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/test/assert"
)

func createClient() *Client {
	config := db.Configuration{
		DatabaseName: "db-unit-test",
	}
	client := NewClient(config)
	err := client.Connect()
	if err != nil {
		panic("Could not connect DB client: " + err.Error())
	}
	return client
}

func TestObjectBoxEvents(t *testing.T) {
	client := createClient()
	defer client.Disconnect()

	event := models.Event{Device: "my device"}
	objectId, err := client.AddEvent(&event)
	assert.NoErr(t, err)
	t.Logf("Added object ID %v", objectId)
	assert.Eq(t, objectId, event.ID)

	event = models.Event{Device: "2nd device"}
	objectId, err = client.AddEvent(&event)
	assert.NoErr(t, err)

	t.Logf("Added 2nd object ID %v", objectId)

	eventRead, err := client.EventById(string(objectId))
	assert.NoErr(t, err)

	if event.ID != eventRead.ID || event.Device != eventRead.Device {
		t.Fatalf("Event data error: %v vs. %v", event, eventRead)
	}
}

func TestObjectBoxReadings(t *testing.T) {
	client := createClient()
	defer client.Disconnect()

	assert.NoErr(t, client.ScrubAllEvents())
	countPre, err := client.eventBox.Count()
	assert.NoErr(t, err)
	assert.Eq(t, uint64(0), countPre)
	countPre, err = client.readingBox.Count()
	assert.Eq(t, uint64(0), countPre)

	objectId, err := client.AddReading(models.Reading{Name: "reading1", Device: "device42"})
	assert.NoErr(t, err)
	t.Logf("Added reading ID %v", objectId)
	assert.NotEq(t, "", objectId)

	objectId2, err := client.AddReading(models.Reading{Name: "reading2", Device: "device42"})
	assert.NoErr(t, err)

	t.Logf("Added 2nd reading ID %v", objectId2)

	objectId3, err := client.AddReading(models.Reading{Name: "reading3", Device: "device43"})
	assert.NoErr(t, err)

	t.Logf("Added 3rd reading ID %v", objectId3)
	count, err := client.ReadingCount()
	assert.NoErr(t, err)

	assert.Eq(t, 3, count)

	readingRead, err := client.ReadingById(string(objectId2))
	if err != nil {
		t.Fatalf("Could not get 2nd reading by ID: %v", err)
	}
	assert.Eq(t, objectId2, readingRead.Id)
	assert.Eq(t, "device42", readingRead.Device)

	all, err := client.Readings()
	assert.NoErr(t, err)
	assert.Eq(t, 3, len(all))
	assert.Eq(t, "device42", all[0].Device)
	assert.Eq(t, "device42", all[1].Device)
	assert.Eq(t, "device43", all[2].Device)

	readings, err := client.ReadingsByDevice("device42", 10)
	assert.Eq(t, 2, len(readings))
	assert.Eq(t, objectId, readings[0].Id)
	assert.Eq(t, "device42", readings[0].Device)
	assert.Eq(t, objectId2, readings[1].Id)
	assert.Eq(t, "device42", readings[1].Device)

	// Limit
	readings, err = client.ReadingsByDevice("device42", 1)
	assert.Eq(t, 1, len(readings))

	readings, err = client.ReadingsByDevice("device43", 10)
	assert.Eq(t, 1, len(readings))
	assert.Eq(t, objectId3, readings[0].Id)
	assert.Eq(t, "device43", readings[0].Device)
}
