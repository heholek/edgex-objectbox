package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	. "github.com/objectbox/objectbox-go/objectbox"
)

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox

	strictReads bool
	asyncPut    bool

	eventBox           *obx.EventBox
	readingBox         *obx.ReadingBox
	valueDescriptorBox *obx.ValueDescriptorBox

	queries coreDataQueries
}

func NewClient(config db.Configuration) (*ObjectBoxClient, error) {
	println(VersionInfo())
	client := &ObjectBoxClient{config: config}
	return client, client.Connect()
}

// Considers client.strictReads
func (client *ObjectBoxClient) storeForReads() *ObjectBox {
	store := client.objectBox
	if client.strictReads {
		store.AwaitAsyncCompletion()
	}
	return store
}

// Considers client.strictReads
// TODO this should be moved to the core/objectbox-go - it's also necessary for remove, update, queries
func (client *ObjectBoxClient) eventBoxForReads() *obx.EventBox {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.eventBox
}

// Considers client.strictReads
func (client *ObjectBoxClient) readingBoxForReads() *obx.ReadingBox {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.readingBox
}

// Considers client.strictReads
func (client *ObjectBoxClient) valueDescriptorBoxForReads() *obx.ValueDescriptorBox {
	if client.strictReads {
		client.objectBox.AwaitAsyncCompletion()
	}
	return client.valueDescriptorBox
}

func (client *ObjectBoxClient) Connect() error {
	objectBox, err := NewBuilder().Directory(client.config.DatabaseName).Model(obx.ObjectBoxModel()).Build()
	if err != nil {
		return err
	}
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)

	client.objectBox = objectBox
	client.eventBox = obx.BoxForEvent(objectBox)
	client.readingBox = obx.BoxForReading(objectBox)
	client.valueDescriptorBox = obx.BoxForValueDescriptor(objectBox)

	// don't use asyncPut by default, unique constraint violation fails silently
	// TODO consider removing altogether or moving to the core
	//client.asyncPut = true
	//client.strictReads = true

	return client.initCoreData()
}

func (client *ObjectBoxClient) Disconnect() {
	client.eventBox = nil
	client.readingBox = nil
	client.valueDescriptorBox = nil

	objectBoxToDestroy := client.objectBox
	client.objectBox = nil
	if objectBoxToDestroy != nil {
		objectBoxToDestroy.Close()
	}
}

func (client *ObjectBoxClient) CloseSession() {
	client.Disconnect()
}
