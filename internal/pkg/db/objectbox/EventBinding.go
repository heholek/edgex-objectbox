package objectbox

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
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
	flattenEntity(object.(*models.Event), fbb, id)
}
