// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type deviceProfile_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var DeviceProfileBinding = deviceProfile_EntityInfo{
	Entity: objectbox.Entity{
		Id: 5,
	},
	Uid: 2211424933411195666,
}

// DeviceProfile_ contains type-based Property helpers to facilitate some common operations such as Queries.
var DeviceProfile_ = struct {
	Timestamps_Created  *objectbox.PropertyInt64
	Timestamps_Modified *objectbox.PropertyInt64
	Timestamps_Origin   *objectbox.PropertyInt64
	Description         *objectbox.PropertyString
	Id                  *objectbox.PropertyUint64
	Name                *objectbox.PropertyString
	Manufacturer        *objectbox.PropertyString
	Model               *objectbox.PropertyString
	Labels              *objectbox.PropertyStringVector
	DeviceResources     *objectbox.PropertyByteVector
	DeviceCommands      *objectbox.PropertyByteVector
	CoreCommands        *objectbox.RelationToMany
}{
	Timestamps_Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Timestamps_Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Timestamps_Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Description: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Manufacturer: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Model: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	Labels: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	DeviceResources: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	DeviceCommands: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &DeviceProfileBinding.Entity,
		},
	},
	CoreCommands: &objectbox.RelationToMany{
		Id:     1,
		Source: &DeviceProfileBinding.Entity,
		Target: &CommandBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (deviceProfile_EntityInfo) GeneratorVersion() int {
	return 2
}

// AddToModel is called by ObjectBox during model build
func (deviceProfile_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("DeviceProfile", 5, 2211424933411195666)
	model.Property("Timestamps_Created", objectbox.PropertyType_Long, 1, 2642282108714980568)
	model.Property("Timestamps_Modified", objectbox.PropertyType_Long, 2, 2107950979126715489)
	model.Property("Timestamps_Origin", objectbox.PropertyType_Long, 3, 238686394209975473)
	model.Property("Description", objectbox.PropertyType_String, 4, 7373170799262197703)
	model.Property("Id", objectbox.PropertyType_Long, 5, 2058046092051357134)
	model.PropertyFlags(objectbox.PropertyFlags_ID | objectbox.PropertyFlags_UNSIGNED)
	model.Property("Name", objectbox.PropertyType_String, 6, 3395397931582751949)
	model.PropertyFlags(objectbox.PropertyFlags_UNIQUE)
	model.PropertyIndex(4, 2508596274702817237)
	model.Property("Manufacturer", objectbox.PropertyType_String, 7, 7853971810229114061)
	model.Property("Model", objectbox.PropertyType_String, 8, 6511893101827755872)
	model.Property("Labels", objectbox.PropertyType_StringVector, 9, 3915749655245396678)
	model.Property("DeviceResources", objectbox.PropertyType_ByteVector, 10, 349654535178347918)
	model.Property("DeviceCommands", objectbox.PropertyType_ByteVector, 11, 5969260086641149857)
	model.EntityLastPropertyId(11, 5969260086641149857)
	model.Relation(1, 2308887105935734223, CommandBinding.Id, CommandBinding.Uid)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (deviceProfile_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*DeviceProfile); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(DeviceProfile).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (deviceProfile_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*DeviceProfile); ok {
		obj.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(DeviceProfile).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (deviceProfile_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	if err := txn.RunWithCursor(DeviceProfileBinding.Id, func(cursor *objectbox.Cursor) error {
		return cursor.RelationReplace(1, CommandBinding.Id, id, object, object.(*DeviceProfile).CoreCommands)
	}); err != nil {
		return err
	}
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (deviceProfile_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *DeviceProfile
	if objPtr, ok := object.(*DeviceProfile); ok {
		obj = objPtr
	} else {
		objVal := object.(DeviceProfile)
		obj = &objVal
	}

	var offsetDescription = fbutils.CreateStringOffset(fbb, obj.DescribedObject.Description)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetManufacturer = fbutils.CreateStringOffset(fbb, obj.Manufacturer)
	var offsetModel = fbutils.CreateStringOffset(fbb, obj.Model)
	var offsetLabels = fbutils.CreateStringVectorOffset(fbb, obj.Labels)
	var offsetDeviceResources = fbutils.CreateByteVectorOffset(fbb, deviceResourcesJsonToDatabaseValue(obj.DeviceResources))
	var offsetDeviceCommands = fbutils.CreateByteVectorOffset(fbb, profileResourcesJsonToDatabaseValue(obj.DeviceCommands))

	// build the FlatBuffers object
	fbb.StartObject(11)
	fbutils.SetInt64Slot(fbb, 0, obj.DescribedObject.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.DescribedObject.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.DescribedObject.Timestamps.Origin)
	fbutils.SetUOffsetTSlot(fbb, 3, offsetDescription)
	fbutils.SetUint64Slot(fbb, 4, id)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetManufacturer)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetModel)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetLabels)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetDeviceResources)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetDeviceCommands)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (deviceProfile_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(12, 0)

	var relCoreCommands []Command
	if err := txn.RunWithCursor(DeviceProfileBinding.Id, func(cursor *objectbox.Cursor) error {
		if rSlice, err := cursor.RelationGetAll(1, CommandBinding.Id, id); err != nil {
			return err
		} else {
			relCoreCommands = rSlice.([]Command)
			return nil
		}
	}); err != nil {
		return nil, err
	}

	return &DeviceProfile{
		DescribedObject: models.DescribedObject{
			Timestamps: Timestamps{
				Created:  table.GetInt64Slot(4, 0),
				Modified: table.GetInt64Slot(6, 0),
				Origin:   table.GetInt64Slot(8, 0),
			},
			Description: fbutils.GetStringSlot(table, 10),
		},
		Id:              objectbox.StringIdConvertToEntityProperty(id),
		Name:            fbutils.GetStringSlot(table, 14),
		Manufacturer:    fbutils.GetStringSlot(table, 16),
		Model:           fbutils.GetStringSlot(table, 18),
		Labels:          fbutils.GetStringVectorSlot(table, 20),
		DeviceResources: deviceResourcesJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 22)),
		DeviceCommands:  profileResourcesJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 24)),
		CoreCommands:    relCoreCommands,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (deviceProfile_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]DeviceProfile, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (deviceProfile_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]DeviceProfile), *object.(*DeviceProfile))
}

