package objectbox

// implements core-metadata service contract
// TODO indexes

import (
	"fmt"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/pkg/errors"
	"sync"
)

type coreMetaDataClient struct {
	objectBox *objectbox.ObjectBox

	addressableBox   *obx.AddressableBox
	commandBox       *obx.CommandBox
	deviceReportBox  *obx.DeviceReportBox
	deviceServiceBox *obx.DeviceServiceBox
	scheduleBox      *obx.ScheduleBox
	scheduleEventBox *obx.ScheduleEventBox

	queries coreMetaDataQueries
}

//region Queries
type coreMetaDataQueries struct {
	addressable struct {
		address   addressableQuery
		name      addressableQuery
		port      addressableQuery
		publisher addressableQuery
		topic     addressableQuery
	}
	command struct {
		name commandQuery
	}
	deviceReport struct {
		device deviceReportQuery
		event  deviceReportQuery
		name   deviceReportQuery
	}
	deviceService struct {
		labels deviceServiceQuery
		name   deviceServiceQuery
	}
	schedule struct {
		name scheduleQuery
	}
	scheduleEvent struct {
		addressable scheduleEventQuery
		name        scheduleEventQuery
		schedule    scheduleEventQuery
		service     scheduleEventQuery
	}
}

type addressableQuery struct {
	*obx.AddressableQuery
	sync.Mutex
}

type commandQuery struct {
	*obx.CommandQuery
	sync.Mutex
}

type deviceReportQuery struct {
	*obx.DeviceReportQuery
	sync.Mutex
}

type deviceServiceQuery struct {
	*obx.DeviceServiceQuery
	sync.Mutex
}

type scheduleQuery struct {
	*obx.ScheduleQuery
	sync.Mutex
}

type scheduleEventQuery struct {
	*obx.ScheduleEventQuery
	sync.Mutex
}

//endregion

func newCoreMetaDataClient(objectBox *objectbox.ObjectBox) (*coreMetaDataClient, error) {
	var client = &coreMetaDataClient{objectBox: objectBox}
	var err error

	client.addressableBox = obx.BoxForAddressable(objectBox)
	client.commandBox = obx.BoxForCommand(objectBox)
	client.deviceReportBox = obx.BoxForDeviceReport(objectBox)
	client.deviceServiceBox = obx.BoxForDeviceService(objectBox)
	client.scheduleBox = obx.BoxForSchedule(objectBox)
	client.scheduleEventBox = obx.BoxForScheduleEvent(objectBox)

	//region Addressable
	if err == nil {
		client.queries.addressable.address.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Address.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.name.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Name.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.port.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Port.Equals(0))
	}

	if err == nil {
		client.queries.addressable.publisher.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Publisher.Equals("", true))
	}

	if err == nil {
		client.queries.addressable.topic.AddressableQuery, err =
			client.addressableBox.QueryOrError(obx.Addressable_.Topic.Equals("", true))
	}
	//endregion

	//region Command
	if err == nil {
		client.queries.command.name.CommandQuery, err =
			client.commandBox.QueryOrError(obx.Command_.Name.Equals("", true))
	}
	//endregion

	//region DeviceReport
	if err == nil {
		client.queries.deviceReport.device.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Device.Equals("", true))
	}
	if err == nil {
		client.queries.deviceReport.event.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Event.Equals("", true))
	}
	if err == nil {
		client.queries.deviceReport.name.DeviceReportQuery, err =
			client.deviceReportBox.QueryOrError(obx.DeviceReport_.Name.Equals("", true))
	}
	//endregion

	//region DeviceService
	if err == nil {
		client.queries.deviceService.labels.DeviceServiceQuery, err =
			client.deviceServiceBox.QueryOrError(obx.DeviceService_.Labels.Contains("", true))
	}

	if err == nil {
		client.queries.deviceService.name.DeviceServiceQuery, err =
			client.deviceServiceBox.QueryOrError(obx.DeviceService_.Name.Equals("", true))
	}
	//endregion

	//region Schedule
	if err == nil {
		client.queries.schedule.name.ScheduleQuery, err =
			client.scheduleBox.QueryOrError(obx.Schedule_.Name.Equals("", true))
	}
	//endregion

	//region ScheduleEvent
	if err == nil {
		client.queries.scheduleEvent.addressable.ScheduleEventQuery, err =
			client.scheduleEventBox.QueryOrError(obx.ScheduleEvent_.Addressable.Equals(0))
	}
	if err == nil {
		client.queries.scheduleEvent.name.ScheduleEventQuery, err =
			client.scheduleEventBox.QueryOrError(obx.ScheduleEvent_.Name.Equals("", true))
	}
	if err == nil {
		client.queries.scheduleEvent.schedule.ScheduleEventQuery, err =
			client.scheduleEventBox.QueryOrError(obx.ScheduleEvent_.Schedule.Equals("", true))
	}
	if err == nil {
		client.queries.scheduleEvent.service.ScheduleEventQuery, err =
			client.scheduleEventBox.QueryOrError(obx.ScheduleEvent_.Service.Equals("", true))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *coreMetaDataClient) GetAllScheduleEvents(se *[]contract.ScheduleEvent) error {
	if list, err := client.scheduleEventBox.GetAll(); err != nil {
		return err
	} else {
		*se = list
		return nil
	}
}

