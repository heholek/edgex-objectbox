// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	"errors"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type notification_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var NotificationBinding = notification_EntityInfo{
	Entity: objectbox.Entity{
		Id: 11,
	},
	Uid: 3035506812026122772,
}

// Notification_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Notification_ = struct {
	Created     *objectbox.PropertyInt64
	Modified    *objectbox.PropertyInt64
	Origin      *objectbox.PropertyInt64
	ID          *objectbox.PropertyUint64
	Slug        *objectbox.PropertyString
	Sender      *objectbox.PropertyString
	Category    *objectbox.PropertyString
	Severity    *objectbox.PropertyString
	Content     *objectbox.PropertyString
	Description *objectbox.PropertyString
	Status      *objectbox.PropertyString
	Labels      *objectbox.PropertyStringVector
	ContentType *objectbox.PropertyString
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &NotificationBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &NotificationBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &NotificationBinding.Entity,
		},
	},
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &NotificationBinding.Entity,
		},
	},
	Slug: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &NotificationBinding.Entity,
		},
	},
	Sender: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &NotificationBinding.Entity,
		},
	},
	Category: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &NotificationBinding.Entity,
		},
	},
	Severity: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &NotificationBinding.Entity,
		},
	},
	Content: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &NotificationBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &NotificationBinding.Entity,
		},
	},
	Status: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &NotificationBinding.Entity,
		},
	},
	Labels: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &NotificationBinding.Entity,
		},
	},
	ContentType: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     13,
			Entity: &NotificationBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (notification_EntityInfo) GeneratorVersion() int {
	return 5
}

// AddToModel is called by ObjectBox during model build
func (notification_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Notification", 11, 3035506812026122772)
	model.Property("Created", 6, 1, 4223274968704546949)
	model.Property("Modified", 6, 2, 6140168204663194960)
	model.Property("Origin", 6, 3, 5119840969146189721)
	model.Property("ID", 6, 4, 1898755377907498256)
	model.PropertyFlags(1)
	model.Property("Slug", 9, 5, 5614391508684015285)
	model.PropertyFlags(32)
	model.PropertyIndex(11, 8355138627950004661)
	model.Property("Sender", 9, 6, 1698316443180001548)
	model.Property("Category", 9, 7, 6130508587898235784)
	model.Property("Severity", 9, 8, 7998250136380607007)
	model.Property("Content", 9, 9, 3265040071503511803)
	model.Property("Description", 9, 10, 3025723844696420299)
	model.Property("Status", 9, 11, 8672829707123022794)
	model.Property("Labels", 30, 12, 6661777360307054348)
	model.Property("ContentType", 9, 13, 3705832325243429372)
	model.EntityLastPropertyId(13, 3705832325243429372)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (notification_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Notification); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID)
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Notification).ID)
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (notification_EntityInfo) SetId(object interface{}, id uint64) error {
	if obj, ok := object.(*Notification); ok {
		var err error
		obj.ID, err = objectbox.StringIdConvertToEntityProperty(id)
		return err
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Notification).ID
		return nil
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (notification_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (notification_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *Notification
	if objPtr, ok := object.(*Notification); ok {
		obj = objPtr
	} else {
		objVal := object.(Notification)
		obj = &objVal
	}

	var offsetSlug = fbutils.CreateStringOffset(fbb, obj.Slug)
	var offsetSender = fbutils.CreateStringOffset(fbb, obj.Sender)
	var offsetCategory = fbutils.CreateStringOffset(fbb, string(obj.Category))
	var offsetSeverity = fbutils.CreateStringOffset(fbb, string(obj.Severity))
	var offsetContent = fbutils.CreateStringOffset(fbb, obj.Content)
	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.Description)
	var offsetStatus = fbutils.CreateStringOffset(fbb, string(obj.Status))
	var offsetLabels = fbutils.CreateStringVectorOffset(fbb, obj.Labels)
	var offsetContentType = fbutils.CreateStringOffset(fbb, obj.ContentType)

	// build the FlatBuffers object
	fbb.StartObject(13)
	fbutils.SetInt64Slot(fbb, 0, obj.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Timestamps.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetSlug)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetSender)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetCategory)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetSeverity)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetContent)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetDescription)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetStatus)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetLabels)
	fbutils.SetUOffsetTSlot(fbb, 12, offsetContentType)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (notification_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'Notification' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	propID, err := objectbox.StringIdConvertToEntityProperty(fbutils.GetUint64Slot(table, 10))
	if err != nil {
		return nil, errors.New("converter objectbox.StringIdConvertToEntityProperty() failed on Notification.ID: " + err.Error())
	}

	return &Notification{
		Timestamps: models.Timestamps{
			Created:  fbutils.GetInt64Slot(table, 4),
			Modified: fbutils.GetInt64Slot(table, 6),
			Origin:   fbutils.GetInt64Slot(table, 8),
		},
		ID:          propID,
		Slug:        fbutils.GetStringSlot(table, 12),
		Sender:      fbutils.GetStringSlot(table, 14),
		Category:    models.NotificationsCategory(fbutils.GetStringSlot(table, 16)),
		Severity:    models.NotificationsSeverity(fbutils.GetStringSlot(table, 18)),
		Content:     fbutils.GetStringSlot(table, 20),
		Description: fbutils.GetStringSlot(table, 22),
		Status:      models.NotificationsStatus(fbutils.GetStringSlot(table, 24)),
		Labels:      fbutils.GetStringVectorSlot(table, 26),
		ContentType: fbutils.GetStringSlot(table, 28),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (notification_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Notification, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (notification_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]Notification), Notification{})
	}
	return append(slice.([]Notification), *object.(*Notification))
}

