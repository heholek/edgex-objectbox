package objectbox

import (
	"strconv"
)

func idObxToString(id uint64) string {
	return strconv.FormatUint(id, 10)
}
func idStringToObx(id string) (uint64, error) {
	return strconv.ParseUint(id, 10, 64)
}