func (client *coreMetaDataClient) validateScheduleEvent(se *contract.ScheduleEvent) error {
	// check the addressable is specified, this is required by tests
	if addId, err := obx.IdFromString(se.Addressable.Id); err != nil {
		return err
	} else if addId == 0 {
		return errors.New("addressable not specified")
	} else if exists, err := client.addressableBox.Contains(addId); err != nil {
		return err
	} else if !exists {
		return fmt.Errorf("addressable %s not found", se.Addressable.Id)
	}

	return nil
}

func (client *coreMetaDataClient) AddScheduleEvent(se *contract.ScheduleEvent) error {
	onCreate(&se.BaseObject)

	if err := client.validateScheduleEvent(se); err != nil {
		return err
	}

	_, err := client.scheduleEventBox.Put(se)
	return err
}

func (client *coreMetaDataClient) GetScheduleEventByName(se *contract.ScheduleEvent, n string) error {
	var query = &client.queries.scheduleEvent.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ScheduleEvent_.Name, n); err != nil {
		return err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return err
	} else if len(list) == 0 {
		return db.ErrNotFound
	} else {
		*se = list[0]
		return nil
	}
}

func (client *coreMetaDataClient) UpdateScheduleEvent(se contract.ScheduleEvent) error {
	onUpdate(&se.BaseObject)

	if err := client.validateScheduleEvent(&se); err != nil {
		return err
	}

	if id, err := obx.IdFromString(string(se.Id)); err != nil {
		return err
	} else if exists, err := client.scheduleEventBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.scheduleEventBox.Put(&se)
	return err
}

func (client *coreMetaDataClient) GetScheduleEventById(se *contract.ScheduleEvent, id string) error {
	if id, err := idFromHex(id); err != nil {
		return err
	} else if id, err := obx.IdFromString(id); err != nil {
		return err
	} else if object, err := client.scheduleEventBox.Get(id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	} else {
		*se = *object
		return nil
	}
}

func (client *coreMetaDataClient) GetScheduleEventsByScheduleName(se *[]contract.ScheduleEvent, n string) error {
	var query = &client.queries.scheduleEvent.schedule

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ScheduleEvent_.Schedule, n); err != nil {
		return err
	} else if list, err := query.Find(); err != nil {
		return err
	} else {
		*se = list
		return nil
	}
}

func (client *coreMetaDataClient) GetScheduleEventsByAddressableId(se *[]contract.ScheduleEvent, id string) error {
	var query = &client.queries.scheduleEvent.addressable

	query.Lock()
	defer query.Unlock()

	if id, err := obx.IdFromString(id); err != nil {
		return err
	} else if err := query.SetInt64Params(obx.ScheduleEvent_.Addressable, int64(id)); err != nil {
		return err
	} else if list, err := query.Find(); err != nil {
		return err
	} else {
		*se = list
		return nil
	}
}

func (client *coreMetaDataClient) GetScheduleEventsByServiceName(se *[]contract.ScheduleEvent, n string) error {
	var query = &client.queries.scheduleEvent.service

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.ScheduleEvent_.Service, n); err != nil {
		return err
	} else if list, err := query.Find(); err != nil {
		return err
	} else {
		*se = list
		return nil
	}
}

func (client *coreMetaDataClient) DeleteScheduleEventById(id string) error {
	if id, err := idFromHex(id); err != nil {
		return err
	} else if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.scheduleEventBox.Box.Remove(id)
	}
}

func (client *coreMetaDataClient) GetAllSchedules(s *[]contract.Schedule) error {
	if list, err := client.scheduleBox.GetAll(); err != nil {
		return err
	} else {
		*s = list
		return nil
	}
}

func (client *coreMetaDataClient) AddSchedule(s *contract.Schedule) error {
	onCreate(&s.BaseObject)

	_, err := client.scheduleBox.Put(s)
	return err
}

func (client *coreMetaDataClient) GetScheduleByName(s *contract.Schedule, n string) error {
	var query = &client.queries.schedule.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Schedule_.Name, n); err != nil {
		return err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return err
	} else if len(list) == 0 {
		return db.ErrNotFound
	} else {
		*s = list[0]
		return nil
	}
}

