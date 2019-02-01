package obx

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/objectbox/objectbox-go/objectbox"
	"strconv"
)

func IdToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
func IdFromString(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}

func bsonIdToEntityProperty(dbValue uint64) bson.ObjectId {
	return bson.ObjectId(objectbox.StringIdConvertToEntityProperty(dbValue))
}

func bsonIdToDatabaseValue(goValue bson.ObjectId) uint64 {
	return objectbox.StringIdConvertToDatabaseValue(string(goValue))
}

func interfaceJsonToEntityProperty(dbValue []byte) interface{} {
	if dbValue == nil {
		return nil
	}

	var value interface{}

	if err := json.Unmarshal(dbValue, &value); err != nil {
		panic(err)
	}

	return value
}

func interfaceJsonToDatabaseValue(goValue interface{}) []byte {
	if goValue == nil {
		return nil
	}

	if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}