// Box provides CRUD access to Notification objects
type NotificationBox struct {
	*objectbox.Box
}

// BoxForNotification opens a box of Notification objects
func BoxForNotification(ob *objectbox.ObjectBox) *NotificationBox {
	return &NotificationBox{
		Box: ob.InternalBox(11),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Notification.ID property on the passed object will be assigned the new ID as well.
func (box *NotificationBox) Put(object *Notification) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Notification.ID property on the passed object will be assigned the new ID as well.
func (box *NotificationBox) Insert(object *Notification) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *NotificationBox) Update(object *Notification) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *NotificationBox) PutAsync(object *Notification) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Notification.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Notification.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *NotificationBox) PutMany(objects []Notification) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *NotificationBox) Get(id uint64) (*Notification, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Notification), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *NotificationBox) GetMany(ids ...uint64) ([]Notification, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *NotificationBox) GetManyExisting(ids ...uint64) ([]Notification, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// GetAll reads all stored objects
func (box *NotificationBox) GetAll() ([]Notification, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// Remove deletes a single object
func (box *NotificationBox) Remove(object *Notification) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *NotificationBox) RemoveMany(objects ...*Notification) (uint64, error) {
	var ids = make([]uint64, len(objects))
	var err error
	for k, object := range objects {
		ids[k], err = objectbox.StringIdConvertToDatabaseValue(object.ID)
		if err != nil {
			return 0, errors.New("converter objectbox.StringIdConvertToDatabaseValue() failed on Notification.ID: " + err.Error())
		}
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Notification_ struct to create conditions.
// Keep the *NotificationQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *NotificationBox) Query(conditions ...objectbox.Condition) *NotificationQuery {
	return &NotificationQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Notification_ struct to create conditions.
// Keep the *NotificationQuery if you intend to execute the query multiple times.
func (box *NotificationBox) QueryOrError(conditions ...objectbox.Condition) (*NotificationQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &NotificationQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See NotificationAsyncBox for more information.
func (box *NotificationBox) Async() *NotificationAsyncBox {
	return &NotificationAsyncBox{AsyncBox: box.Box.Async()}
}

// NotificationAsyncBox provides asynchronous operations on Notification objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type NotificationAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForNotification creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use NotificationBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForNotification(ob *objectbox.ObjectBox, timeoutMs uint64) *NotificationAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 11, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 11: %s" + err.Error())
	}
	return &NotificationAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the ID property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *NotificationAsyncBox) Put(object *Notification) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The ID property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *NotificationAsyncBox) Insert(object *Notification) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *NotificationAsyncBox) Update(object *Notification) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *NotificationAsyncBox) Remove(object *Notification) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all Notification which ID is either 42 or 47:
// 		box.Query(Notification_.ID.In(42, 47)).Find()
type NotificationQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *NotificationQuery) Find() ([]Notification, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Notification), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *NotificationQuery) Offset(offset uint64) *NotificationQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *NotificationQuery) Limit(limit uint64) *NotificationQuery {
	query.Query.Limit(limit)
	return query
}
