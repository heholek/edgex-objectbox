package objectbox

// implements core-data service contract
// TODO queries are not "async-put safe", i. e. there might be changes that have not been written
// TODO indexes

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"sync"
)

//region Queries
type coreDataQueries struct {
	events struct {
		all       eventQuery
		createdB  eventQuery
		createdLT eventQuery
		device    eventQuery
		pushedGT  eventQuery
	}
	readings struct {
		createdB      readingQuery
		device        readingQuery
		deviceAndName readingQuery
		name          readingQuery
		names         readingQuery
	}
}

type eventQuery struct {
	*obx.EventQuery
	sync.Mutex
}

type readingQuery struct {
	*obx.ReadingQuery
	sync.Mutex
}

//endregion

func (client *ObjectBoxClient) initCoreData() error {
	var err error

	//region Events
	if err == nil {
		client.queries.events.all.EventQuery, err = client.eventBox.QueryOrError()
	}

	if err == nil {
		client.queries.events.device.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Device.Equals("", true))
	}

	if err == nil {
		client.queries.events.createdB.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Created.Between(0, 0))
	}

	if err == nil {
		client.queries.events.createdLT.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Created.LessThan(0))
	}

	if err == nil {
		client.queries.events.pushedGT.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Pushed.GreaterThan(0))
	}
	//endregion

	//region Readings
	if err == nil {
		client.queries.readings.device.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true))
	}

	if err == nil {
		client.queries.readings.deviceAndName.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true), obx.Reading_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.readings.name.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.readings.names.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Name.In(true))
	}

	if err == nil {
		client.queries.readings.createdB.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Created.Between(0, 0))
	}
	//endregion

	return err
}

func (client *ObjectBoxClient) Events() ([]contract.Event, error) {
	return client.eventBoxForReads().GetAll()
}

func (client *ObjectBoxClient) EventsWithLimit(limit int) ([]contract.Event, error) {
	// TODO there is no test for this method in the test/db_data.go
	var query = &client.queries.events.all

	query.Lock()
	defer query.Unlock()

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) AddEvent(event contract.Event) (objectId string, err error) {
	if event.Created == 0 {
		event.Created = db.MakeTimestamp()
	}

	// TODO readings

	var id uint64
	if client.asyncPut {
		id, err = client.eventBox.PutAsync(&event)
	} else {
		id, err = client.eventBox.Put(&event)
	}
	if err != nil {
		return
	}

	event.ID = obx.IdToString(id)
	return event.ID, nil
}

func (client *ObjectBoxClient) UpdateEvent(e contract.Event) error {
	e.Modified = db.MakeTimestamp()

	// check whether it exists, otherwise this function must fail
	if object, err := client.eventById(e.ID); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	var err error
	if client.asyncPut {
		_, err = client.eventBox.PutAsync(&e)
	} else {
		_, err = client.eventBox.Put(&e)
	}

	return err
}

func (client *ObjectBoxClient) EventById(idString string) (contract.Event, error) {
	object, err := client.eventById(idString)
	if object == nil || err != nil {
		return contract.Event{}, err
	}
	return *object, nil
}

func (client *ObjectBoxClient) eventById(idString string) (*contract.Event, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.eventBoxForReads().Get(id)
}

func (client *ObjectBoxClient) EventCount() (count int, err error) {
	countLong, err := client.eventBoxForReads().Count()
	if err == nil {
		count = int(countLong)
	}
	return
}

func (client *ObjectBoxClient) EventCountByDeviceId(id string) (int, error) {
	var query = &client.queries.events.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return 0, err
	}

	count, err := query.Count()
	return int(count), err
}

func (client *ObjectBoxClient) DeleteEventById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	client.objectBox.AwaitAsyncCompletion()
	return client.eventBox.Box.Remove(id)
}

func (client *ObjectBoxClient) EventsForDeviceLimit(id string, limit int) ([]contract.Event, error) {
	var query = &client.queries.events.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) EventsForDevice(id string) ([]contract.Event, error) {
	return client.EventsForDeviceLimit(id, 0)
}

func (client *ObjectBoxClient) EventsByCreationTime(start, end int64, limit int) ([]contract.Event, error) {
	var query = &client.queries.events.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) EventsOlderThanAge(age int64) ([]contract.Event, error) {
	var time = (db.MakeTimestamp()) - age

	var query = &client.queries.events.createdLT

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, time); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *ObjectBoxClient) EventsPushed() ([]contract.Event, error) {
	var query = &client.queries.events.pushedGT

	query.Lock()
	defer query.Unlock()

	return query.Find()
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
	if r.Created == 0 {
		r.Created = db.MakeTimestamp()
	}

	var id uint64
	if client.asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}
	if err != nil {
		return
	}
	r.Id = obx.IdToString(id)
	return r.Id, nil
}

func (client *ObjectBoxClient) UpdateReading(r contract.Reading) error {
	r.Modified = db.MakeTimestamp()

	// check whether it exists, otherwise this function must fail
	if object, err := client.readingById(r.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	var err error
	if client.asyncPut {
		_, err = client.readingBox.PutAsync(&r)
	} else {
		_, err = client.readingBox.Put(&r)
	}

	return err
}

func (client *ObjectBoxClient) ReadingById(idString string) (contract.Reading, error) {
	object, err := client.readingById(idString)
	if object == nil || err != nil {
		return contract.Reading{}, err
	}
	return *object, nil
}

func (client *ObjectBoxClient) readingById(idString string) (*contract.Reading, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.readingBoxForReads().Get(id)
}

func (client *ObjectBoxClient) ReadingCount() (count int, err error) {
	countLong, err := client.readingBoxForReads().Count()
	count = int(countLong)
	return
}

func (client *ObjectBoxClient) DeleteReadingById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	client.objectBox.AwaitAsyncCompletion()
	return client.readingBox.Box.Remove(id)
}

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.readings.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByValueDescriptor(name string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.readings.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Name, name); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.readings.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.Reading_.Name, names...); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByCreationTime(start, end int64, limit int) ([]contract.Reading, error) {
	var query = &client.queries.readings.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Reading_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.readings.deviceAndName

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, err
	}
	if err := query.SetStringParams(obx.Reading_.Name, valueDescriptor); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (ObjectBoxClient) AddValueDescriptor(v contract.ValueDescriptor) (string, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptors() ([]contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) UpdateValueDescriptor(v contract.ValueDescriptor) error {
	panic(notImplemented())
}

func (ObjectBoxClient) DeleteValueDescriptorById(id string) error {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorByName(name string) (contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorsByName(names []string) ([]contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorById(id string) (contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorsByUomLabel(uomLabel string) ([]contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorsByLabel(label string) ([]contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ValueDescriptorsByType(t string) ([]contract.ValueDescriptor, error) {
	panic(notImplemented())
}

func (ObjectBoxClient) ScrubAllValueDescriptors() error {
	panic(notImplemented())
}

func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}

//endregion
