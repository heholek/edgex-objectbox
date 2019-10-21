package objectbox

// implements core-data service contract

import (
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	correlation "github.com/objectbox/edgex-objectbox/internal/pkg/correlation/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/obx"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type coreDataClient struct {
	objectBox *objectbox.ObjectBox

	eventBox           *obx.EventBox           // no async - has relation
	readingBox         *obx.ReadingBox         // async used
	valueDescriptorBox *obx.ValueDescriptorBox // no async - a config

	readingAsync *AsyncView

	queries coreDataQueries
}

//region Queries
type coreDataQueries struct {
	event struct {
		all       eventQuery
		checksum  eventQuery
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

	client.readingAsync = newAsyncView(client.readingBox.Box)

	//region Event
	if err == nil {
		client.queries.event.all.EventQuery, err = client.eventBox.QueryOrError()
	}

	if err == nil {
		client.queries.event.device.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Device.Equals("", true))
	}

	if err == nil {
		client.queries.event.checksum.EventQuery, err =
			client.eventBox.QueryOrError(obx.Event_.Checksum.Equals("", true))
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
		return nil, mapError(err)
	}
}

func (client *coreDataClient) awaitAsync() {
	if async {
		client.readingAsync.async.AwaitSubmitted()
		client.readingAsync.Clear()
	}
}

func correlationEventsToContractEvents(events []correlation.Event) []contract.Event {
	var result []contract.Event
	for k := range events {
		result = append(result, events[k].Event)
	}
	return result
}

func (client *coreDataClient) Events() ([]contract.Event, error) {
	result, err := client.eventBox.GetAll()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) EventsWithLimit(limit int) ([]contract.Event, error) {
	// TODO there is no test for this method in the test/db_data.go
	var query = &client.queries.event.all

	query.Lock()
	defer query.Unlock()

	result, err := query.Limit(uint64(limit)).Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) AddEvent(e correlation.Event) (string, error) {
	client.awaitAsync()

	if e.Created == 0 {
		e.Created = db.MakeTimestamp()
	}

	for i := range e.Readings {
		var reading = &e.Readings[i]
		if reading.Created == 0 {
			reading.Created = db.MakeTimestamp()
		}

		// same thing as Mongo binding does
		if reading.Device == "" {
			reading.Device = e.Device
		}
	}

	id, err := client.eventBox.Put(&e)
	return obx.IdToString(id), mapError(err)
}

func (client *coreDataClient) UpdateEvent(e correlation.Event) error {
	client.awaitAsync()

	// as we don't do lazy-loading externally, if the slice is nil, it's empty, not a "not-yet-loaded" lazy one
	if e.Event.Readings == nil {
		e.Event.Readings = []contract.Reading{}
	}

	e.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(e.ID); err != nil {
		return mapError(err)
	} else if exists, err := client.eventBox.Contains(id); err != nil {
		return mapError(err)
	} else if !exists {
		return mapError(db.ErrNotFound)
	}

	_, err := client.eventBox.Put(&e)
	return mapError(err)
}

func (client *coreDataClient) EventById(id string) (contract.Event, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Event{}, mapError(err)
	} else if object, err := client.eventBox.Get(id); err != nil {
		return contract.Event{}, mapError(err)
	} else if object == nil {
		return contract.Event{}, mapError(db.ErrNotFound)
	} else {
		return object.Event, nil
	}
}

func (client *coreDataClient) EventsByChecksum(checksum string) ([]contract.Event, error) {
	var query = &client.queries.event.checksum

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Checksum, checksum); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) EventCount() (int, error) {
	countLong, err := client.eventBox.Count()
	return int(countLong), mapError(err)
}

func (client *coreDataClient) EventCountByDeviceId(id string) (int, error) {
	var query = &client.queries.event.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return 0, mapError(err)
	}

	count, err := query.Limit(0).Count()
	return int(count), mapError(err)
}

func (client *coreDataClient) DeleteEventById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	// TODO: Do this async to speed up deleteEventsByAge() when AsyncBox is available in objectbox-go
	return mapError(client.eventBox.RemoveId(id))
}

func (client *coreDataClient) EventsForDeviceLimit(id string, limit int) ([]contract.Event, error) {
	var query = &client.queries.event.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Event_.Device, id); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) EventsForDevice(id string) ([]contract.Event, error) {
	result, err := client.EventsForDeviceLimit(id, 0)
	return result, mapError(err)
}

func (client *coreDataClient) EventsByCreationTime(start, end int64, limit int) ([]contract.Event, error) {
	var query = &client.queries.event.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, start, end); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) EventsOlderThanAge(age int64) ([]contract.Event, error) {
	var time = db.MakeTimestamp() - age

	var query = &client.queries.event.createdLT

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Event_.Created, time); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) EventsPushed() ([]contract.Event, error) {
	var query = &client.queries.event.pushedGT

	query.Lock()
	defer query.Unlock()

	result, err := query.Limit(0).Find()
	return correlationEventsToContractEvents(result), mapError(err)
}

func (client *coreDataClient) ScrubAllEvents() error {
	client.awaitAsync()

	return client.objectBox.RunInWriteTx(func() error {
		if err := client.eventBox.RemoveAll(); err != nil {
			return mapError(err)
		}
		return mapError(client.readingBox.RemoveAll())
	})
}

func (client *coreDataClient) Readings() ([]contract.Reading, error) {
	client.awaitAsync()

	result, err := client.readingBox.GetAll()
	return result, mapError(err)
}

