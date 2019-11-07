package objectbox

// implements core-command service contract

// WARNING: In Edinburgh, this was all part of core-metadata. In Fuji, EdgeX decided to implement new methods
// (GetCommandsByDeviceId, GetCommandByNameAndDeviceId) in a new microservice, keeping the rest in the original one.
// This breaks service isolation and prevents using ObjectBox (or any embedded database).
//
// In order to circumvent this, we update the core-command service to forward calls to the core-metadata service and
// keep the commands DB as part of the metadata DB
// In the future development for Geneva release, this should be resolved in upstream EdgeX.

import (
	contract "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/defs"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/obx"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type coreCommandClient struct {
	objectBox *objectbox.ObjectBox

	commandBox *obx.CommandBox // no async - a config

	queries coreCommandQueries
}

//region Queries
type coreCommandQueries struct {
	command struct {
		name          commandQuery
		device        commandQuery
		nameAndDevice commandQuery
	}
}

type commandQuery struct {
	*obx.CommandQuery
	sync.Mutex
}

//endregion

func newCoreCommandClient(objectBox *objectbox.ObjectBox) (*coreCommandClient, error) {
	var client = &coreCommandClient{objectBox: objectBox}
	var err error

	client.commandBox = obx.BoxForCommand(objectBox)

	//region Command
	if err == nil {
		client.queries.command.name.CommandQuery, err =
			client.commandBox.QueryOrError(obx.Command_.Name.Equals("", true))
	}
	if err == nil {
		client.queries.command.device.CommandQuery, err =
			client.commandBox.QueryOrError(obx.Command_.DeviceId.Equals(0))
	}
	if err == nil {
		client.queries.command.nameAndDevice.CommandQuery, err =
			client.commandBox.QueryOrError(obx.Command_.Name.Equals("", true), obx.Command_.DeviceId.Equals(0))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, mapError(err)
	}
}

// Fuji: used in core-metadata
func (client *coreCommandClient) GetAllCommands() ([]contract.Command, error) {
	commands, err := client.commandBox.GetAll()
	return defs.CommandSliceToContract(commands), mapError(err)
}

// Fuji: used in core-metadata
func (client *coreCommandClient) GetCommandById(id string) (contract.Command, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Command{}, mapError(err)
	} else if object, err := client.commandBox.Get(id); err != nil {
		return contract.Command{}, mapError(err)
	} else if object == nil {
		return contract.Command{}, mapError(db.ErrNotFound)
	} else {
		return object.ToContract(), nil
	}
}

// Fuji: used in core-metadata
func (client *coreCommandClient) GetCommandsByName(name string) ([]contract.Command, error) {
	var query = &client.queries.command.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Command_.Name, name); err != nil {
		return nil, mapError(err)
	}

	commands, err := query.Limit(0).Find()
	return defs.CommandSliceToContract(commands), mapError(err)
}

// Fuji: used in core-metadata and core-command
func (client *coreCommandClient) GetCommandsByDeviceId(idString string) ([]contract.Command, error) {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return nil, mapError(err)
	}

	var query = &client.queries.command.device

	query.Lock()
	defer query.Unlock()

	err = query.SetInt64Params(obx.Command_.DeviceId, int64(id))
	if err != nil {
		return nil, mapError(err)
	}

	commands, err := query.Limit(0).Find()
	return defs.CommandSliceToContract(commands), mapError(err)
}

// Fuji: used in core-command
func (client *coreCommandClient) GetCommandByNameAndDeviceId(cname string, did string) (contract.Command, error) {
	id, err := obx.IdFromString(did)
	if err != nil {
		return contract.Command{}, mapError(err)
	}

	var query = &client.queries.command.nameAndDevice

	query.Lock()
	defer query.Unlock()

	err = query.SetStringParams(obx.Command_.Name, cname)
	if err != nil {
		return contract.Command{}, mapError(err)
	}

	err = query.SetInt64Params(obx.Command_.DeviceId, int64(id))
	if err != nil {
		return contract.Command{}, mapError(err)
	}

	commands, err := query.Limit(1).Find()
	if err != nil {
		return contract.Command{}, mapError(err)
	}

	if len(commands) < 1 {
		return contract.Command{}, db.ErrNotFound
	}

	return commands[0].ToContract(), nil
}

func (client *coreCommandClient) addCommand(c contract.Command, deviceId uint64) (string, error) {
	onCreate(&c.Timestamps)
	var cmd = defs.Command{
		Timestamps: c.Timestamps,
		Id:         c.Id,
		Name:       c.Name,
		Get: defs.Get{Action: defs.Action{
			Path:      c.Get.Action.Path,
			Responses: c.Get.Responses,
			URL:       c.Get.URL,
		}},
		Put: defs.Put{
			Action: defs.Action{
				Path:      c.Put.Action.Path,
				Responses: c.Put.Action.Responses,
				URL:       c.Put.Action.URL,
			},
			ParameterNames: c.Put.ParameterNames,
		},

		DeviceId: deviceId,
	}
	id, err := client.commandBox.Put(&cmd)
	return obx.IdToString(id), mapError(err)
}

func (client *coreCommandClient) deleteCommandsByDeviceId(idString string) error {
	id, err := obx.IdFromString(idString)
	if err != nil {
		return mapError(err)
	}

	var query = &client.queries.command.device

	query.Lock()
	defer query.Unlock()

	err = query.SetInt64Params(obx.Command_.DeviceId, int64(id))
	if err != nil {
		return mapError(err)
	}

	_, err = query.Limit(0).Remove()
	return mapError(err)
}

//func (client *coreCommandClient) UpdateCommand(c contract.Command) error {
//	onUpdate(&c.Timestamps)
//
//	if id, err := obx.IdFromString(c.Id); err != nil {
//		return mapError(err)
//	} else if exists, err := client.commandBox.Contains(id); err != nil {
//		return mapError(err)
//	} else if !exists {
//		return mapError(db.ErrNotFound)
//	}
//
//	_, err := client.commandBox.Put(&c)
//	return mapError(err)
//}
