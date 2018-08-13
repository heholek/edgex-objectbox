package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/flatcoredata"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox
}

func NewClient(config db.Configuration) *ObjectBoxClient {
	client := &ObjectBoxClient{config: config}
	return client
}

func (ObjectBoxClient) CloseSession() {
	panic("implement me")
}

func (client *ObjectBoxClient) Connect() (err error) {
	model, err := createCoreDataModel()
	if err != nil {
		return
	}
	objectBox, err := NewObjectBox(model, client.config.DatabaseName)
	if err != nil {
		return
	}
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)
	objectBox.RegisterBinding(EventBinding{})
	objectBox.RegisterBinding(ReadingBinding{})
	client.objectBox = objectBox
	return
}

func (client *ObjectBoxClient) Events() (events []models.Event, err error) {
	err = client.objectBox.Strict().RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		var bytes []byte
		for bytes, err = cursor.First(); bytes != nil; bytes, err = cursor.Next() {
			if err != nil || bytes == nil {
				return
			}
			flatEvent := flatcoredata.GetRootAsEvent(bytes, flatbuffers.UOffsetT(0))
			events = append(events, *toModelEvent(flatEvent))
		}
		return
	})
	return
}

func (client *ObjectBoxClient) AddEvent(event *models.Event) (objectId bson.ObjectId, err error) {
	var id uint64
	if true {
		err = client.objectBox.RunWithCursor(1, false, func(cursor *Cursor) (err error) {
			id, err = cursor.Put(event)
			return
		})
	} else {
		var box *Box // explicit to avoid shadowing
		box, err = client.objectBox.Box(1)
		if err != nil {
			return
		}
		id, err = box.PutAsync(event)
	}
	if err != nil {
		return
	}

	stringId := bson.ObjectId(strconv.FormatUint(id, 10))
	event.ID = stringId
	return stringId, nil
}

func (client *ObjectBoxClient) UpdateEvent(e models.Event) error {
	panic("implement me")
}

func (client *ObjectBoxClient) EventById(idString string) (event models.Event, err error) {
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return
	}
	client.objectBox.Strict().RunWithCursor(1, true, func(cursor *Cursor) (err error) {
		bytes, err := cursor.Get(uint64(id))
		if bytes == nil || err != nil {
			return
		}
		flatEvent := flatcoredata.GetRootAsEvent(bytes, flatbuffers.UOffsetT(0))
		event = *toModelEvent(flatEvent)
		return
	})
	return
}

func (client *ObjectBoxClient) EventCount() (count int, err error) {
	err = client.objectBox.Strict().RunWithCursor(1, true, func(cursor *Cursor) (err error) {
		var countLong uint64
		countLong, err = cursor.Count()
		count = int(countLong)
		return
	})
	return
}

func (ObjectBoxClient) EventCountByDeviceId(id string) (int, error) {
	panic("implement me")
}

func (ObjectBoxClient) DeleteEventById(id string) error {
	panic("implement me")
}

func (ObjectBoxClient) EventsForDeviceLimit(id string, limit int) ([]models.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsForDevice(id string) ([]models.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsByCreationTime(startTime, endTime int64, limit int) ([]models.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]models.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsOlderThanAge(age int64) ([]models.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsPushed() ([]models.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) ScrubAllEvents() (err error) {
	err = client.objectBox.RunWithCursor(2, false, func(cursor *Cursor) (err error) {
		return cursor.RemoveAll()
	})
	if err != nil {
		return
	}
	return client.objectBox.RunWithCursor(1, false, func(cursor *Cursor) (err error) {
		return cursor.RemoveAll()
	})
}

func (client *ObjectBoxClient) Readings() (readings []models.Reading, err error) {
	err = client.objectBox.Strict().RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		var bytes []byte
		for bytes, err = cursor.First(); bytes != nil; bytes, err = cursor.Next() {
			if err != nil || bytes == nil {
				return
			}
			flatReading := flatcoredata.GetRootAsReading(bytes, flatbuffers.UOffsetT(0))
			readings = append(readings, *toModelReading(flatReading))
		}
		return
	})
	return
}

func (client *ObjectBoxClient) AddReading(r models.Reading) (objectId bson.ObjectId, err error) {
	var id uint64
	if false {
		err = client.objectBox.RunWithCursor(2, false, func(cursor *Cursor) (err error) {
			id, err = cursor.Put(&r)
			return
		})
	} else {
		var box *Box // Explicit to avoid shadowing
		box, err = client.objectBox.Box(2)
		if err != nil {
			return
		}
		id, err = box.PutAsync(&r)
	}
	if err != nil {
		return
	}
	stringId := bson.ObjectId(strconv.FormatUint(id, 10))
	r.Id = stringId
	return stringId, nil
}

func (ObjectBoxClient) UpdateReading(r models.Reading) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingById(idString string) (reading models.Reading, err error) {
	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		return
	}
	client.objectBox.RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		bytes, err := cursor.Get(uint64(id))
		if err != nil {
			return
		}
		reading = *toModelReadingFromBytes(bytes)
		return
	})
	return
}

