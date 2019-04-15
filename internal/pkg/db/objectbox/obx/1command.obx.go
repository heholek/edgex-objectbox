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

type command_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var CommandBinding = command_EntityInfo{
	Entity: objectbox.Entity{
		Id: 2,
	},
	Uid: 3466110984159220104,
}

// Command_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Command_ = struct {
	Created            *objectbox.PropertyInt64
	Modified           *objectbox.PropertyInt64
	Origin             *objectbox.PropertyInt64
	Id                 *objectbox.PropertyUint64
	Name               *objectbox.PropertyString
	Get_Path           *objectbox.PropertyString
	Get_Responses      *objectbox.PropertyByteVector
	Get_URL            *objectbox.PropertyString
	Put_Path           *objectbox.PropertyString
	Put_Responses      *objectbox.PropertyByteVector
	Put_URL            *objectbox.PropertyString
	Put_ParameterNames *objectbox.PropertyStringVector
}{
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &CommandBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &CommandBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &CommandBinding.Entity,
		},
	},
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &CommandBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &CommandBinding.Entity,
		},
	},
	Get_Path: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &CommandBinding.Entity,
		},
	},
	Get_Responses: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &CommandBinding.Entity,
		},
	},
	Get_URL: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &CommandBinding.Entity,
		},
	},
	Put_Path: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &CommandBinding.Entity,
		},
	},
	Put_Responses: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     10,
			Entity: &CommandBinding.Entity,
		},
	},
	Put_URL: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     11,
			Entity: &CommandBinding.Entity,
		},
	},
	Put_ParameterNames: &objectbox.PropertyStringVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     12,
			Entity: &CommandBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (command_EntityInfo) GeneratorVersion() int {
	return 2
}

// AddToModel is called by ObjectBox during model build
func (command_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Command", 2, 3466110984159220104)
	model.Property("Created", objectbox.PropertyType_Long, 1, 8976154384675543155)
	model.Property("Modified", objectbox.PropertyType_Long, 2, 4173457774608518837)
	model.Property("Origin", objectbox.PropertyType_Long, 3, 2804731238210135713)
	model.Property("Id", objectbox.PropertyType_Long, 4, 7187431837387194143)
	model.PropertyFlags(objectbox.PropertyFlags_ID | objectbox.PropertyFlags_UNSIGNED)
	model.Property("Name", objectbox.PropertyType_String, 5, 1838205786009926106)
	model.Property("Get_Path", objectbox.PropertyType_String, 6, 4675672987288618325)
	model.Property("Get_Responses", objectbox.PropertyType_ByteVector, 7, 6760659952611457052)
	model.Property("Get_URL", objectbox.PropertyType_String, 8, 3771849917208761361)
	model.Property("Put_Path", objectbox.PropertyType_String, 9, 101551801493008702)
	model.Property("Put_Responses", objectbox.PropertyType_ByteVector, 10, 7143390721286976146)
	model.Property("Put_URL", objectbox.PropertyType_String, 11, 5130326444177191018)
	model.Property("Put_ParameterNames", objectbox.PropertyType_StringVector, 12, 706434787225063093)
	model.EntityLastPropertyId(12, 706434787225063093)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (command_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Command); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Command).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (command_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Command); ok {
		obj.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Command).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (command_EntityInfo) PutRelated(txn *objectbox.Transaction, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (command_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *Command
	if objPtr, ok := object.(*Command); ok {
		obj = objPtr
	} else {
		objVal := object.(Command)
		obj = &objVal
	}

	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetGet_Path = fbutils.CreateStringOffset(fbb, obj.Get.Action.Path)
	var offsetGet_Responses = fbutils.CreateByteVectorOffset(fbb, responsesJsonToDatabaseValue(obj.Get.Action.Responses))
	var offsetGet_URL = fbutils.CreateStringOffset(fbb, obj.Get.Action.URL)
	var offsetPut_Path = fbutils.CreateStringOffset(fbb, obj.Put.Action.Path)
	var offsetPut_Responses = fbutils.CreateByteVectorOffset(fbb, responsesJsonToDatabaseValue(obj.Put.Action.Responses))
	var offsetPut_URL = fbutils.CreateStringOffset(fbb, obj.Put.Action.URL)
	var offsetPut_ParameterNames = fbutils.CreateStringVectorOffset(fbb, obj.Put.ParameterNames)

	// build the FlatBuffers object
	fbb.StartObject(12)
	fbutils.SetInt64Slot(fbb, 0, obj.Timestamps.Created)
	fbutils.SetInt64Slot(fbb, 1, obj.Timestamps.Modified)
	fbutils.SetInt64Slot(fbb, 2, obj.Timestamps.Origin)
	fbutils.SetUint64Slot(fbb, 3, id)
	fbutils.SetUOffsetTSlot(fbb, 4, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetGet_Path)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetGet_Responses)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetGet_URL)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetPut_Path)
	fbutils.SetUOffsetTSlot(fbb, 9, offsetPut_Responses)
	fbutils.SetUOffsetTSlot(fbb, 10, offsetPut_URL)
	fbutils.SetUOffsetTSlot(fbb, 11, offsetPut_ParameterNames)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (command_EntityInfo) Load(txn *objectbox.Transaction, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(10, 0)

	return &Command{
		Timestamps: models.Timestamps{
			Created:  table.GetInt64Slot(4, 0),
			Modified: table.GetInt64Slot(6, 0),
			Origin:   table.GetInt64Slot(8, 0),
		},
		Id:   objectbox.StringIdConvertToEntityProperty(id),
		Name: fbutils.GetStringSlot(table, 12),
		Get: &Get{
			Action: Action{
				Path:      fbutils.GetStringSlot(table, 14),
				Responses: responsesJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 16)),
				URL:       fbutils.GetStringSlot(table, 18),
			},
		},
		Put: &Put{
			Action: Action{
				Path:      fbutils.GetStringSlot(table, 20),
				Responses: responsesJsonToEntityProperty(fbutils.GetByteVectorSlot(table, 22)),
				URL:       fbutils.GetStringSlot(table, 24),
			},
			ParameterNames: fbutils.GetStringVectorSlot(table, 26),
		},
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (command_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Command, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (command_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Command), *object.(*Command))
}

