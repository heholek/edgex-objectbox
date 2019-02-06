package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/test"
	"os"
	"testing"

	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/influxdata/influxdb/pkg/testing/assert"
)

func TestObjectBox(t *testing.T) {
	var config = db.Configuration{DatabaseName: "testdata"}

	defer os.RemoveAll(config.DatabaseName)

	//if client, err := NewClient(config); err != nil {
	//	t.Fatalf("Could not connect: %v", err)
	//} else {
	//	test.TestDataDB(t, client)
	//}
	//
	//if client, err := NewClient(config); err != nil {
	//	t.Fatalf("Could not connect: %v", err)
	//} else {
	//	test.TestMetadataDB(t, client)
	//}
	//
	//if client, err := NewClient(config); err != nil {
	//	t.Fatalf("Could not connect: %v", err)
	//} else {
	//	test.TestExportDB(t, client)
	//}
	//
	//if client, err := NewClient(config); err != nil {
	//	t.Fatalf("Could not connect: %v", err)
	//} else {
	//	test.TestSchedulerDB(t, client)
	//}

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
	client, err := NewClient(config)
	if err != nil {
		panic("Could not connect DB client: " + err.Error())
	}
	return client
}

func TestObjectBoxEvents(t *testing.T) {
	client := createClient()
	defer client.Disconnect()

	event := models.Event{
		Device: "my device",
	}
	objectId, err := client.AddEvent(event)
	if err != nil {
		t.Fatalf("Could not add event: %v", err)
	}
	t.Logf("Added object ID %v", objectId)

	event.Device = "2nd device"
	objectId, err = client.AddEvent(event)
	if err != nil {
		t.Fatalf("Could not add 2nd event: %v", err)
	}
	t.Logf("Added 2nd object ID %v", objectId)

	eventRead, err := client.EventById(string(objectId))
	if err != nil {
		t.Fatalf("Could not get 2nd event by ID: %v", err)
	}
	if objectId != eventRead.ID || event.Device != eventRead.Device {
		t.Fatalf("Event data error: %v vs. %v", event, eventRead)
	}
}

func TestObjectBoxReadings(t *testing.T) {
	client := createClient()
	defer client.Disconnect()

	assert.NoError(t, client.ScrubAllEvents())
	countPre, err := client.eventBox.Count()
	assert.NoError(t, err)
	ok := assert.Equal(t, countPre, uint64(0))
	countPre, err = client.readingBox.Count()
	ok = ok && assert.Equal(t, countPre, uint64(0))
	if !ok {
		t.Fatal("Non-zero counts")
	}

	reading := models.Reading{
		Name:   "reading1",
		Device: "device42",
	}
	objectId, err := client.AddReading(reading)
	assert.NoError(t, err)
	t.Logf("Added reading ID %v", objectId)
	assert.NotEqual(t, objectId, "")

	reading.Name = "reading2"
	objectId2, err := client.AddReading(reading)
	assert.NoError(t, err)

	t.Logf("Added 2nd reading ID %v", objectId2)

	reading.Name = "reading3"
	reading.Device = "device43"
	objectId3, err := client.AddReading(reading)
	assert.NoError(t, err)

	t.Logf("Added 3rd reading ID %v", objectId3)
	count, err := client.ReadingCount()
	assert.NoError(t, err)

	assert.Equal(t, count, 3)

	readingRead, err := client.ReadingById(string(objectId2))
	if err != nil {
		t.Fatalf("Could not get 2nd reading by ID: %v", err)
	}
	assert.Equal(t, readingRead.Id, objectId2)
	assert.Equal(t, readingRead.Device, "device42")

	all, err := client.Readings()
	assert.NoError(t, err)
	if assert.Equal(t, len(all), 3) {
		assert.Equal(t, all[0].Device, "device42")
		assert.Equal(t, all[1].Device, "device42")
		assert.Equal(t, all[2].Device, "device43")
	}

	readings, err := client.ReadingsByDevice("device42", 10)
	if assert.Equal(t, len(readings), 2) {
		assert.Equal(t, readings[0].Id, objectId)
		assert.Equal(t, readings[0].Device, "device42")
		assert.Equal(t, readings[1].Id, objectId2)
		assert.Equal(t, readings[1].Device, "device42")
	}

	// Limit
	readings, err = client.ReadingsByDevice("device42", 1)
	assert.Equal(t, len(readings), 1)

	readings, err = client.ReadingsByDevice("device43", 10)
	if assert.Equal(t, len(readings), 1) {
		assert.Equal(t, readings[0].Id, objectId3)
		assert.Equal(t, readings[0].Device, "device43")
	}
}