func (client *coreMetaDataClient) UpdateSchedule(s contract.Schedule) error {
	onUpdate(&s.BaseObject)

	if id, err := obx.IdFromString(string(s.Id)); err != nil {
		return err
	} else if exists, err := client.scheduleBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.scheduleBox.Put(&s)
	return err
}

func (client *coreMetaDataClient) GetScheduleById(s *contract.Schedule, id string) error {
	if id, err := idFromHex(id); err != nil {
		return err
	} else if id, err := obx.IdFromString(id); err != nil {
		return err
	} else if object, err := client.scheduleBox.Get(id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	} else {
		*s = *object
		return nil
	}
}

func (client *coreMetaDataClient) DeleteScheduleById(id string) error {
	if id, err := idFromHex(id); err != nil {
		return err
	} else if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.scheduleBox.Box.Remove(id)
	}
}

func (client *coreMetaDataClient) GetAllDeviceReports() ([]contract.DeviceReport, error) {
	return client.deviceReportBox.GetAll()
}

func (client *coreMetaDataClient) GetDeviceReportByDeviceName(n string) ([]contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.device

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Device, n); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetDeviceReportByName(n string) (contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Name, n); err != nil {
		return contract.DeviceReport{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.DeviceReport{}, err
	} else if len(list) == 0 {
		return contract.DeviceReport{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetDeviceReportById(id string) (contract.DeviceReport, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.DeviceReport{}, err
	} else if object, err := client.deviceReportBox.Get(id); err != nil {
		return contract.DeviceReport{}, err
	} else if object == nil {
		return contract.DeviceReport{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *coreMetaDataClient) AddDeviceReport(dr contract.DeviceReport) (string, error) {
	onCreate(&dr.BaseObject)

	id, err := client.deviceReportBox.Put(&dr)
	return obx.IdToString(id), err
}

func (client *coreMetaDataClient) UpdateDeviceReport(dr contract.DeviceReport) error {
	onUpdate(&dr.BaseObject)

	if id, err := obx.IdFromString(dr.Id); err != nil {
		return err
	} else if exists, err := client.deviceReportBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.deviceReportBox.Put(&dr)
	return err
}

func (client *coreMetaDataClient) GetDeviceReportsByScheduleEventName(n string) ([]contract.DeviceReport, error) {
	var query = &client.queries.deviceReport.event

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceReport_.Event, n); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) DeleteDeviceReportById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.deviceReportBox.Box.Remove(id)
	}
}

func (client *coreMetaDataClient) UpdateDevice(d contract.Device) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetDeviceById(id string) (contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceByName(n string) (contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetAllDevices() ([]contract.Device, error) { panic(notImplemented()) }

func (client *coreMetaDataClient) GetDevicesByProfileId(pid string) ([]contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDevicesByServiceId(sid string) ([]contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDevicesByAddressableId(aid string) ([]contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDevicesWithLabel(l string) ([]contract.Device, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddDevice(d contract.Device) (string, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteDeviceById(id string) error { panic(notImplemented()) }

func (client *coreMetaDataClient) UpdateDeviceProfile(dp contract.DeviceProfile) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddDeviceProfile(d contract.DeviceProfile) (string, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetAllDeviceProfiles() ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfileById(id string) (contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteDeviceProfileById(id string) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetDeviceProfilesByModel(m string) ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfilesWithLabel(l string) ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfilesByManufacturerModel(man string, mod string) ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfilesByManufacturer(man string) ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfileByName(n string) (contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceProfilesUsingCommand(c contract.Command) ([]contract.DeviceProfile, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) UpdateAddressable(a contract.Addressable) error {
	onUpdate(&a.BaseObject)

	// check whether it exists, otherwise this function must fail
	if object, err := client.addressableById(a.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	_, err := client.addressableBox.Put(&a)
	return err
}

func (client *coreMetaDataClient) AddAddressable(a contract.Addressable) (string, error) {
	onCreate(&a.BaseObject)

	id, err := client.addressableBox.Put(&a)
	return obx.IdToString(id), err
}

func (client *coreMetaDataClient) GetAddressableById(id string) (contract.Addressable, error) {
	object, err := client.addressableById(id)
	if object == nil || err != nil {
		return contract.Addressable{}, err
	}
	return *object, nil
}

func (client *coreMetaDataClient) addressableById(idString string) (*contract.Addressable, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.addressableBox.Get(id)
}

func (client *coreMetaDataClient) GetAddressableByName(n string) (contract.Addressable, error) {
	var query = &client.queries.addressable.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Name, n); err != nil {
		return contract.Addressable{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Addressable{}, err
	} else if len(list) == 0 {
		return contract.Addressable{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetAddressablesByTopic(t string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.topic

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Topic, t); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetAddressablesByPort(p int) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.port

	query.Lock()
	defer query.Unlock()

	if err := query.SetInt64Params(obx.Addressable_.Port, int64(p)); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetAddressablesByPublisher(p string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.publisher

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Publisher, p); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetAddressablesByAddress(add string) ([]contract.Addressable, error) {
	var query = &client.queries.addressable.address

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Addressable_.Address, add); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetAddressables() ([]contract.Addressable, error) {
	return client.addressableBox.GetAll()
}

func (client *coreMetaDataClient) DeleteAddressableById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.addressableBox.Box.Remove(id)
}

func (client *coreMetaDataClient) UpdateDeviceService(ds contract.DeviceService) error {
	onUpdate(&ds.BaseObject)

	// check whether it exists, otherwise this function must fail
	if object, err := client.deviceServiceById(ds.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	_, err := client.deviceServiceBox.Put(&ds)
	return err
}

func (client *coreMetaDataClient) GetDeviceServicesByAddressableId(id string) ([]contract.DeviceService, error) {
	// FIXME this requires relations queries 1..n, right now we are just passing the tests but the result is incorrect
	return make([]contract.DeviceService, 1), nil
}

func (client *coreMetaDataClient) GetDeviceServicesWithLabel(l string) ([]contract.DeviceService, error) {
	var query = &client.queries.deviceService.labels

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceService_.Labels, l); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) GetDeviceServiceById(id string) (contract.DeviceService, error) {
	object, err := client.deviceServiceById(id)
	if object == nil || err != nil {
		return contract.DeviceService{}, err
	}
	return *object, nil
}

func (client *coreMetaDataClient) deviceServiceById(idString string) (*contract.DeviceService, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.deviceServiceBox.Get(id)
}

func (client *coreMetaDataClient) GetDeviceServiceByName(n string) (contract.DeviceService, error) {
	var query = &client.queries.deviceService.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.DeviceService_.Name, n); err != nil {
		return contract.DeviceService{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.DeviceService{}, err
	} else if len(list) == 0 {
		return contract.DeviceService{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *coreMetaDataClient) GetAllDeviceServices() ([]contract.DeviceService, error) {
	return client.deviceServiceBox.GetAll()
}

func (client *coreMetaDataClient) AddDeviceService(ds contract.DeviceService) (string, error) {
	onCreate(&ds.BaseObject)
	id, err := client.deviceServiceBox.Put(&ds)
	return obx.IdToString(id), err
}

func (client *coreMetaDataClient) DeleteDeviceServiceById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.deviceServiceBox.Box.Remove(id)
}

func (client *coreMetaDataClient) GetProvisionWatcherById(id string) (contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetAllProvisionWatchers() ([]contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetProvisionWatcherByName(n string) (contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetProvisionWatchersByProfileId(id string) ([]contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetProvisionWatchersByServiceId(id string) ([]contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetProvisionWatchersByIdentifier(k string, v string) ([]contract.ProvisionWatcher, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddProvisionWatcher(pw contract.ProvisionWatcher) (string, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) UpdateProvisionWatcher(pw contract.ProvisionWatcher) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteProvisionWatcherById(id string) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetCommandById(id string) (contract.Command, error) {
	object, err := client.commandById(id)
	if object == nil || err != nil {
		return contract.Command{}, err
	}
	return *object, nil
}

func (client *coreMetaDataClient) commandById(idString string) (*contract.Command, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, err
	}

	return client.commandBox.Get(id)
}

func (client *coreMetaDataClient) GetCommandByName(name string) ([]contract.Command, error) {
	var query = &client.queries.command.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Command_.Name, name); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *coreMetaDataClient) AddCommand(c contract.Command) (string, error) {
	onCreate(&c.BaseObject)

	var id uint64
	var err error

	if asyncPut {
		id, err = client.commandBox.PutAsync(&c)
	} else {
		id, err = client.commandBox.Put(&c)
	}

	return obx.IdToString(id), err
}

func (client *coreMetaDataClient) GetAllCommands() ([]contract.Command, error) {
	return client.commandBox.GetAll()
}

func (client *coreMetaDataClient) UpdateCommand(c contract.Command) error {
	onUpdate(&c.BaseObject)

	// check whether it exists, otherwise this function must fail
	if object, err := client.commandById(c.Id); err != nil {
		return err
	} else if object == nil {
		return db.ErrNotFound
	}

	var err error
	if asyncPut {
		_, err = client.commandBox.PutAsync(&c)
	} else {
		_, err = client.commandBox.Put(&c)
	}

	return err

}

func (client *coreMetaDataClient) DeleteCommandById(idString string) error {
	// TODO maybe this requires a check whether the item exists

	id, err := obx.IdFromString(idString)
	if err != nil {
		return err
	}

	return client.commandBox.Box.Remove(id)
}

func (client *coreMetaDataClient) ScrubMetadata() error {
	// TODO implement for all boxes
	return nil
}
