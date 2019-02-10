package objectbox

// implements core-data service contract

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type coreDataClient struct {
	objectBox *objectbox.ObjectBox

	eventBox           *obx.EventBox
	readingBox         *obx.ReadingBox
	valueDescriptorBox *obx.ValueDescriptorBox

	queries coreDataQueries
}

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
		labels   valueDescriptorQuery
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

func newCoreDataClient(objectBox *objectbox.ObjectBox) (*coreDataClient, error) {
	var client = &coreDataClient{objectBox: objectBox}
	var err error

	client.eventBox = obx.BoxForEvent(client.objectBox)
	client.readingBox = obx.BoxForReading(client.objectBox)
	client.valueDescriptorBox = obx.BoxForValueDescriptor(client.objectBox)

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
		client.queries.valueDescriptor.labels.ValueDescriptorQuery, err =
			client.valueDescriptorBox.QueryOrError(obx.ValueDescriptor_.Labels.Contains("", true))
	}

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

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *coreDataClient) Events() ([]contract.Event, error) {
	return client.eventBox.GetAll()
}

func (client *coreDataClient) EventsWithLimit(limit int) ([]contract.Event, error) {
	// TODO there is no test for this method in the test/db_data.go
	var query = &client.queries.event.all

	query.Lock()
	defer query.Unlock()

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) AddEvent(event contract.Event) (string, error) {
	if event.Created == 0 {
		event.Created = db.MakeTimestamp()
	}

	// TODO currently tests don't add any readings to the event

	var id uint64
	var err error

	if asyncPut {
		id, err = client.eventBox.PutAsync(&event)
	} else {
		id, err = client.eventBox.Put(&event)
	}

	return obx.IdToString(id), err
}

func (client *coreDataClient) UpdateEvent(e contract.Event) error {
	e.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(e.ID); err != nil {
		return err
	} else if exists, err := client.eventBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	var err error
	if asyncPut {
		_, err = client.eventBox.PutAsync(&e)
	} else {
		_, err = client.eventBox.Put(&e)
	}

	return err
}

func (client *coreDataClient) EventById(id string) (contract.Event, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Event{}, err
	} else if object, err := client.eventBox.Get(id); err != nil {
		return contract.Event{}, err
	} else if object == nil {
		return contract.Event{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *coreDataClient) EventCount() (count int, err error) {
	countLong, err := client.eventBox.Count()
	if err == nil {
		count = int(countLong)
	}
	return
}

func (client *coreDataClient) EventCountByDeviceId(id string) (int, error) {
	var query = &client.queries.event.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return 0, err
	}

	count, err := query.Count()
	return int(count), err
}

func (client *coreDataClient) DeleteEventById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.eventBox.Box.Remove(id)
}

func (client *coreDataClient) EventsForDeviceLimit(id string, limit int) ([]contract.Event, error) {
	var query = &client.queries.event.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) EventsForDevice(id string) ([]contract.Event, error) {
	return client.EventsForDeviceLimit(id, 0)
}

func (client *coreDataClient) EventsByCreationTime(start, end int64, limit int) ([]contract.Event, error) {
	var query = &client.queries.event.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) EventsOlderThanAge(age int64) ([]contract.Event, error) {
	var time = db.MakeTimestamp() - age

	var query = &client.queries.event.createdLT

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, time); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreDataClient) EventsPushed() ([]contract.Event, error) {
	var query = &client.queries.event.pushedGT

	query.Lock()
	defer query.Unlock()

	return query.Find()
}

func (client *coreDataClient) ScrubAllEvents() error {
	if err := client.eventBox.RemoveAll(); err != nil {
		return err
	}
	return client.readingBox.RemoveAll()
}

func (client *coreDataClient) Readings() ([]contract.Reading, error) {
	return client.readingBox.GetAll()
}

func (client *coreDataClient) AddReading(r contract.Reading) (string, error) {
	if r.Created == 0 {
		r.Created = db.MakeTimestamp()
	}

	var id uint64
	var err error

	if asyncPut {
		id, err = client.readingBox.PutAsync(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}

	return obx.IdToString(id), err
}

func (client *coreDataClient) UpdateReading(r contract.Reading) error {
	r.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(r.Id); err != nil {
		return err
	} else if exists, err := client.readingBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	var err error
	if asyncPut {
		_, err = client.readingBox.PutAsync(&r)
	} else {
		_, err = client.readingBox.Put(&r)
	}

	return err
}

func (client *coreDataClient) ReadingById(id string) (contract.Reading, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Reading{}, err
	} else if object, err := client.readingBox.Get(id); err != nil {
		return contract.Reading{}, err
	} else if object == nil {
		return contract.Reading{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *coreDataClient) ReadingCount() (int, error) {
	count, err := client.readingBox.Count()
	return int(count), err
}

func (client *coreDataClient) DeleteReadingById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.readingBox.Box.Remove(id)
}

func (client *coreDataClient) ReadingsByDevice(deviceId string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) ReadingsByValueDescriptor(name string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Name, name); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.Reading_.Name, names...); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) ReadingsByCreationTime(start, end int64, limit int) ([]contract.Reading, error) {
	var query = &client.queries.reading.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Reading_.Created, start, end); err != nil {
		return nil, err
	}

	return query.Limit(uint64(limit)).Find()
}

func (client *coreDataClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]contract.Reading, error) {
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

func (client *coreDataClient) AddValueDescriptor(v contract.ValueDescriptor) (string, error) {
	if v.Created == 0 {
		v.Created = db.MakeTimestamp()
	}

	// TODO tests don't set Max, Min, Default (interface{})

	id, err := client.valueDescriptorBox.Put(&v)
	return obx.IdToString(id), err
}

func (client *coreDataClient) ValueDescriptors() ([]contract.ValueDescriptor, error) {
	return client.valueDescriptorBox.GetAll()
}

func (client *coreDataClient) UpdateValueDescriptor(v contract.ValueDescriptor) error {
	v.Modified = db.MakeTimestamp()

	// check whether it exists, otherwise this function must fail
	if object, err := client.valueDescriptorById(v.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	_, err := client.valueDescriptorBox.Put(&v)
	return err
}

func (client *coreDataClient) DeleteValueDescriptorById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.valueDescriptorBox.Box.Remove(id)
}

func (client *coreDataClient) ValueDescriptorByName(name string) (contract.ValueDescriptor, error) {
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

func (client *coreDataClient) ValueDescriptorsByName(names []string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.ValueDescriptor_.Name, names...); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreDataClient) ValueDescriptorById(id string) (contract.ValueDescriptor, error) {
	object, err := client.valueDescriptorById(id)
	if object == nil || err != nil {
		return contract.ValueDescriptor{}, err
	}
	return *object, nil
}

func (client *coreDataClient) valueDescriptorById(idString string) (*contract.ValueDescriptor, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.valueDescriptorBox.Get(id)
}

func (client *coreDataClient) ValueDescriptorsByUomLabel(uomLabel string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.uomlabel

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.UomLabel, uomLabel); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreDataClient) ValueDescriptorsByLabel(label string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Labels, label); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreDataClient) ValueDescriptorsByType(t string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.typ

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Type, t); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreDataClient) ScrubAllValueDescriptors() error {
	return client.valueDescriptorBox.RemoveAll()
}