// Box provides CRUD access to Command objects
type CommandBox struct {
	*objectbox.Box
}

// BoxForCommand opens a box of Command objects
func BoxForCommand(ob *objectbox.ObjectBox) *CommandBox {
	return &CommandBox{
		Box: ob.InternalBox(2),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Command.Id property on the passed object will be assigned the new ID as well.
func (box *CommandBox) Put(object *Command) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Command.Id property on the passed object will be assigned the new ID as well.
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
func (box *CommandBox) PutAsync(object *Command) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutAll inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Command.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Command.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *CommandBox) PutAll(objects []Command) ([]uint64, error) {
	return box.Box.PutAll(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *CommandBox) Get(id uint64) (*Command, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Command), nil
}

// Get reads all stored objects
func (box *CommandBox) GetAll() ([]Command, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Command), nil
}

// Remove deletes a single object
func (box *CommandBox) Remove(object *Command) (err error) {
	return box.Box.Remove(objectbox.StringIdConvertToDatabaseValue(object.Id))
}

// Creates a query with the given conditions. Use the fields of the Command_ struct to create conditions.
// Keep the *CommandQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *CommandBox) Query(conditions ...objectbox.Condition) *CommandQuery {
	return &CommandQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Command_ struct to create conditions.
// Keep the *CommandQuery if you intend to execute the query multiple times.
func (box *CommandBox) QueryOrError(conditions ...objectbox.Condition) (*CommandQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &CommandQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Command which Id is either 42 or 47:
// 		box.Query(Command_.Id.In(42, 47)).Find()
type CommandQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *CommandQuery) Find() ([]Command, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Command), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *CommandQuery) Offset(offset uint64) *CommandQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *CommandQuery) Limit(limit uint64) *CommandQuery {
	query.Query.Limit(limit)
	return query
}
