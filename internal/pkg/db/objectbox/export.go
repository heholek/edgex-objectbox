package objectbox

// implements export-client service contract
// TODO indexes

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type exportClient struct {
	objectBox *objectbox.ObjectBox

	registrationBox *obx.RegistrationBox

	queries exportQueries
}

//region Queries
type exportQueries struct {
	registration struct {
		name registrationQuery
	}
}

type registrationQuery struct {
	*obx.RegistrationQuery
	sync.Mutex
}

//endregion

func newExportClient(objectBox *objectbox.ObjectBox) (*exportClient, error) {
	var client = &exportClient{objectBox: objectBox}
	var err error

	client.registrationBox = obx.BoxForRegistration(objectBox)

	//region Registration
	if err == nil {
		client.queries.registration.name.RegistrationQuery, err =
			client.registrationBox.QueryOrError(obx.Registration_.Name.Equals("", true))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *exportClient) Registrations() ([]contract.Registration, error) {
	return client.registrationBox.GetAll()
}

func (client *exportClient) AddRegistration(reg contract.Registration) (string, error) {
	// NOTE this is done instead of onCreate because there is no reg.BaseObject
	if reg.Created == 0 {
		reg.Created = db.MakeTimestamp()
	}

	id, err := client.registrationBox.Put(&reg)
	return obx.IdToString(id), err
}

func (client *exportClient) UpdateRegistration(reg contract.Registration) error {
	// NOTE this is done instead of onUpdate because there is no reg.BaseObject
	reg.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(reg.ID); err != nil {
		return err
	} else if exists, err := client.registrationBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.registrationBox.Put(&reg)
	return err
}

func (client *exportClient) RegistrationById(id string) (contract.Registration, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Registration{}, err
	} else if object, err := client.registrationBox.Get(id); err != nil {
		return contract.Registration{}, err
	} else if object == nil {
		return contract.Registration{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *exportClient) DeleteRegistrationById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.registrationBox.Box.Remove(id)
	}
}

func (client *exportClient) RegistrationByName(name string) (contract.Registration, error) {
	var query = &client.queries.registration.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Registration_.Name, name); err != nil {
		return contract.Registration{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Registration{}, err
	} else if len(list) == 0 {
		return contract.Registration{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *exportClient) DeleteRegistrationByName(name string) error {
	if obj, err := client.RegistrationByName(name); err != nil {
		return err
	} else {
		return client.registrationBox.Remove(&obj)
	}
}

func (client *exportClient) ScrubAllRegistrations() error {
	return client.registrationBox.RemoveAll()
}
