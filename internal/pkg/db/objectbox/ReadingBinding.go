package objectbox

import (
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"strconv"
)

type ReadingBinding struct {
}

func (ReadingBinding) GetTypeId() TypeId {
	return 2
}

func (ReadingBinding) GetTypeName() string {
	return "Reading"
}

func (ReadingBinding) GetId(object interface{}) (id uint64, err error) {
	reading, ok := object.(*models.Reading)
	if !ok {
		// Programming error, OK to panic
		panic("Object has wrong type")
	}
	idString := string(reading.Id)
	if idString == "" {
		return 0, nil
	}
	return strconv.ParseUint(idString, 10, 64)
}

func (ReadingBinding) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	flattenReading(object.(*models.Reading), fbb, id)
}
