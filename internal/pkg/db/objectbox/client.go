package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/objectbox/objectbox-go/objectbox"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"sync"
)

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox

	eventBox   *Box
	readingBox *Box

	queryEventByDeviceId      *Query
	queryEventByDeviceIdMutex sync.Mutex

	queryReadingByDeviceId      *Query
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
func (client *ObjectBoxClient) eventBoxForReads() *Box {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.eventBox
}

// Considers client.strictReads
func (client *ObjectBoxClient) readingBoxForReads() *Box {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.readingBox
}

func (client *ObjectBoxClient) CloseSession() {
	client.Disconnect()
}

func (client *ObjectBoxClient) Connect() (err error) {
	model := NewModel()
	model.GeneratorVersion(1)
	model.RegisterBinding(EventBinding{})
	model.RegisterBinding(ReadingBinding{indexDevice: false}) // TODO make this configurable once EdgeX allows this
	model.LastEntityId(2, 10002)
	model.LastIndexId(1, 20002007) // Index for Reading.Device

	builder := NewBuilder().Directory(client.config.DatabaseName).Model(model)
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)
	objectBox, err := builder.BuildOrError()
	if err != nil {
		return
	}
	client.objectBox = objectBox
	client.eventBox = objectBox.InternalBox(1)
	client.readingBox = objectBox.InternalBox(2)
	client.asyncPut = true
	client.strictReads = true

	queryBuilder := objectBox.InternalNewQueryBuilder(1)
	queryBuilder.StringEquals(3, "", true)
	client.queryEventByDeviceId, err = queryBuilder.Build()
	if err != nil {
		return
	}

	queryBuilder = objectBox.InternalNewQueryBuilder(2)
	queryBuilder.StringEquals(7, "", true)
	client.queryReadingByDeviceId, err = queryBuilder.Build()
	if err != nil {
		return
	}

	return
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

func (client *ObjectBoxClient) Events() (events []models.Event, err error) {
	slice, err := client.eventBoxForReads().GetAll()
	if slice != nil {
		events = slice.([]models.Event)
	}
	return
}

func (client *ObjectBoxClient) EventsWithLimit(limit int) ([]models.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) AddEvent(event *models.Event) (objectId bson.ObjectId, err error) {
	var id uint64
	if client.asyncPut {
		id, err = client.eventBox.PutAsync(event)
	} else {
		id, err = client.eventBox.Put(event)
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
	object, err := client.eventBoxForReads().Get(id)
	if object != nil {
		event = *object.(*models.Event)
	}
	return
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

func (client *ObjectBoxClient) EventsForDevice(deviceId string) (events []models.Event, err error) {
	client.queryEventByDeviceIdMutex.Lock()
	defer client.queryEventByDeviceIdMutex.Unlock()

	err = client.queryEventByDeviceId.InternalSetParamString(3, deviceId)
	if err != nil {
		return nil, err
	}
	slice, err := client.queryEventByDeviceId.Find()
	events = slice.([]models.Event)
	return
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

func (client *ObjectBoxClient) Readings() (readings []models.Reading, err error) {
	slice, err := client.readingBoxForReads().GetAll()
	if slice != nil {
		readings = slice.([]models.Reading)
	}
	return
}

func (client *ObjectBoxClient) AddReading(r models.Reading) (objectId bson.ObjectId, err error) {
	var id uint64
	if client.asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
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
	object, err := client.readingBoxForReads().Get(id)
	if object == nil || err != nil {
		return
	}
	reading = *object.(*models.Reading)
	return
}

func (client *ObjectBoxClient) ReadingCount() (count int, err error) {
	countLong, err := client.readingBoxForReads().Count()
	count = int(countLong)
	return
}

func (ObjectBoxClient) DeleteReadingById(id string) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) (readings []models.Reading, err error) {
	client.queryReadingByDeviceIdMutex.Lock()
	defer client.queryReadingByDeviceIdMutex.Unlock()

	err = client.queryReadingByDeviceId.InternalSetParamString(7, deviceId)
	if err != nil {
		return nil, err
	}
	slice, err := client.queryReadingByDeviceId.Find()
	readings = slice.([]models.Reading)
	if limit > 0 && limit < len(readings) { // TODO clarify semantics of limit == 0
		// TODO put limit in the query
		readings = readings[:limit]
	}
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

func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}