func (client *ObjectBoxClient) ReadingCount() (count int, err error) {
	err = client.objectBox.Strict().RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		var countLong uint64
		countLong, err = cursor.Count()
		count = int(countLong)
		return
	})
	return
}

func (ObjectBoxClient) DeleteReadingById(id string) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) (readings []models.Reading, err error) {
	client.objectBox.Strict().RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		bytesArray, err := cursor.FindByString(7, deviceId)
		if err != nil {
			return
		}
		defer bytesArray.Destroy()
		for _, bytesData := range bytesArray.bytesArray {
			readings = append(readings, *toModelReadingFromBytes(bytesData))
			if len(readings) == limit {
				// TODO consider limit in query builder
				break
			}
		}
		return
	})
	return
}

func (ObjectBoxClient) ReadingsByValueDescriptor(name string, limit int) ([]models.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]models.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByCreationTime(start, end int64, limit int) ([]models.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) AddValueDescriptor(v models.ValueDescriptor) (bson.ObjectId, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptors() ([]models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) UpdateValueDescriptor(v models.ValueDescriptor) error {
	panic("implement me")
}

func (ObjectBoxClient) DeleteValueDescriptorById(id string) error {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorByName(name string) (models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByName(names []string) ([]models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorById(id string) (models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByUomLabel(uomLabel string) ([]models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByLabel(label string) ([]models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByType(t string) ([]models.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ScrubAllValueDescriptors() error {
	panic("implement me")
}

// TODO this is rather a quick hack to make it work, clean up later
func createCoreDataModel() (model *Model, err error) {
	model, err = NewModel()
	if err != nil {
		return
	}

	model.Entity("Event", 1, 10001)
	model.Property("id", PropertyType_Long, 1, 10001001)
	model.PropertyFlags(PropertyFlags_ID)
	model.Property("pushed", PropertyType_Long, 2, 10001002)
	model.Property("device", PropertyType_String, 3, 10001003)
	model.Property("created", PropertyType_Long, 4, 10001004)
	model.Property("modified", PropertyType_Long, 5, 10001005)
	model.Property("origin", PropertyType_Long, 6, 10001006)
	model.Property("scheduleEvent", PropertyType_String, 7, 10001007)
	model.EntityLastPropertyId(7, 10001007)

	model.Entity("Reading", 2, 10002)
	model.Property("id", PropertyType_Long, 1, 10002001)
	model.PropertyFlags(PropertyFlags_ID)
	model.Property("eventId", PropertyType_Long, 2, 10002002)
	//model.Property("eventId", PropertyType_Relation, 2, 10002002)
	//model.PropertyFlags(PropertyFlags_INDEXED)
	model.Property("pushed", PropertyType_Long, 3, 10002003)
	model.Property("created", PropertyType_Long, 4, 10002004)
	model.Property("origin", PropertyType_Long, 5, 10002005)
	model.Property("modified", PropertyType_Long, 6, 10002006)
	model.Property("device", PropertyType_String, 7, 10002007)
	model.Property("name", PropertyType_String, 8, 10002008)
	model.Property("value", PropertyType_String, 9, 10002009)
	model.EntityLastPropertyId(9, 10002009)

	model.LastEntityId(2, 10002)

	if model.err != nil {
		err = model.err
		model = nil
		return
	}

	return
}
