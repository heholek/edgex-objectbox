package obx

import (
	"encoding/json"
	"github.com/edgexfoundry/go-mod-core-contracts/models"
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
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

func interfaceJsonToEntityProperty(dbValue []byte) (interface{}, error) {
	if dbValue == nil {
		return nil, nil
	}

	var value interface{}
	if err := json.Unmarshal(dbValue, &value); err != nil {
		return nil, err
	}
	return value, nil
}

func interfaceJsonToDatabaseValue(goValue interface{}) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func mapStringStringJsonToEntityProperty(dbValue []byte) (result map[string]string, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func mapStringStringJsonToDatabaseValue(goValue map[string]string) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func mapStringProtocolPropertiesJsonToEntityProperty(dbValue []byte) (result map[string]models.ProtocolProperties, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func mapStringProtocolPropertiesJsonToDatabaseValue(goValue map[string]models.ProtocolProperties) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func notificationsCategoryToDatabaseValue(goValue []models.NotificationsCategory) ([]string, error) {
	if goValue == nil {
		return nil, nil
	}

	var result = make([]string, len(goValue))

	for k, v := range goValue {
		result[k] = string(v)
	}

	return result, nil
}

func notificationsCategoryToEntityProperty(dbValue []string) ([]models.NotificationsCategory, error) {
	if dbValue == nil {
		return nil, nil
	}

	var result = make([]models.NotificationsCategory, len(dbValue))

	for k, v := range dbValue {
		result[k] = models.NotificationsCategory(v)
	}

	return result, nil
}

func responsesJsonToEntityProperty(dbValue []byte) (result []models.Response, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func responsesJsonToDatabaseValue(goValue []models.Response) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func channelsJsonToEntityProperty(dbValue []byte) (result []models.Channel, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, err
}

func channelsJsonToDatabaseValue(goValue []models.Channel) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func transmissionRecordsJsonToEntityProperty(dbValue []byte) (result []models.TransmissionRecord, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func transmissionRecordsJsonToDatabaseValue(goValue []models.TransmissionRecord) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func deviceResourcesJsonToEntityProperty(dbValue []byte) (result []models.DeviceResource, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func deviceResourcesJsonToDatabaseValue(goValue []models.DeviceResource) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func autoEventsJsonToEntityProperty(dbValue []byte) (result []models.AutoEvent, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func autoEventsJsonToDatabaseValue(goValue []models.AutoEvent) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}

func profileResourcesJsonToEntityProperty(dbValue []byte) (result []models.ProfileResource, err error) {
	if dbValue == nil {
		return nil, nil
	}

	if err := json.Unmarshal(dbValue, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func profileResourcesJsonToDatabaseValue(goValue []models.ProfileResource) ([]byte, error) {
	if goValue == nil {
		return nil, nil
	}

	return json.Marshal(goValue)
}
