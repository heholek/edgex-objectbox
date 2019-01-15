package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	. "github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type ObjectBoxClient struct {
	config    db.Configuration
	objectBox *ObjectBox

	eventBox   *obx.EventBox
	readingBox *obx.ReadingBox

	queryEventByDeviceId      *obx.EventQuery
	queryEventByDeviceIdMutex sync.Mutex

	queryReadingByDeviceId      *obx.ReadingQuery
	queryReadingByDeviceIdMutex sync.Mutex

	strictReads bool
	asyncPut    bool
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

func (client *ObjectBoxClient) Connect() error {
	objectBox, err := NewBuilder().Directory(client.config.DatabaseName).Model(obx.ObjectBoxModel()).Build()
	if err != nil {
		return err
	}
	//objectBox.SetDebugFlags(DebugFlags_LOG_ASYNC_QUEUE)

	client.objectBox = objectBox
	client.eventBox = obx.BoxForEvent(objectBox)
	client.readingBox = obx.BoxForReading(objectBox)
	client.asyncPut = true
	client.strictReads = true

	client.queryEventByDeviceId, err = client.eventBox.QueryOrError(obx.Event_.Device.Equals("", true))
	if err != nil {
		return err
	}

	client.queryReadingByDeviceId, err = client.readingBox.QueryOrError(obx.Reading_.Device.Equals("", true))
	if err != nil {
		return err
	}

	return err
}

func (client *ObjectBoxClient) Disconnect() {
	client.eventBox = nil
	client.readingBox = nil
	objectBoxToDestroy := client.objectBox
	client.objectBox = nil
	if objectBoxToDestroy != nil {
		objectBoxToDestroy.Close()
	}
}