func (client *coreDataClient) AddReading(r contract.Reading) (string, error) {
	if r.Created == 0 {
		r.Created = db.MakeTimestamp()
	}

	var id uint64
	var err error

	if async {
		id, err = client.readingAsync.Insert(&r)
	} else {
		id, err = client.readingBox.Put(&r)
	}

	return obx.IdToString(id), mapError(err)
}

func (client *coreDataClient) UpdateReading(r contract.Reading) error {
	r.Modified = db.MakeTimestamp()

	id, err := obx.IdFromString(r.Id)
	if err != nil {
		return mapError(err)
	}

	if async {
		err = client.readingAsync.Update(id, &r)
	} else {
		_, err = client.readingBox.Put(&r)
	}

	return mapError(err)
}

func (client *coreDataClient) ReadingById(id string) (contract.Reading, error) {
	client.awaitAsync()

	if id, err := obx.IdFromString(id); err != nil {
		return contract.Reading{}, mapError(err)
	} else if object, err := client.readingBox.Get(id); err != nil {
		return contract.Reading{}, mapError(err)
	} else if object == nil {
		return contract.Reading{}, mapError(db.ErrNotFound)
	} else {
		return *object, nil
	}
}

func (client *coreDataClient) ReadingCount() (int, error) {
	client.awaitAsync()

	count, err := client.readingBox.Count()
	return int(count), mapError(err)
}

func (client *coreDataClient) DeleteReadingById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	if async {
		return mapError(client.readingAsync.RemoveId(id))
	}

	return mapError(client.readingBox.RemoveId(id))
}

func (client *coreDataClient) ReadingsByDevice(deviceId string, limit int) ([]contract.Reading, error) {
	client.awaitAsync()

	var query = &client.queries.reading.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ReadingsByValueDescriptor(name string, limit int) ([]contract.Reading, error) {
	client.awaitAsync()

	var query = &client.queries.reading.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Name, name); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ReadingsByValueDescriptorNames(names []string, limit int) ([]contract.Reading, error) {
	client.awaitAsync()

	var query = &client.queries.reading.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.Reading_.Name, names...); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ReadingsByCreationTime(start, end int64, limit int) ([]contract.Reading, error) {
	client.awaitAsync()

	var query = &client.queries.reading.createdB

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Reading_.Created, start, end); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ReadingsByDeviceAndValueDescriptor(deviceId, valueDescriptor string, limit int) ([]contract.Reading, error) {
	client.awaitAsync()

	var query = &client.queries.reading.deviceAndName

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Reading_.Device, deviceId); err != nil {
		return nil, mapError(err)
	}
	if err := query.SetStringParams(obx.Reading_.Name, valueDescriptor); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(uint64(limit)).Find()
	return result, mapError(err)
}

func (client *coreDataClient) AddValueDescriptor(v contract.ValueDescriptor) (string, error) {
	if v.Created == 0 {
		v.Created = db.MakeTimestamp()
	}

	// TODO tests don't set Max, Min, Default (interface{})

	id, err := client.valueDescriptorBox.Put(&v)
	return obx.IdToString(id), mapError(err)
}

func (client *coreDataClient) ValueDescriptors() ([]contract.ValueDescriptor, error) {
	result, err := client.valueDescriptorBox.GetAll()
	return result, mapError(err)
}

func (client *coreDataClient) UpdateValueDescriptor(v contract.ValueDescriptor) error {
	v.Modified = db.MakeTimestamp()

	// check whether it exists, otherwise this function must fail
	if object, err := client.valueDescriptorById(v.Id); err != nil {
		return mapError(err)
	} else if object == nil {
		return mapError(db.ErrNotFound)
	}

	_, err := client.valueDescriptorBox.Put(&v)
	return mapError(err)
}

func (client *coreDataClient) DeleteValueDescriptorById(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	return mapError(client.valueDescriptorBox.RemoveId(id))
}

func (client *coreDataClient) ValueDescriptorByName(name string) (contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Name, name); err != nil {
		return contract.ValueDescriptor{}, mapError(err)
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.ValueDescriptor{}, mapError(err)
	} else if len(list) == 0 {
		return contract.ValueDescriptor{}, mapError(db.ErrNotFound)
	} else {
		return list[0], nil
	}
}

func (client *coreDataClient) ValueDescriptorsByName(names []string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.names

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParamsIn(obx.ValueDescriptor_.Name, names...); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ValueDescriptorById(id string) (contract.ValueDescriptor, error) {
	object, err := client.valueDescriptorById(id)
	if object == nil || err != nil {
		return contract.ValueDescriptor{}, mapError(err)
	}
	return *object, nil
}

func (client *coreDataClient) valueDescriptorById(idString string) (*contract.ValueDescriptor, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, mapError(err)
	}

	result, err := client.valueDescriptorBox.Get(id)
	return result, mapError(err)
}

func (client *coreDataClient) ValueDescriptorsByUomLabel(uomLabel string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.uomlabel

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.UomLabel, uomLabel); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ValueDescriptorsByLabel(label string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Labels, label); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ValueDescriptorsByType(t string) ([]contract.ValueDescriptor, error) {
	var query = &client.queries.valueDescriptor.typ

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ValueDescriptor_.Type, t); err != nil {
		return nil, mapError(err)
	}

	result, err := query.Limit(0).Find()
	return result, mapError(err)
}

func (client *coreDataClient) ScrubAllValueDescriptors() error {
	return client.valueDescriptorBox.RemoveAll()
}
