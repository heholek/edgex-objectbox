package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/flatcoredata"
	. "github.com/edgexfoundry/edgex-go/internal/pkg/objectbox"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

type EventBinding struct {
}

func (EventBinding) GetTypeId() TypeId {
	return 1
}

func (EventBinding) GetTypeName() string {
	return "Event"
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

func (EventBinding) AppendToSlice(slice interface{}, object interface{}) interface{} {
	if slice == nil {
		slice = make([]models.Event, 0, 16)
	}
	return append(slice.([]models.Event), *object.(*models.Event))
}
