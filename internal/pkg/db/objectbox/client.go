package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	. "github.com/edgexfoundry/edgex-go/internal/pkg/objectbox"
	"github.com/edgexfoundry/edgex-go/pkg/models"
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
	builder := NewObjectBoxBuilder().Name(client.config.DatabaseName).LastEntityId(2, 10002)
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)
	builder.RegisterBinding(EventBinding{})
	builder.RegisterBinding(ReadingBinding{})
	objectBox, err := builder.Build()
	if err != nil {
		return
	}
	client.objectBox = objectBox
	return
}

func (client *ObjectBoxClient) Disconnect() {
	objectBoxToDestroy := client.objectBox
	client.objectBox = nil
	if objectBoxToDestroy != nil {
		objectBoxToDestroy.Destroy()
	}
}

func (client *ObjectBoxClient) Events() (events []models.Event, err error) {
	err = client.objectBox.Strict().RunWithCursor(2, true, func(cursor *Cursor) (err error) {
		slice, err := cursor.GetAll()
		if slice != nil {
			events = slice.([]models.Event)
		}
		return
	})
	return
}

func (client *ObjectBoxClient) AddEvent(event *models.Event) (objectId bson.ObjectId, err error) {
	var id uint64
	if false {
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
		object, err := cursor.Get(uint64(id))
		if object != nil {
			event = *object.(*models.Event)
		}
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
		slice, err := cursor.GetAll()
		if slice != nil {
			readings = slice.([]models.Reading)
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
		object, err := cursor.Get(uint64(id))
		if err != nil {
			return
		}
		reading = *object.(*models.Reading)
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
		for _, bytesData := range bytesArray.BytesArray {
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
