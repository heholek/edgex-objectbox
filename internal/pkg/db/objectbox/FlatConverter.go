package objectbox

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/flatcoredata"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/google/flatbuffers/go"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

const Unavailable = flatbuffers.UOffsetT(0)

func flattenEntity(event *models.Event, fbb *flatbuffers.Builder, id uint64) {
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

func flattenReading(reading models.Reading, fbb *flatbuffers.Builder, id uint64) {
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
