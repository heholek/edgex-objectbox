package objectbox

/*
#cgo LDFLAGS: -L ${SRCDIR}/libs -lobjectboxc
#include <stdlib.h>
#include <string.h>
#include "objectbox.h"
*/
import "C"

import (
	"github.com/google/flatbuffers/go"
	"unsafe"
)

type Box struct {
	box     *C.OB_box
	binding ObjectBinding
	// FIXME not synchronized:
	fbb *flatbuffers.Builder
}

func (box *Box) Destroy() (err error) {
	rc := C.ob_box_destroy(box.box)
	box.box = nil
	if rc != 0 {
		err = createError()
	}
	return
}

func (box *Box) IdForPut(idCandidate uint64) (id uint64, err error) {
	id = uint64(C.ob_box_id_for_put(box.box, C.uint64_t(idCandidate)))
	if id == 0 {
		err = createError()
	}
	return
}

func (box *Box) PutAsync(object interface{}) (id uint64, err error) {
	idFromObject, err := box.binding.GetId(object)
	if err != nil {
		return
	}
	checkForPreviousValue := idFromObject != 0
	id, err = box.IdForPut(idFromObject)
	if err != nil {
		return
	}
	box.binding.Flatten(object, box.fbb, id)
	return id, box.finishInternalFbbAndPutAsync(id, checkForPreviousValue)
}

func (box *Box) finishInternalFbbAndPutAsync(id uint64, checkForPreviousObject bool) (err error) {
	fbb := box.fbb
	fbb.Finish(fbb.EndObject())
	bytes := fbb.FinishedBytes()

	cCheckPrevious := 0
	if checkForPreviousObject {
		cCheckPrevious = 1
	}
	rc := C.ob_box_put_async(box.box, C.uint64_t(id), unsafe.Pointer(&bytes[0]), C.size_t(len(bytes)),
		C.int(cCheckPrevious))
	if rc != 0 {
		err = createError()
	}

	// Reset to have a clear state for the next caller
	fbb.Reset()

	return
}
