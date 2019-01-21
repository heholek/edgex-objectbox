package objectbox

// implements core-metadata service contract
// TODO indexes

import (
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
)

type coreMetaDataClient struct {
	objectBox *objectbox.ObjectBox

	queries coreMetaDataQueries
}

//region Queries
type coreMetaDataQueries struct {
}

//endregion

func newCoreMetaDataClient(objectBox *objectbox.ObjectBox) (*coreMetaDataClient, error) {
	var client = &coreMetaDataClient{objectBox: objectBox}
	var err error

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
func (client *coreMetaDataClient) UpdateDevice(d contract.Device) error   { panic(notImplemented()) }
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
	panic(notImplemented())
}
func (client *coreMetaDataClient) AddAddressable(a contract.Addressable) (string, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressableById(id string) (contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressableByName(n string) (contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressablesByTopic(t string) ([]contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressablesByPort(p int) ([]contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressablesByPublisher(p string) ([]contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressablesByAddress(add string) ([]contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAddressables() ([]contract.Addressable, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) DeleteAddressableById(id string) error { panic(notImplemented()) }
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
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetCommandByName(id string) ([]contract.Command, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) AddCommand(c contract.Command) (string, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) GetAllCommands() ([]contract.Command, error) {
	panic(notImplemented())
}
func (client *coreMetaDataClient) UpdateCommand(c contract.Command) error { panic(notImplemented()) }
func (client *coreMetaDataClient) DeleteCommandById(id string) error      { panic(notImplemented()) }
func (client *coreMetaDataClient) ScrubMetadata() error                   { panic(notImplemented()) }
