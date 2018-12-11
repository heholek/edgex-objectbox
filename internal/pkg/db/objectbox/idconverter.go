package objectbox

import (
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

func idObxToBson(id uint64) bson.ObjectId {
	return bson.ObjectId(strconv.FormatUint(id, 10))
}
func idStringToObx(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}
