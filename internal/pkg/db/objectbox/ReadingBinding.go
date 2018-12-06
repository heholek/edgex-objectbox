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
type ReadingBinding struct {
	indexDevice bool
}

func (binding ReadingBinding) SetId(object interface{}, id uint64) error {
	object.(*models.Reading).Id = bson.ObjectId(strconv.FormatUint(id, 10))
	return nil
}

func (binding ReadingBinding) GeneratorVersion() int {
	return 1
}

func (binding ReadingBinding) AddToModel(model *Model) {
	model.Entity("Reading", 2, 10002)
	model.Property("id", PropertyType_Long, 1, 10002001)
	model.PropertyFlags(PropertyFlags_ID)
	model.Property("eventId", PropertyType_Long, 2, 10002002)
	//model.Property("eventId", PropertyType_Relation, 2, 10002002)
	//model.PropertyFlags(PropertyFlags_INDEXED)
	model.Property("pushed", PropertyType_Long, 3, 10002003)
	model.Property("created", PropertyType_Long, 4, 10002004)
	model.Property("origin", PropertyType_Long, 5, 10002005)
	model.Property("modified", PropertyType_Long, 6, 10002006)

	model.Property("device", PropertyType_String, 7, 10002007)
	if binding.indexDevice {
		model.PropertyFlags(PropertyFlags_INDEXED)
		model.PropertyIndex(1, 20002007)
	}

	model.Property("name", PropertyType_String, 8, 10002008)
	model.Property("value", PropertyType_String, 9, 10002009)
	model.EntityLastPropertyId(9, 10002009)
}

func (ReadingBinding) GetId(object interface{}) (id uint64, err error) {
	idString := string(object.(*models.Reading).Id)
	if idString == "" {
		return 0, nil
	}
	return strconv.ParseUint(idString, 10, 64)
}

func (ReadingBinding) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) {
	flattenModelReading(object.(*models.Reading), fbb, id)
}

func flattenModelReading(reading *models.Reading, fbb *flatbuffers.Builder, id uint64) {
	offsetDevice := Unavailable
	if reading.Device != "" {
		offsetDevice = fbb.CreateString(reading.Device)
	}
	offsetName := Unavailable
	if reading.Name != "" {
		offsetName = fbb.CreateString(reading.Name)
	}
	offsetValue := Unavailable
	if reading.Value != "" {
		offsetValue = fbb.CreateString(reading.Value)
	}

	flatcoredata.ReadingStart(fbb)

	flatcoredata.ReadingAddId(fbb, id)
	flatcoredata.ReadingAddCreated(fbb, reading.Created)
	flatcoredata.ReadingAddOrigin(fbb, reading.Origin)
	flatcoredata.ReadingAddModified(fbb, reading.Modified)
	flatcoredata.ReadingAddPushed(fbb, reading.Pushed)

	if offsetDevice != Unavailable {
		flatcoredata.ReadingAddDevice(fbb, offsetDevice)
	}
	if offsetName != Unavailable {
		flatcoredata.ReadingAddName(fbb, offsetName)
	}
	if offsetValue != Unavailable {
		flatcoredata.ReadingAddValue(fbb, offsetValue)
	}
}

func (ReadingBinding) ToObject(bytes []byte) interface{} {
	flatReading := flatcoredata.GetRootAsReading(bytes, flatbuffers.UOffsetT(0))
	return toModelReading(flatReading)
}

func (ReadingBinding) MakeSlice(capacity int) interface{} {
	return make([]models.Reading, 0, capacity)
}

func (ReadingBinding) AppendToSlice(slice interface{}, object interface{}) (sliceNew interface{}) {
	return append(slice.([]models.Reading), *object.(*models.Reading))
}

func toModelReadingFromBytes(bytesData []byte) *models.Reading {
	flatReading := flatcoredata.GetRootAsReading(bytesData, flatbuffers.UOffsetT(0))
	return toModelReading(flatReading)
}

func toModelReading(src *flatcoredata.Reading) *models.Reading {
	return &models.Reading{
		Id:       bson.ObjectId(strconv.FormatUint(src.Id(), 10)),
		Pushed:   src.Pushed(),
		Created:  src.Created(),
		Origin:   src.Origin(),
		Modified: src.Modified(),
		Device:   string(src.Device()),
		Name:     string(src.Name()),
		Value:    string(src.Value()),
	}

}
