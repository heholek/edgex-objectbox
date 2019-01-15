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
	event struct {
		all       eventQuery
		createdB  eventQuery
		createdLT eventQuery
		device    eventQuery
		pushedGT  eventQuery
	}
	reading struct {
		createdB      readingQuery
		device        readingQuery
		deviceAndName readingQuery
		name          readingQuery
		names         readingQuery
	}
	valueDescriptor struct {
		name     valueDescriptorQuery
		names    valueDescriptorQuery
		typ      valueDescriptorQuery
		uomlabel valueDescriptorQuery
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

type valueDescriptorQuery struct {
	*obx.ValueDescriptorQuery
	sync.Mutex
}

//endregion

func (client *ObjectBoxClient) initCoreData() error {
	var err error

	//region Event
	if err == nil {
		client.queries.event.all.EventQuery, err = client.eventBox.QueryOrError()
	}

	if err == nil {
		client.queries.event.device.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Device.Equals("", true))
	}

	if err == nil {
		client.queries.event.createdB.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Created.Between(0, 0))
	}

	if err == nil {
		client.queries.event.createdLT.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Created.LessThan(0))
	}

	if err == nil {
		client.queries.event.pushedGT.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Pushed.GreaterThan(0))
	}
	//endregion

	//region Reading
	if err == nil {
		client.queries.reading.device.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true))
	}

	if err == nil {
		client.queries.reading.deviceAndName.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true), obx.Reading_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.reading.name.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.reading.names.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Name.In(true))
	}

	if err == nil {
		client.queries.reading.createdB.ReadingQuery, err =
			client.readingBox.QueryOrError(obx.Reading_.Created.Between(0, 0))
	}
	//endregion

	//region ValueDescriptor
	if err == nil {
		client.queries.valueDescriptor.name.ValueDescriptorQuery, err =
			client.valueDescriptorBox.QueryOrError(obx.ValueDescriptor_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.valueDescriptor.names.ValueDescriptorQuery, err =
			client.valueDescriptorBox.QueryOrError(obx.ValueDescriptor_.Name.In(true))
	}

	if err == nil {
		client.queries.valueDescriptor.typ.ValueDescriptorQuery, err =
			client.valueDescriptorBox.QueryOrError(obx.ValueDescriptor_.Type.Equals("", true))
	}

	if err == nil {
		client.queries.valueDescriptor.uomlabel.ValueDescriptorQuery, err =
			client.valueDescriptorBox.QueryOrError(obx.ValueDescriptor_.UomLabel.Equals("", true))
	}
	//endregion

	return err
}

func (client *ObjectBoxClient) Events() ([]contract.Event, error) {
	return client.eventBoxForReads().GetAll()
}

func (client *ObjectBoxClient) EventsWithLimit(limit int) ([]contract.Event, error) {
	// TODO there is no test for this method in the test/db_data.go
	var query = &client.queries.event.all

	query.Lock()
	defer query.Unlock()

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) AddEvent(event contract.Event) (string, error) {
	if event.Created == 0 {
		event.Created = db.MakeTimestamp()
	}

	// TODO readings

	var id uint64
	var err error

	if client.asyncPut {
		id, err = client.eventBox.PutAsync(&event)
	} else {
		id, err = client.eventBox.Put(&event)
	}

	return obx.IdToString(id), err
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

func (client *ObjectBoxClient) EventById(id string) (contract.Event, error) {
	object, err := client.eventById(id)
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
	var query = &client.queries.event.device

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

	return client.eventBoxForReads().Box.Remove(id)
}

func (client *ObjectBoxClient) EventsForDeviceLimit(id string, limit int) ([]contract.Event, error) {
	var query = &client.queries.event.device

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
	var query = &client.queries.event.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) EventsOlderThanAge(age int64) ([]contract.Event, error) {
	var time = (db.MakeTimestamp()) - age

	var query = &client.queries.event.createdLT

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, time); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *ObjectBoxClient) EventsPushed() ([]contract.Event, error) {
	var query = &client.queries.event.pushedGT

	query.Lock()
	defer query.Unlock()

	return query.Find()
}

