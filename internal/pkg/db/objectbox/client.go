package objectbox

import (
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db/objectbox/obx"
	. "github.com/objectbox/objectbox-go/objectbox"
)

// NOTE: REST API tests (blackbox) sometime import data with relation IDs missing (e.g. Transmission.Notification.ID).
//  This causes ObjectBox to try & insert the related entity but then fails if there's a unique constraint. To improve
//  compatibility, we try to find related entities before inserting them. e.g. AddTransmission, AddRegistration

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox

	// embedded services
	*coreDataClient
	*coreMetaDataClient
	*coreCommandClient
	*exportClient
	*schedulerClient
	*notificationsClient
}

// use Async where possible
const async = true

func NewClient(config db.Configuration) (*ObjectBoxClient, error) {
	println(VersionInfo())
	client := &ObjectBoxClient{config: config}
	return client, client.Connect()
}

func (client *ObjectBoxClient) Connect() error {
	objectBox, err := NewBuilder().
		Directory(client.config.DatabaseName).
		Model(obx.ObjectBoxModel()).
		Build()

	if err != nil {
		return err
	}

	client.objectBox = objectBox

	if err == nil {
		client.coreDataClient, err = newCoreDataClient(objectBox)
	}

	if err == nil {
		client.coreCommandClient, err = newCoreCommandClient(objectBox)
	}

	if err == nil {
		client.coreMetaDataClient, err = newCoreMetaDataClient(objectBox, client.coreCommandClient)
	}

	if err == nil {
		client.exportClient, err = newExportClient(objectBox)
	}

	if err == nil {
		client.schedulerClient, err = newSchedulerClient(objectBox)
	}

	if err == nil {
		client.notificationsClient, err = newNotificationsClient(objectBox)
	}

	if err != nil {
		client.CloseSession()
	}
	return err
}

func (client *ObjectBoxClient) CloseSession() {
	objectBoxToDestroy := client.objectBox
	client.objectBox = nil

	// remove service references
	client.coreDataClient = nil
	client.coreMetaDataClient = nil
	client.coreCommandClient = nil
	client.exportClient = nil
	client.schedulerClient = nil
	client.notificationsClient = nil

	if objectBoxToDestroy != nil {
		objectBoxToDestroy.AwaitAsyncCompletion()
		objectBoxToDestroy.Close()
	}
}

// NOTE this is not in the upstream
func (client *ObjectBoxClient) EnsureAllDurable(async bool) error {
	client.objectBox.AwaitAsyncCompletion()
	return nil
}
