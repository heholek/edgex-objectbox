// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type reading_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var ReadingBinding = reading_EntityInfo{
	Entity: objectbox.Entity{
		Id: 4,
	},
	Uid: 8012333558975129213,
}

// Reading_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Reading_ = struct {
	Id          *objectbox.PropertyUint64
	Pushed      *objectbox.PropertyInt64
	Created     *objectbox.PropertyInt64
	Origin      *objectbox.PropertyInt64
	Modified    *objectbox.PropertyInt64
	Device      *objectbox.PropertyString
	Name        *objectbox.PropertyString
	Value       *objectbox.PropertyString
	BinaryValue *objectbox.PropertyByteVector
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &ReadingBinding.Entity,
		},
	},
	Pushed: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &ReadingBinding.Entity,
		},
	},
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &ReadingBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &ReadingBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &ReadingBinding.Entity,
		},
	},
	Device: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &ReadingBinding.Entity,
		},
	},
	Name: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     7,
			Entity: &ReadingBinding.Entity,
		},
	},
	Value: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     8,
			Entity: &ReadingBinding.Entity,
		},
	},
	BinaryValue: &objectbox.PropertyByteVector{
		BaseProperty: &objectbox.BaseProperty{
			Id:     9,
			Entity: &ReadingBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (reading_EntityInfo) GeneratorVersion() int {
	return 3
}

// AddToModel is called by ObjectBox during model build
func (reading_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Reading", 4, 8012333558975129213)
	model.Property("Id", 6, 1, 5443847882081660610)
	model.PropertyFlags(8193)
	model.Property("Pushed", 6, 2, 1634534132378761166)
	model.Property("Created", 6, 3, 6022679825524196879)
	model.Property("Origin", 6, 4, 4564442774242088945)
	model.Property("Modified", 6, 5, 3044189181647527381)
	model.Property("Device", 9, 6, 960087973838288432)
	model.Property("Name", 9, 7, 8155400751319283038)
	model.Property("Value", 9, 8, 5963400646570440874)
	model.Property("BinaryValue", 23, 9, 7724532463827666508)
	model.EntityLastPropertyId(9, 7724532463827666508)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (reading_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Reading); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.Id), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Reading).Id), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (reading_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Reading); ok {
		obj.Id = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Reading).Id
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (reading_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (reading_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *Reading
	if objPtr, ok := object.(*Reading); ok {
		obj = objPtr
	} else {
		objVal := object.(Reading)
		obj = &objVal
	}

	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)
	var offsetName = fbutils.CreateStringOffset(fbb, obj.Name)
	var offsetValue = fbutils.CreateStringOffset(fbb, obj.Value)
	var offsetBinaryValue = fbutils.CreateByteVectorOffset(fbb, obj.BinaryValue)

	// build the FlatBuffers object
	fbb.StartObject(9)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Pushed)
	fbutils.SetInt64Slot(fbb, 2, obj.Created)
	fbutils.SetInt64Slot(fbb, 3, obj.Origin)
	fbutils.SetInt64Slot(fbb, 4, obj.Modified)
	fbutils.SetUOffsetTSlot(fbb, 5, offsetDevice)
	fbutils.SetUOffsetTSlot(fbb, 6, offsetName)
	fbutils.SetUOffsetTSlot(fbb, 7, offsetValue)
	fbutils.SetUOffsetTSlot(fbb, 8, offsetBinaryValue)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (reading_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(4, 0)

	return &Reading{
		Id:          objectbox.StringIdConvertToEntityProperty(id),
		Pushed:      fbutils.GetInt64Slot(table, 6),
		Created:     fbutils.GetInt64Slot(table, 8),
		Origin:      fbutils.GetInt64Slot(table, 10),
		Modified:    fbutils.GetInt64Slot(table, 12),
		Device:      fbutils.GetStringSlot(table, 14),
		Name:        fbutils.GetStringSlot(table, 16),
		Value:       fbutils.GetStringSlot(table, 18),
		BinaryValue: fbutils.GetByteVectorSlot(table, 20),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (reading_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Reading, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (reading_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Reading), *object.(*Reading))
}

// Box provides CRUD access to Reading objects
type ReadingBox struct {
	*objectbox.Box
}

// BoxForReading opens a box of Reading objects
func BoxForReading(ob *objectbox.ObjectBox) *ReadingBox {
	return &ReadingBox{
		Box: ob.InternalBox(4),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Reading.Id property on the passed object will be assigned the new ID as well.
func (box *ReadingBox) Put(object *Reading) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Reading.Id property on the passed object will be assigned the new ID as well.
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
func (box *ReadingBox) PutAsync(object *Reading) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Reading.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Reading.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *ReadingBox) PutMany(objects []Reading) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *ReadingBox) Get(id uint64) (*Reading, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Reading), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *ReadingBox) GetMany(ids ...uint64) ([]Reading, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Reading), nil
}

// GetAll reads all stored objects
func (box *ReadingBox) GetAll() ([]Reading, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Reading), nil
}

// Remove deletes a single object
func (box *ReadingBox) Remove(object *Reading) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *ReadingBox) RemoveMany(objects ...*Reading) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = objectbox.StringIdConvertToDatabaseValue(object.Id)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Reading_ struct to create conditions.
// Keep the *ReadingQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *ReadingBox) Query(conditions ...objectbox.Condition) *ReadingQuery {
	return &ReadingQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Reading_ struct to create conditions.
// Keep the *ReadingQuery if you intend to execute the query multiple times.
func (box *ReadingBox) QueryOrError(conditions ...objectbox.Condition) (*ReadingQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &ReadingQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Reading which Id is either 42 or 47:
// 		box.Query(Reading_.Id.In(42, 47)).Find()
type ReadingQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *ReadingQuery) Find() ([]Reading, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Reading), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *ReadingQuery) Offset(offset uint64) *ReadingQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *ReadingQuery) Limit(limit uint64) *ReadingQuery {
	query.Query.Limit(limit)
	return query
}
