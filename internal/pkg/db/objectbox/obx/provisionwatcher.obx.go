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

type provisionWatcher_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ProvisionWatcherBinding = provisionWatcher_EntityInfo{
	Entity: objectbox.Entity{
		Id: 12,
	},
	Uid: 678769668479103726,
}

// ProvisionWatcher_ contains type-based Property helpers to facilitate some common operations such as Queries.
var ProvisionWatcher_ = struct {
	Created        *objectbox.PropertyInt64
	Modified       *objectbox.PropertyInt64
	Origin         *objectbox.PropertyInt64
	Id             *objectbox.PropertyUint64
	Name           *objectbox.PropertyString
	Identifiers    *objectbox.PropertyByteVector
	Profile        *objectbox.RelationToOne
	Service        *objectbox.RelationToOne
	AdminState     *objectbox.PropertyString
	OperatingState *objectbox.PropertyString
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Identifiers: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	Profile: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     7,
			Entity: &ProvisionWatcherBinding.Entity,
		},
		Target: &DeviceProfileBinding.Entity,
	},
	Service: &objectbox.RelationToOne{
		Property: &objectbox.BaseProperty{
			Id:     8,
			Entity: &ProvisionWatcherBinding.Entity,
		},
		Target: &DeviceServiceBinding.Entity,
	},
	AdminState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
	OperatingState: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &ProvisionWatcherBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (provisionWatcher_EntityInfo) GeneratorVersion() int {
	return 5
}

