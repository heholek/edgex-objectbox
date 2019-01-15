package obx

import (
	"bytes"
	"encoding/gob"
	"strconv"
)

func IdToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
func IdFromString(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}

// TODO benchmark whether it's faster to construct encoder or use a global one with a mutex

func interfaceGobToEntityProperty(dbValue []byte) interface{} {
	if dbValue == nil {
		return nil
	}

	var b = bytes.NewBuffer(dbValue)
	var decoder = gob.NewDecoder(b)

	var value interface{}
	if err := decoder.Decode(&value); err != nil {
		panic(err)
	}

	return value
}

func interfaceGobToDatabaseValue(goValue interface{}) []byte {
	if goValue == nil {
		return nil
	}

	var b bytes.Buffer
	var encoder = gob.NewEncoder(&b)

	if err := encoder.Encode(goValue); err != nil {
		panic(err)
	}

	return b.Bytes()
}