func (client *ObjectBoxClient) ScrubAllEvents() error {
	if err := client.eventBox.RemoveAll(); err != nil {
		return err
	}
	return client.readingBoxForReads().RemoveAll()
}

func (client *ObjectBoxClient) Readings() ([]contract.Reading, error) {
	return client.readingBoxForReads().GetAll()
}

func (client *ObjectBoxClient) AddReading(r contract.Reading) (string, error) {
	if r.Created == 0 {
		r.Created = db.MakeTimestamp()
	}

	var id uint64
	var err error

	if client.asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}

	return obx.IdToString(id), err
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

func (client *ObjectBoxClient) ReadingById(id string) (contract.Reading, error) {
	object, err := client.readingById(id)
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

func (client *ObjectBoxClient) ReadingCount() (int, error) {
	count, err := client.readingBoxForReads().Count()
	return int(count), err
}

func (client *ObjectBoxClient) DeleteReadingById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.readingBoxForReads().Box.Remove(id)
}

func (client *ObjectBoxClient) ReadingsByDevice(deviceId string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByValueDescriptor(name string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Name, name); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.Reading_.Name, names...); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByCreationTime(start, end int64, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Reading_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *ObjectBoxClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.deviceAndName

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

func (client *ObjectBoxClient) AddValueDescriptor(v contract.ValueDescriptor) (string, error) {
	if v.Created == 0 {
		v.Created = db.MakeTimestamp()
	}

	var id uint64
	var err error

	if client.asyncPut {
		id, err = client.valueDescriptorBox.PutAsync(&v)
	} else {
		id, err = client.valueDescriptorBox.Put(&v)
	}

	return obx.IdToString(id), err
}

func (client *ObjectBoxClient) ValueDescriptors() ([]contract.ValueDescriptor, error) {
	return client.valueDescriptorBoxForReads().GetAll()
}

func (client *ObjectBoxClient) UpdateValueDescriptor(v contract.ValueDescriptor) error {
	v.Modified = db.MakeTimestamp()

	// check whether it exists, otherwise this function must fail
	if object, err := client.valueDescriptorById(v.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	var err error
	if client.asyncPut {
		_, err = client.valueDescriptorBox.PutAsync(&v)
	} else {
		_, err = client.valueDescriptorBox.Put(&v)
	}

	return err
}

func (client *ObjectBoxClient) DeleteValueDescriptorById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.valueDescriptorBoxForReads().Box.Remove(id)
}

func (client *ObjectBoxClient) ValueDescriptorByName(name string) (contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Name, name); err != nil {
		return contract.ValueDescriptor{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.ValueDescriptor{}, err
	} else if len(list) == 0 {
		return contract.ValueDescriptor{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *ObjectBoxClient) ValueDescriptorsByName(names []string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.ValueDescriptor_.Name, names...); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *ObjectBoxClient) ValueDescriptorById(id string) (contract.ValueDescriptor, error) {
	object, err := client.valueDescriptorById(id)
	if object == nil || err != nil {
		return contract.ValueDescriptor{}, err
	}
	return *object, nil
}

func (client *ObjectBoxClient) valueDescriptorById(idString string) (*contract.ValueDescriptor, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.valueDescriptorBoxForReads().Get(id)
}

func (client *ObjectBoxClient) ValueDescriptorsByUomLabel(uomLabel string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.uomlabel

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.UomLabel, uomLabel); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *ObjectBoxClient) ValueDescriptorsByLabel(label string) ([]contract.ValueDescriptor, error) {
	// TODO implement queries on `[]string` in the core
	if objects, err := client.ValueDescriptors(); err != nil {
		return nil, err
	} else {
		// manually search all value descriptors for the given label
		var result = make([]contract.ValueDescriptor, 0)
		for _, object := range objects {
			for _, str := range object.Labels {
				if label == str {
					result = append(result, object)
					break
				}
			}
		}
		return result, nil
	}
}

func (client *ObjectBoxClient) ValueDescriptorsByType(t string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.typ

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Type, t); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *ObjectBoxClient) ScrubAllValueDescriptors() error {
	return client.valueDescriptorBox.RemoveAll()
}

func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}

//endregion
