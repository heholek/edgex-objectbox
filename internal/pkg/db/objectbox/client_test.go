package objectbox

import (
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/test"
	"os"
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/models"
	correlation "github.com/objectbox/edgex-objectbox/internal/pkg/correlation/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/objectbox-go/test/assert"
)

func TestObjectBox(t *testing.T) {
	var config = db.Configuration{DatabaseName: "testdata"}

	defer os.RemoveAll(config.DatabaseName)

	if client, err := NewClient(config); err != nil {
		t.Fatalf("Could not connect: %v", err)
	} else {
		test.TestDataDB(t, client)
	}

	if client, err := NewClient(config); err != nil {
		t.Fatalf("Could not connect: %v", err)
	} else {
		test.TestMetadataDB(t, client)
	}

	if client, err := NewClient(config); err != nil {
		t.Fatalf("Could not connect: %v", err)
	} else {
		test.TestExportDB(t, client)
	}

	if client, err := NewClient(config); err != nil {
		t.Fatalf("Could not connect: %v", err)
	} else {
		test.TestSchedulerDB(t, client)
	}

	if client, err := NewClient(config); err != nil {
		t.Fatalf("Could not connect: %v", err)
	} else {
		test.TestNotificationsDB(t, client)
	}
}

func createClient() *ObjectBoxClient {
	config := db.Configuration{
		DatabaseName: "unit-test",
	}
	os.RemoveAll(config.DatabaseName)
	client, err := NewClient(config)
	if err != nil {
		panic("Could not connect DB client: " + err.Error())
	}
	return client
}

func TestObjectBoxEvents(t *testing.T) {
	client := createClient()
	defer client.CloseSession()

	event := correlation.Event{
		Event: models.Event{
			Device: "my device",
		},
	}
	objectId, err := client.AddEvent(event)
	assert.NoErr(t, err)

	event.Device = "2nd device"
	objectId, err = client.AddEvent(event)
	assert.NoErr(t, err)

	eventRead, err := client.EventById(string(objectId))
	assert.NoErr(t, err)

	if objectId != eventRead.ID || event.Device != eventRead.Device {
		t.Fatalf("Event data error: %v vs. %v", event, eventRead)
	}

	allEvents, err := client.Events()
	assert.NoErr(t, err)
	assert.Eq(t, 2, len(allEvents))

	allEvents, err = client.EventsWithLimit(10)
	assert.NoErr(t, err)
	assert.Eq(t, 2, len(allEvents))

	allEvents, err = client.EventsWithLimit(1)
	assert.NoErr(t, err)
	assert.Eq(t, 1, len(allEvents))
}

func TestObjectBoxReadings(t *testing.T) {
	client := createClient()
	defer client.CloseSession()

	assert.NoErr(t, client.ScrubAllEvents())
	countPre, err := client.eventBox.Count()
	assert.NoErr(t, err)
	assert.Eq(t, countPre, uint64(0))
	countPre, err = client.readingBox.Count()
	assert.Eq(t, countPre, uint64(0))

	reading := models.Reading{
		Name:   "reading1",
		Device: "device42",
	}
	objectId, err := client.AddReading(reading)
	assert.NoErr(t, err)
	t.Logf("Added reading ID %v", objectId)
	assert.NotEq(t, objectId, "")

	reading.Name = "reading2"
	objectId2, err := client.AddReading(reading)
	assert.NoErr(t, err)

	t.Logf("Added 2nd reading ID %v", objectId2)

	reading.Name = "reading3"
	reading.Device = "device43"
	objectId3, err := client.AddReading(reading)
	assert.NoErr(t, err)

	t.Logf("Added 3rd reading ID %v", objectId3)
	count, err := client.ReadingCount()
	assert.NoErr(t, err)

	assert.Eq(t, count, 3)

	readingRead, err := client.ReadingById(string(objectId2))
	if err != nil {
		t.Fatalf("Could not get 2nd reading by ID: %v", err)
	}
	assert.Eq(t, readingRead.Id, objectId2)
	assert.Eq(t, readingRead.Device, "device42")

	all, err := client.Readings()
	assert.NoErr(t, err)
	assert.Eq(t, len(all), 3)
	assert.Eq(t, all[0].Device, "device42")
	assert.Eq(t, all[1].Device, "device42")
	assert.Eq(t, all[2].Device, "device43")

	readings, err := client.ReadingsByDevice("device42", 10)
	assert.Eq(t, len(readings), 2)
	assert.Eq(t, readings[0].Id, objectId)
	assert.Eq(t, readings[0].Device, "device42")
	assert.Eq(t, readings[1].Id, objectId2)
	assert.Eq(t, readings[1].Device, "device42")

	// Limit
	readings, err = client.ReadingsByDevice("device42", 1)
	assert.Eq(t, len(readings), 1)

	readings, err = client.ReadingsByDevice("device43", 10)
	assert.Eq(t, len(readings), 1)
	assert.Eq(t, readings[0].Id, objectId3)
	assert.Eq(t, readings[0].Device, "device43")
}