// Box provides CRUD access to DeviceProfile objects
type DeviceProfileBox struct {
	*objectbox.Box
}

// BoxForDeviceProfile opens a box of DeviceProfile objects
func BoxForDeviceProfile(ob *objectbox.ObjectBox) *DeviceProfileBox {
	return &DeviceProfileBox{
		Box: ob.InternalBox(5),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the DeviceProfile.Id property on the passed object will be assigned the new ID as well.
func (box *DeviceProfileBox) Put(object *DeviceProfile) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the DeviceProfile.Id property on the passed object will be assigned the new ID as well.
//
// It's executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "Put & Forget:" you gain faster puts as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
//
// In situations with (extremely) high async load, this method may be throttled (~1ms) or delayed (<1s).
// In the unlikely event that the object could not be enqueued after delaying, an error will be returned.
//
// Note that this method does not give you hard durability guarantees like the synchronous Put provides.
// There is a small time window (typically 3 ms) in which the data may not have been committed durably yet.
func (box *DeviceProfileBox) PutAsync(object *DeviceProfile) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the DeviceProfile.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the DeviceProfile.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *DeviceProfileBox) PutAll(objects []DeviceProfile) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *DeviceProfileBox) Get(id uint64) (*DeviceProfile, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*DeviceProfile), nil
}

// Get reads all stored objects
func (box *DeviceProfileBox) GetAll() ([]DeviceProfile, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceProfile), nil
}

// Remove deletes a single object
func (box *DeviceProfileBox) Remove(object *DeviceProfile) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.Id))
}

// Creates a query with the given conditions. Use the fields of the DeviceProfile_ struct to create conditions.
// Keep the *DeviceProfileQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *DeviceProfileBox) Query(conditions ...objectbox.Condition) *DeviceProfileQuery {
	return &DeviceProfileQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the DeviceProfile_ struct to create conditions.
// Keep the *DeviceProfileQuery if you intend to execute the query multiple times.
func (box *DeviceProfileBox) QueryOrError(conditions ...objectbox.Condition) (*DeviceProfileQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &DeviceProfileQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all DeviceProfile which Id is either 42 or 47:
// 		box.Query(DeviceProfile_.Id.In(42, 47)).Find()
type DeviceProfileQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *DeviceProfileQuery) Find() ([]DeviceProfile, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]DeviceProfile), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *DeviceProfileQuery) Offset(offset uint64) *DeviceProfileQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *DeviceProfileQuery) Limit(limit uint64) *DeviceProfileQuery {
	query.Query.Limit(limit)
	return query
}
