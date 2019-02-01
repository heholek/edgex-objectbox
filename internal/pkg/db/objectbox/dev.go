package objectbox

import (
	"encoding/hex"
	"fmt"
	"path/filepath"
	"runtime"
)

func notImplemented() string {
	pc, _, _, _ := runtime.Caller(1)
	fun := filepath.Ext(runtime.FuncForPC(pc).Name())[1:]
	return fmt.Sprintf("method %s is not implemented", fun)
}

// FIXME temporary while some tests still call id.Hex()
// after that is fixed, this function should be removed, it's calls would be useless
func idFromHex(id string) (string, error) {
	decoded, err := hex.DecodeString(id)
	return string(decoded), err
}
