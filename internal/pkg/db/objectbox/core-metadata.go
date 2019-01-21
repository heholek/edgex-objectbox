package objectbox

// implements core-metadata service contract
// TODO indexes

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type coreMetaDataClient struct {
	objectBox *objectbox.ObjectBox

	commandBox     *obx.CommandBox
	addressableBox *obx.AddressableBox

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
}

type addressableQuery struct {
	*obx.AddressableQuery
	sync.Mutex
}

type commandQuery struct {
	*obx.CommandQuery
	sync.Mutex
}

//endregion

func newCoreMetaDataClient(objectBox *objectbox.ObjectBox) (*coreMetaDataClient, error) {
	var client = &coreMetaDataClient{objectBox: objectBox}
	var err error

	client.commandBox = obx.BoxForCommand(objectBox)
	client.addressableBox = obx.BoxForAddressable(objectBox)

	//region Command
	if err == nil {
		client.queries.command.name.CommandQuery, err =
			client.commandBox.QueryOrError(obx.Command_.Name.Equals("", true))
	}
	//endregion

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

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *coreMetaDataClient) GetAllScheduleEvents(se *[]contract.ScheduleEvent) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddScheduleEvent(se *contract.ScheduleEvent) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetScheduleEventByName(se *contract.ScheduleEvent, n string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) UpdateScheduleEvent(se contract.ScheduleEvent) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetScheduleEventById(se *contract.ScheduleEvent, id string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetScheduleEventsByScheduleName(se *[]contract.ScheduleEvent, n string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetScheduleEventsByAddressableId(se *[]contract.ScheduleEvent, id string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetScheduleEventsByServiceName(se *[]contract.ScheduleEvent, n string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteScheduleEventById(id string) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetAllSchedules(s *[]contract.Schedule) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddSchedule(s *contract.Schedule) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetScheduleByName(s *contract.Schedule, n string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) UpdateSchedule(s contract.Schedule) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetScheduleById(s *contract.Schedule, id string) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteScheduleById(id string) error { panic(notImplemented()) }

func (client *coreMetaDataClient) GetAllDeviceReports() ([]contract.DeviceReport, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceReportByDeviceName(n string) ([]contract.DeviceReport, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceReportByName(n string) (contract.DeviceReport, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceReportById(id string) (contract.DeviceReport, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddDeviceReport(dr contract.DeviceReport) (string, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) UpdateDeviceReport(dr contract.DeviceReport) error {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceReportsByScheduleEventName(n string) ([]contract.DeviceReport, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteDeviceReportById(id string) error { panic(notImplemented()) }

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
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceServicesByAddressableId(id string) ([]contract.DeviceService, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceServicesWithLabel(l string) ([]contract.DeviceService, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceServiceById(id string) (contract.DeviceService, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetDeviceServiceByName(n string) (contract.DeviceService, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) GetAllDeviceServices() ([]contract.DeviceService, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) AddDeviceService(ds contract.DeviceService) (string, error) {
	panic(notImplemented())
}

func (client *coreMetaDataClient) DeleteDeviceServiceById(id string) error { panic(notImplemented()) }

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
