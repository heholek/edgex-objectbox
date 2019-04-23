package obx

import (
	"encoding/json"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"strconv"
)

func IdToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}

func IdFromString(id string) (uint64, error) {
	// using db.ErrNotFound instead of db.ErrInvalidObjectId for compatibility with EdgeX test definitions (uuid IDs)
	var errInvalid = db.ErrNotFound

	var result uint64
	if id == "" {
		result = 0
	} else if id, err := strconv.ParseUint(id, 10, 64); err != nil {
		return 0, errInvalid
	} else {
		result = id
	}

	if result == 0 {
		return 0, errInvalid
	}
	return result, nil
}

func interfaceJsonToEntityProperty(dbValue []byte) interface{} {
	if dbValue == nil {
		return nil
	}

	var value interface{}
	if err := json.Unmarshal(dbValue, &value); err != nil {
		panic(err)
	} else {
		return value
	}
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

func mapStringStringJsonToEntityProperty(dbValue []byte) (result map[string]string) {
	if dbValue == nil {
		return nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		panic(err)
	} else {
		return result
	}
}

func mapStringStringJsonToDatabaseValue(goValue map[string]string) []byte {
	if goValue == nil {
		return nil
	}

	if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func mapStringProtocolPropertiesJsonToEntityProperty(dbValue []byte) (result map[string]models.ProtocolProperties) {
	if dbValue == nil {
		return nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		panic(err)
	} else {
		return result
	}
}

func mapStringProtocolPropertiesJsonToDatabaseValue(goValue map[string]models.ProtocolProperties) []byte {
	if goValue == nil {
		return nil
	}

	if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func notificationsCategoryToDatabaseValue(goValue []models.NotificationsCategory) []string {
	if goValue == nil {
		return nil
	}

	var result = make([]string, len(goValue))

	for k, v := range goValue {
		result[k] = string(v)
	}

	return result
}

func notificationsCategoryToEntityProperty(dbValue []string) []models.NotificationsCategory {
	if dbValue == nil {
		return nil
	}

	var result = make([]models.NotificationsCategory, len(dbValue))

	for k, v := range dbValue {
		result[k] = models.NotificationsCategory(v)
	}

	return result
}

func responsesJsonToEntityProperty(dbValue []byte) (result []models.Response) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func responsesJsonToDatabaseValue(goValue []models.Response) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func channelsJsonToEntityProperty(dbValue []byte) (result []models.Channel) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func channelsJsonToDatabaseValue(goValue []models.Channel) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func transmissionRecordsJsonToEntityProperty(dbValue []byte) (result []models.TransmissionRecord) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func transmissionRecordsJsonToDatabaseValue(goValue []models.TransmissionRecord) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func deviceResourcesJsonToEntityProperty(dbValue []byte) (result []models.DeviceResource) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func deviceResourcesJsonToDatabaseValue(goValue []models.DeviceResource) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func autoEventsJsonToEntityProperty(dbValue []byte) (result []models.AutoEvent) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func autoEventsJsonToDatabaseValue(goValue []models.AutoEvent) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}

func profileResourcesJsonToEntityProperty(dbValue []byte) (result []models.ProfileResource) {
	if dbValue != nil {
		if err := json.Unmarshal(dbValue, &result); err != nil {
			panic(err)
		}
	}

	return
}

func profileResourcesJsonToDatabaseValue(goValue []models.ProfileResource) []byte {
	if goValue == nil {
		return nil
	} else if bytes, err := json.Marshal(goValue); err != nil {
		panic(err)
	} else {
		return bytes
	}
}
