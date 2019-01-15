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

// decodes the given byte slice as a complex number
func interfaceGobToEntityProperty(dbValue []byte) interface{} {
	// NOTE that constructing the decoder each time is inefficient and only serves as an example for the property converters
	var b = bytes.NewBuffer(dbValue)
	var decoder = gob.NewDecoder(b)

	var value complex128
	if err := decoder.Decode(&value); err != nil {
		panic(err)
	}

	return value
}

// encodes the given complex number as a byte slice
func interfaceGobToDatabaseValue(goValue interface{}) []byte {
	// NOTE that constructing the encoder each time is inefficient and only serves as an example for the property converters
	var b bytes.Buffer
	var encoder = gob.NewEncoder(&b)

	if err := encoder.Encode(goValue); err != nil {
		panic(err)
	}

	return b.Bytes()
}
