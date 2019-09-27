// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package obx

import (
	. "github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type event_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var EventBinding = event_EntityInfo{
	Entity: objectbox.Entity{
		Id: 8,
	},
	Uid: 5261868944228209948,
}

// Event_ contains type-based Property helpers to facilitate some common operations such as Queries.
var Event_ = struct {
	ID       *objectbox.PropertyUint64
	Pushed   *objectbox.PropertyInt64
	Device   *objectbox.PropertyString
	Created  *objectbox.PropertyInt64
	Modified *objectbox.PropertyInt64
	Origin   *objectbox.PropertyInt64
	Readings *objectbox.RelationToMany
}{
	ID: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &EventBinding.Entity,
		},
	},
	Pushed: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &EventBinding.Entity,
		},
	},
	Device: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     3,
			Entity: &EventBinding.Entity,
		},
	},
	Created: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     4,
			Entity: &EventBinding.Entity,
		},
	},
	Modified: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     5,
			Entity: &EventBinding.Entity,
		},
	},
	Origin: &objectbox.PropertyInt64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     6,
			Entity: &EventBinding.Entity,
		},
	},
	Readings: &objectbox.RelationToMany{
		Id:     2,
		Source: &EventBinding.Entity,
		Target: &ReadingBinding.Entity,
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (event_EntityInfo) GeneratorVersion() int {
	return 3
}

// AddToModel is called by ObjectBox during model build
func (event_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("Event", 8, 5261868944228209948)
	model.Property("ID", 6, 1, 8408127641507867598)
	model.PropertyFlags(8193)
	model.Property("Pushed", 6, 2, 5256971571136440340)
	model.Property("Device", 9, 3, 8701916792943957528)
	model.Property("Created", 6, 4, 8617743842335964745)
	model.Property("Modified", 6, 5, 9004481679338464951)
	model.Property("Origin", 6, 6, 536564806295219253)
	model.EntityLastPropertyId(6, 536564806295219253)
	model.Relation(2, 6583600503460504451, ReadingBinding.Id, ReadingBinding.Uid)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (event_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*Event); ok {
		return objectbox.StringIdConvertToDatabaseValue(obj.ID), nil
	} else {
		return objectbox.StringIdConvertToDatabaseValue(object.(Event).ID), nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (event_EntityInfo) SetId(object interface{}, id uint64) {
	if obj, ok := object.(*Event); ok {
		obj.ID = objectbox.StringIdConvertToEntityProperty(id)
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(Event).ID
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (event_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	if err := BoxForEvent(ob).RelationReplace(Event_.Readings, id, object, object.(*Event).Readings); err != nil {
		return err
	}

	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (event_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *Event
	if objPtr, ok := object.(*Event); ok {
		obj = objPtr
	} else {
		objVal := object.(Event)
		obj = &objVal
	}

	var offsetDevice = fbutils.CreateStringOffset(fbb, obj.Device)

	// build the FlatBuffers object
	fbb.StartObject(6)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetInt64Slot(fbb, 1, obj.Pushed)
	fbutils.SetUOffsetTSlot(fbb, 2, offsetDevice)
	fbutils.SetInt64Slot(fbb, 3, obj.Created)
	fbutils.SetInt64Slot(fbb, 4, obj.Modified)
	fbutils.SetInt64Slot(fbb, 5, obj.Origin)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (event_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}
	var id = table.GetUint64Slot(4, 0)

	var relReadings []Reading
	if rIds, err := BoxForEvent(ob).RelationIds(Event_.Readings, id); err != nil {
		return nil, err
	} else if rSlice, err := BoxForReading(ob).GetMany(rIds...); err != nil {
		return nil, err
	} else {
		relReadings = rSlice
	}

	return &Event{
		ID:       objectbox.StringIdConvertToEntityProperty(id),
		Pushed:   fbutils.GetInt64Slot(table, 6),
		Device:   fbutils.GetStringSlot(table, 8),
		Created:  fbutils.GetInt64Slot(table, 10),
		Modified: fbutils.GetInt64Slot(table, 12),
		Origin:   fbutils.GetInt64Slot(table, 14),
		Readings: relReadings,
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (event_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]Event, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (event_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]Event), *object.(*Event))
}

// Box provides CRUD access to Event objects
type EventBox struct {
	*objectbox.Box
}

// BoxForEvent opens a box of Event objects
func BoxForEvent(ob *objectbox.ObjectBox) *EventBox {
	return &EventBox{
		Box: ob.InternalBox(8),
	}
}

// Put synchronously inserts/updates a single object.
// In case the ID is not specified, it would be assigned automatically (auto-increment).
// When inserting, the Event.ID property on the passed object will be assigned the new ID as well.
func (box *EventBox) Put(object *Event) (uint64, error) {
	return box.Box.Put(object)
}

// PutAsync asynchronously inserts/updates a single object.
// When inserting, the Event.ID property on the passed object will be assigned the new ID as well.
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
func (box *EventBox) PutAsync(object *Event) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case IDs are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the Event.ID property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the Event.ID assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EventBox) PutMany(objects []Event) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EventBox) Get(id uint64) (*Event, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*Event), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *EventBox) GetMany(ids ...uint64) ([]Event, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]Event), nil
}

// GetAll reads all stored objects
func (box *EventBox) GetAll() ([]Event, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]Event), nil
}

// Remove deletes a single object
func (box *EventBox) Remove(object *Event) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *EventBox) RemoveMany(objects ...*Event) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = objectbox.StringIdConvertToDatabaseValue(object.ID)
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EventBox) Query(conditions ...objectbox.Condition) *EventQuery {
	return &EventQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the Event_ struct to create conditions.
// Keep the *EventQuery if you intend to execute the query multiple times.
func (box *EventBox) QueryOrError(conditions ...objectbox.Condition) (*EventQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EventQuery{query}, nil
	}
}

// Query provides a way to search stored objects
//
// For example, you can find all Event which ID is either 42 or 47:
// 		box.Query(Event_.ID.In(42, 47)).Find()
type EventQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EventQuery) Find() ([]Event, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]Event), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *EventQuery) Offset(offset uint64) *EventQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *EventQuery) Limit(limit uint64) *EventQuery {
	query.Query.Limit(limit)
	return query
}
