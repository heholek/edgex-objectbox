package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/objectbox/objectbox-go/objectbox"
	"gopkg.in/mgo.v2/bson"
	"sync"
)

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox

	eventBox   *obx.EventBox
	readingBox *obx.ReadingBox

	queryEventByDeviceId      *obx.EventQuery
	queryEventByDeviceIdMutex sync.Mutex

	queryReadingByDeviceId      *obx.ReadingQuery
	queryReadingByDeviceIdMutex sync.Mutex

	strictReads bool
	asyncPut    bool
}

func NewClient(config db.Configuration) *ObjectBoxClient {
	println(VersionInfo())
	client := &ObjectBoxClient{config: config}
	return client
}

// Considers client.strictReads
func (client *ObjectBoxClient) storeForReads() *ObjectBox {
	store := client.objectBox
	if client.strictReads {
		store.AwaitAsyncCompletion()
	}
	return store
}

// Considers client.strictReads
func (client *ObjectBoxClient) eventBoxForReads() *obx.EventBox {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.eventBox
}

// Considers client.strictReads
func (client *ObjectBoxClient) readingBoxForReads() *obx.ReadingBox {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.readingBox
}

func (client *ObjectBoxClient) CloseSession() {
	client.Disconnect()
}

func (client *ObjectBoxClient) Connect() error {
	objectBox, err := NewBuilder().Directory(client.config.DatabaseName).Model(obx.ObjectBoxModel()).Build()
	if err != nil {
		return err
	}
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)

	client.objectBox = objectBox
	client.eventBox = obx.BoxForEvent(objectBox)
	client.readingBox = obx.BoxForReading(objectBox)
	client.asyncPut = true
	client.strictReads = true

	client.queryEventByDeviceId, err = client.eventBox.QueryOrError(obx.Event_.Device.Equals("", true))
	if err != nil {
		return err
	}

	client.queryReadingByDeviceId, err = client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true))
	if err != nil {
		return err
	}

	return err
}

func (client *ObjectBoxClient) Disconnect() {
	client.eventBox = nil
	client.readingBox = nil
	objectBoxToDestroy := client.objectBox
	client.objectBox = nil
	if objectBoxToDestroy != nil {
		objectBoxToDestroy.Close()
	}
}

func (client *ObjectBoxClient) Events() ([]models.Event, error) {
	slice, err := client.eventBoxForReads().GetAll()
	if err != nil {
		return nil, err
	}

	// TODO this needs to be done by the binding
	var events = make([]models.Event, 0, len(slice))
	for _, ptr := range slice {
		events = append(events, *ptr)
	}

	return events, nil
}

func (client *ObjectBoxClient) EventsWithLimit(limit int) ([]models.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) AddEvent(event *models.Event) (objectId bson.ObjectId, err error) {
	var id string
	if client.asyncPut {
		id, err = client.eventBox.PutAsync(event)
	} else {
		id, err = client.eventBox.Put(event)
	}
	if err != nil {
		return
	}

	event.ID = bson.ObjectId(id)
	return event.ID, nil
}

func (client *ObjectBoxClient) UpdateEvent(e models.Event) error {
	panic("implement me")
}

func (client *ObjectBoxClient) EventById(idString string) (models.Event, error) {
	object, err := client.eventBoxForReads().Get(idString)
	if object == nil || err != nil {
		return models.Event{}, err
	}
	return *object, nil
}

func (client *ObjectBoxClient) EventCount() (count int, err error) {
	countLong, err := client.eventBoxForReads().Count()
	if err == nil {
		count = int(countLong)
	}
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

func (client *ObjectBoxClient) EventsForDevice(deviceId string) ([]models.Event, error) {
	client.queryEventByDeviceIdMutex.Lock()
	client.queryEventByDeviceId.InternalSetParamString(obx.Event_.Device.Id, deviceId)
	slice, err := client.queryEventByDeviceId.Find()
	client.queryEventByDeviceIdMutex.Unlock()

	if err != nil {
		return nil, err
	}

	// TODO this needs to be done by the binding
	var events = make([]models.Event, 0, len(slice))
	for _, ptr := range slice {
		events = append(events, *ptr)
	}

	return events, nil
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
	err = client.eventBox.RemoveAll()
	if err != nil {
		return
	}
	return client.readingBoxForReads().RemoveAll()
}

func (client *ObjectBoxClient) Readings() ([]models.Reading, error) {
	slice, err := client.readingBoxForReads().GetAll()
	if err != nil {
		return nil, err
	}

	// TODO this needs to be done by the binding
	var readings = make([]models.Reading, 0, len(slice))
	for _, ptr := range slice {
		readings = append(readings, *ptr)
	}

	return readings, nil
}

func (client *ObjectBoxClient) AddReading(r models.Reading) (objectId bson.ObjectId, err error) {
	var id string
	if client.asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}
	if err != nil {
		return
	}
	r.Id = bson.ObjectId(id)
	return r.Id, nil
}

func (ObjectBoxClient) UpdateReading(r models.Reading) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingById(idString string) (models.Reading, error) {
	object, err := client.readingBoxForReads().Get(idString)
	if object == nil || err != nil {
		return models.Reading{}, err
	}
	return *object, nil
}

func (client *ObjectBoxClient) ReadingCount() (count int, err error) {
	countLong, err := client.readingBoxForReads().Count()
	count = int(countLong)
	return
}

func (ObjectBoxClient) DeleteReadingById(id string) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) ([]models.Reading, error) {
	client.queryReadingByDeviceIdMutex.Lock()
	client.queryReadingByDeviceId.InternalSetParamString(obx.Reading_.Device.Id, deviceId)
	slice, err := client.queryReadingByDeviceId.Find()
	client.queryReadingByDeviceIdMutex.Unlock()

	if err != nil {
		return nil, err
	}

	// TODO this needs to be done by the binding
	var readings = make([]models.Reading, 0, len(slice))
	for _, ptr := range slice {
		readings = append(readings, *ptr)
	}

	return readings, nil
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

func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}