// AddToModel is called by ObjectBox during model build
func (provisionWatcher_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("ProvisionWatcher", 12, 678769668479103726)
	model.Property("Created", 6, 1, 5400688834240787055)
	model.Property("Modified", 6, 2, 6936168735950790495)
	model.Property("Origin", 6, 3, 7324362782559730584)
	model.Property("Id", 6, 4, 4214810142464261408)
	model.PropertyFlags(1)
	model.Property("Name", 9, 5, 1922646737032650250)
	model.PropertyFlags(32)
	model.PropertyIndex(12, 8347802915601169150)
	model.Property("Identifiers", 23, 6, 6659681934816812940)
	model.Property("Profile", 11, 7, 4488324877991092491)
	model.PropertyFlags(8712)
	model.PropertyRelation("DeviceProfile", 13, 5329856807707561884)
	model.Property("Service", 11, 8, 6873740026061739530)
	model.PropertyFlags(8712)
	model.PropertyRelation("DeviceService", 14, 3453358122163741587)
	model.Property("AdminState", 9, 10, 92900689631506377)
	model.Property("OperatingState", 9, 9, 3437982289020393516)
	model.EntityLastPropertyId(10, 92900689631506377)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (provisionWatcher_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*ProvisionWatcher); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id)
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(ProvisionWatcher).Id)
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (provisionWatcher_EntityInfo) SetId(object interface{}, id uint64) error {
	if obj, ok := object.(*ProvisionWatcher); ok {
		var err error
		obj.Id, err = objectbox.StringIdConvertToEntityProperty(id)
		return err
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(ProvisionWatcher).Id
		return nil
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (provisionWatcher_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if rel := &object.(*ProvisionWatcher).Profile; rel != nil {
		if rId, err := DeviceProfileBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForDeviceProfile(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	if rel := &object.(*ProvisionWatcher).Service; rel != nil {
		if rId, err := DeviceServiceBinding.GetId(rel); err != nil {
			return err
		} else if rId == 0 {
			// NOTE Put/PutAsync() has a side-effect of setting the rel.ID
			if _, err := BoxForDeviceService(ob).Put(rel); err != nil {
				return err
			}
		}
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (provisionWatcher_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *ProvisionWatcher
	if objPtr, ok := object.(*ProvisionWatcher); ok {
		obj = objPtr
	} else {
		objVal := object.(ProvisionWatcher)
		obj = &objVal
	}

	var propIdentifiers []byte
	{
		var err error
		propIdentifiers, err = mapStringStringJsonToDatabaseValue(obj.Identifiers)
		if err != nil {
			return errors.New("converter mapStringStringJsonToDatabaseValue() failed on ProvisionWatcher.Identifiers: " + err.Error())
		}
	}

	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetIdentifiers = fbutils.CreateByteVectorOffset(fbb, propIdentifiers)
	var offsetAdminState = fbutils.CreateStringOffset(fbb, string(obj.AdminState))
	var offsetOperatingState = fbutils.CreateStringOffset(fbb, string(obj.OperatingState))

	var rIdProfile uint64
	if rel := &obj.Profile; rel != nil {
		if rId, err := DeviceProfileBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdProfile = rId
		}
	}

	var rIdService uint64
	if rel := &obj.Service; rel != nil {
		if rId, err := DeviceServiceBinding.GetId(rel); err != nil {
			return err
		} else {
			rIdService = rId
		}
	}

	// build the FlatBuffers object
	fbb.StartObject(10)
	fbutils.SetInt64Slot(fbb, 0, obj.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Timestamps.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetIdentifiers)
	fbutils.SetUint64Slot(fbb, 6, rIdProfile)
	fbutils.SetUint64Slot(fbb, 7, rIdService)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetAdminState)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetOperatingState)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (provisionWatcher_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	if len(bytes) == 0 { // sanity check, should "never" happen
		return nil, errors.New("can't deserialize an object of type 'ProvisionWatcher' - no data received")
	}

	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	propId, err := objectbox.StringIdConvertToEntityProperty(fbutils.GetUint64Slot(table, 10))
	if err != nil {
		return nil, errors.New("converter objectbox.StringIdConvertToEntityProperty() failed on ProvisionWatcher.Id: " + err.Error())
	}

	propIdentifiers, err := mapStringStringJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 14))
	if err != nil {
		return nil, errors.New("converter mapStringStringJsonToEntityProperty() failed on ProvisionWatcher.Identifiers: " + err.Error())
	}

	var relProfile *DeviceProfile
	if rId := fbutils.GetUint64Slot(table, 16); rId > 0 {
		if rObject, err := BoxForDeviceProfile(ob).Get(rId); err != nil {
			return nil, err
		} else if rObject == nil {
			relProfile = &DeviceProfile{}
		} else {
			relProfile = rObject
		}
	} else {
		relProfile = &DeviceProfile{}
	}

	var relService *DeviceService
	if rId := fbutils.GetUint64Slot(table, 18); rId > 0 {
		if rObject, err := BoxForDeviceService(ob).Get(rId); err != nil {
			return nil, err
		} else if rObject == nil {
			relService = &DeviceService{}
		} else {
			relService = rObject
		}
	} else {
		relService = &DeviceService{}
	}

	return &ProvisionWatcher{
		Timestamps: models.Timestamps{
			Created:  fbutils.GetInt64Slot(table, 4),
			Modified: fbutils.GetInt64Slot(table, 6),
			Origin:   fbutils.GetInt64Slot(table, 8),
		},
		Id:             propId,
		Name:           fbutils.GetStringSlot(table, 12),
		Identifiers:    propIdentifiers,
		Profile:        *relProfile,
		Service:        *relService,
		AdminState:     models.AdminState(fbutils.GetStringSlot(table, 22)),
		OperatingState: models.OperatingState(fbutils.GetStringSlot(table, 20)),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (provisionWatcher_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]ProvisionWatcher, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (provisionWatcher_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if object == nil {
		return append(slice.([]ProvisionWatcher), ProvisionWatcher{})
	}
	return append(slice.([]ProvisionWatcher), *object.(*ProvisionWatcher))
}

// Box provides CRUD access to ProvisionWatcher objects
type ProvisionWatcherBox struct {
	*objectbox.Box
}

// BoxForProvisionWatcher opens a box of ProvisionWatcher objects
func BoxForProvisionWatcher(ob *objectbox.ObjectBox) *ProvisionWatcherBox {
	return &ProvisionWatcherBox{
		Box: ob.InternalBox(12),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ProvisionWatcher.Id property on the passed object will be assigned the new ID as well.
func (box *ProvisionWatcherBox) Put(object *ProvisionWatcher) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the ProvisionWatcher.Id property on the passed object will be assigned the new ID as well.
func (box *ProvisionWatcherBox) Insert(object *ProvisionWatcher) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *ProvisionWatcherBox) Update(object *ProvisionWatcher) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *ProvisionWatcherBox) PutAsync(object *ProvisionWatcher) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the ProvisionWatcher.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the ProvisionWatcher.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ProvisionWatcherBox) PutMany(objects []ProvisionWatcher) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ProvisionWatcherBox) Get(id uint64) (*ProvisionWatcher, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*ProvisionWatcher), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *ProvisionWatcherBox) GetMany(ids ...uint64) ([]ProvisionWatcher, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// GetManyExisting reads multiple objects at once, skipping those that do not exist.
func (box *ProvisionWatcherBox) GetManyExisting(ids ...uint64) ([]ProvisionWatcher, error) {
	objects, err := box.Box.GetManyExisting(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// GetAll reads all stored objects
func (box *ProvisionWatcherBox) GetAll() ([]ProvisionWatcher, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// Remove deletes a single object
func (box *ProvisionWatcherBox) Remove(object *ProvisionWatcher) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *ProvisionWatcherBox) RemoveMany(objects ...*ProvisionWatcher) (uint64, error) {
	var ids = make([]uint64, len(objects))
	var err error
	for k, object := range objects {
		ids[k], err = objectbox.StringIdConvertToDatabaseValue(object.Id)
		if err != nil {
			return 0, errors.New("converter objectbox.StringIdConvertToDatabaseValue() failed on ProvisionWatcher.Id: " + err.Error())
		}
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the ProvisionWatcher_ struct to create conditions.
// Keep the *ProvisionWatcherQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ProvisionWatcherBox) Query(conditions ...objectbox.Condition) *ProvisionWatcherQuery {
	return &ProvisionWatcherQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the ProvisionWatcher_ struct to create conditions.
// Keep the *ProvisionWatcherQuery if you intend to execute the query multiple times.
func (box *ProvisionWatcherBox) QueryOrError(conditions ...objectbox.Condition) (*ProvisionWatcherQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ProvisionWatcherQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See ProvisionWatcherAsyncBox for more information.
func (box *ProvisionWatcherBox) Async() *ProvisionWatcherAsyncBox {
	return &ProvisionWatcherAsyncBox{AsyncBox: box.Box.Async()}
}

// ProvisionWatcherAsyncBox provides asynchronous operations on ProvisionWatcher objects.
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
type ProvisionWatcherAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForProvisionWatcher creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use ProvisionWatcherBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForProvisionWatcher(ob *objectbox.ObjectBox, timeoutMs uint64) *ProvisionWatcherAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 12, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 12: %s" + err.Error())
	}
	return &ProvisionWatcherAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *ProvisionWatcherAsyncBox) Put(object *ProvisionWatcher) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *ProvisionWatcherAsyncBox) Insert(object *ProvisionWatcher) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *ProvisionWatcherAsyncBox) Update(object *ProvisionWatcher) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *ProvisionWatcherAsyncBox) Remove(object *ProvisionWatcher) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all ProvisionWatcher which Id is either 42 or 47:
// 		box.Query(ProvisionWatcher_.Id.In(42, 47)).Find()
type ProvisionWatcherQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ProvisionWatcherQuery) Find() ([]ProvisionWatcher, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]ProvisionWatcher), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ProvisionWatcherQuery) Offset(offset uint64) *ProvisionWatcherQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ProvisionWatcherQuery) Limit(limit uint64) *ProvisionWatcherQuery {
	query.Query.Limit(limit)
	return query
}
