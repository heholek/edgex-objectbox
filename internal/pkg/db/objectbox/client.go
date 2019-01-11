package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	. "github.com/objectbox/objectbox-go/objectbox"
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

func (client *ObjectBoxClient) Events() ([]contract.Event, error) {
	return client.eventBoxForReads().GetAll()
}

func (client *ObjectBoxClient) EventsWithLimit(limit int) ([]contract.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) AddEvent(event contract.Event) (objectId string, err error) {
	var id uint64
	if client.asyncPut {
		id, err = client.eventBox.PutAsync(&event)
	} else {
		id, err = client.eventBox.Put(&event)
	}
	if err != nil {
		return
	}

	event.ID = idObxToString(id)
	return event.ID, nil
}

func (client *ObjectBoxClient) UpdateEvent(e contract.Event) error {
	panic("implement me")
}

func (client *ObjectBoxClient) EventById(idString string) (contract.Event, error) {
	id, err := idStringToObx(idString)
	if err != nil {
		return contract.Event{}, err
	}

	object, err := client.eventBoxForReads().Get(id)
	if object == nil || err != nil {
		return contract.Event{}, err
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

func (ObjectBoxClient) EventsForDeviceLimit(id string, limit int) ([]contract.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) EventsForDevice(deviceId string) ([]contract.Event, error) {
	client.queryEventByDeviceIdMutex.Lock()
	defer client.queryEventByDeviceIdMutex.Unlock()
	if err := client.queryEventByDeviceId.SetStringParams(obx.Event_.Device, deviceId); err != nil {
		return nil, err
	}
	return client.queryEventByDeviceId.Find()
}

func (ObjectBoxClient) EventsByCreationTime(startTime, endTime int64, limit int) ([]contract.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]contract.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsOlderThanAge(age int64) ([]contract.Event, error) {
	panic("implement me")
}

func (ObjectBoxClient) EventsPushed() ([]contract.Event, error) {
	panic("implement me")
}

func (client *ObjectBoxClient) ScrubAllEvents() (err error) {
	err = client.eventBox.RemoveAll()
	if err != nil {
		return
	}
	return client.readingBoxForReads().RemoveAll()
}

func (client *ObjectBoxClient) Readings() ([]contract.Reading, error) {
	return client.readingBoxForReads().GetAll()
}

func (client *ObjectBoxClient) AddReading(r contract.Reading) (objectId string, err error) {
	var id uint64
	if client.asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}
	if err != nil {
		return
	}
	r.Id = idObxToString(id)
	return r.Id, nil
}

func (ObjectBoxClient) UpdateReading(r contract.Reading) error {
	panic("implement me")
}

func (client *ObjectBoxClient) ReadingById(idString string) (contract.Reading, error) {
	id, err := idStringToObx(idString)
	if err != nil {
		return contract.Reading{}, err
	}

	object, err := client.readingBoxForReads().Get(id)
	if object == nil || err != nil {
		return contract.Reading{}, err
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

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) ([]contract.Reading, error) {
	client.queryReadingByDeviceIdMutex.Lock()
	defer client.queryReadingByDeviceIdMutex.Unlock()
	if err := client.queryReadingByDeviceId.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, err
	}
	return client.queryReadingByDeviceId.Limit(uint64(limit)).Find()
}

func (ObjectBoxClient) ReadingsByValueDescriptor(name string, limit int) ([]contract.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]contract.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) ReadingsByCreationTime(start, end int64, limit int) ([]contract.Reading, error) {
	panic("implement me")
}

func (ObjectBoxClient) AddValueDescriptor(v contract.ValueDescriptor) (string, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptors() ([]contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) UpdateValueDescriptor(v contract.ValueDescriptor) error {
	panic("implement me")
}

func (ObjectBoxClient) DeleteValueDescriptorById(id string) error {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorByName(name string) (contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByName(names []string) ([]contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorById(id string) (contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByUomLabel(uomLabel string) ([]contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByLabel(label string) ([]contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ValueDescriptorsByType(t string) ([]contract.ValueDescriptor, error) {
	panic("implement me")
}

func (ObjectBoxClient) ScrubAllValueDescriptors() error {
	panic("implement me")
}

func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}
