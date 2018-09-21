package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/flatcoredata"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	. "github.com/objectbox/objectbox-go/objectbox"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

// This file could be generated in the future
type EventBinding struct {
}

func (EventBinding) AddToModel(model *Model) {
	model.Entity("Event", 1, 10001)
	model.Property("id", PropertyType_Long, 1, 10001001)
	model.PropertyFlags(PropertyFlags_ID)
	model.Property("pushed", PropertyType_Long, 2, 10001002)
	model.Property("device", PropertyType_String, 3, 10001003)
	model.Property("created", PropertyType_Long, 4, 10001004)
	model.Property("modified", PropertyType_Long, 5, 10001005)
	model.Property("origin", PropertyType_Long, 6, 10001006)
	model.Property("scheduleEvent", PropertyType_String, 7, 10001007)
	model.EntityLastPropertyId(7, 10001007)
}

func (EventBinding) GetId(object interface{}) (id uint64, err error) {
	event, ok := object.(*models.Event)
	if !ok {
		// Programming error, OK to panic
		panic("Object has wrong type")
	}
	idString := string(event.ID)
	if idString == "" {
		return 0, nil
	}
	return strconv.ParseUint(idString, 10, 64)
}

func (EventBinding) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	flattenModelEvent(object.(*models.Event), fbb, id)
}

func flattenModelEvent(event *models.Event, fbb *flatbuffers.Builder, id uint64) {
	offsetDevice := Unavailable
	if event.Device != "" {
		offsetDevice = fbb.CreateString(event.Device)
	}

	flatcoredata.EventStart(fbb)

	flatcoredata.EventAddId(fbb, id)
	if offsetDevice != Unavailable {
		flatcoredata.EventAddDevice(fbb, offsetDevice)
	}
}

func (EventBinding) ToObject(bytes []byte) interface{} {
	flatEvent := flatcoredata.GetRootAsEvent(bytes, flatbuffers.UOffsetT(0))
	return toModelEvent(flatEvent)
}

func toModelEvent(src *flatcoredata.Event) *models.Event {
	return &models.Event{
		ID:       bson.ObjectId(strconv.FormatUint(src.Id(), 10)),
		Pushed:   src.Pushed(),
		Created:  src.Created(),
		Origin:   src.Origin(),
		Modified: src.Modified(),
		Device:   string(src.Device()),
		Event:    string(src.ScheduleEvent()),
	}
}

func (EventBinding) MakeSlice(capacity int) interface{} {
	return make([]models.Event, 0, capacity)
}

func (EventBinding) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]models.Event), *object.(*models.Event))
}
