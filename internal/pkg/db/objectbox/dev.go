package objectbox

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func notImplemented() string {
	pc, _, _, _ := runtime.Caller(1)
	fun := filepath.Ext(runtime.FuncForPC(pc).Name())[1:]
	return fmt.Sprintf("method %s is not implemented", fun)
}
