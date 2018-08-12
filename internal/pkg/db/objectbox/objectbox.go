package objectbox

/*
#cgo LDFLAGS: -L ${SRCDIR}/libs -lobjectboxc
#include <stdlib.h>
#include <string.h>
#include "objectbox.h"
*/
import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/google/flatbuffers/go"
	"runtime"
	"strconv"
	"unsafe"
)

//noinspection GoUnusedConst
const (
	PropertyType_Bool       = 1
	PropertyType_Byte       = 2
	PropertyType_Short      = 3
	PropertyType_Char       = 4
	PropertyType_Int        = 5
	PropertyType_Long       = 6
	PropertyType_Float      = 7
	PropertyType_Double     = 8
	PropertyType_String     = 9
	PropertyType_Date       = 10
	PropertyType_Relation   = 11
	PropertyType_ByteVector = 23
)

//noinspection GoUnusedConst
const (
	/// One long property on an entity must be the ID
	PropertyFlags_ID = 1

	/// On languages like Java, a non-primitive type is used (aka wrapper types, allowing null)
	PropertyFlags_NON_PRIMITIVE_TYPE = 2

	/// Unused yet
	PropertyFlags_NOT_NULL = 4
	PropertyFlags_INDEXED  = 8
	PropertyFlags_RESERVED = 16
	/// Unused yet: Unique index
	PropertyFlags_UNIQUE = 32
	/// Unused yet: Use a persisted sequence to enforce ID to rise monotonic (no ID reuse)
	PropertyFlags_ID_MONOTONIC_SEQUENCE = 64
	/// Allow IDs to be assigned by the developer
	PropertyFlags_ID_SELF_ASSIGNABLE = 128
	/// Unused yet
	PropertyFlags_INDEX_PARTIAL_SKIP_NULL = 256
	/// Unused yet, used by References for 1) back-references and 2) to clear references to deleted objects (required for ID reuse)
	PropertyFlags_INDEX_PARTIAL_SKIP_ZERO = 512
	/// Virtual properties may not have a dedicated field in their entity class, e.g. target IDs of to-one relations
	PropertyFlags_VIRTUAL = 1024
	/// Index uses a 32 bit hash instead of the value
	/// (32 bits is shorter on disk, runs well on 32 bit systems, and should be OK even with a few collisions)

	PropertyFlags_INDEX_HASH = 2048
	/// Index uses a 64 bit hash instead of the value
	/// (recommended mostly for 64 bit machines with values longer >200 bytes; small values are faster with a 32 bit hash)
	PropertyFlags_INDEX_HASH64 = 4096
)

//noinspection GoUnusedConst
const (
	DebugFlags_LOG_TRANSACTIONS_READ  = 1
	DebugFlags_LOG_TRANSACTIONS_WRITE = 2
	DebugFlags_LOG_QUERIES            = 4
	DebugFlags_LOG_QUERY_PARAMETERS   = 8
	DebugFlags_LOG_ASYNC_QUEUE        = 16
)

type Model struct {
	model *C.OB_model
	err   error
}

type ObjectBox struct {
	store *C.OB_store
}

type Transaction struct {
	txn *C.OB_txn
}

type Cursor struct {
	cursor *C.OB_cursor
	fbb    *flatbuffers.Builder
}

type Box struct {
	box *C.OB_box
	fbb *flatbuffers.Builder
}

type TableArray struct {
	tableArray *C.OB_table_array
}

type BytesArray struct {
	bytesArray  [][]byte
	cBytesArray *C.OB_bytes_array
}

type TxnFun func(transaction *Transaction) (err error)
type CursorFun func(cursor *Cursor) (err error)

func NewModel() (model *Model, err error) {
	model = &Model{}
	model.model = C.ob_model_create()
	if model.model == nil {
		model = nil
		err = createError()
	}
	return
}

func (model *Model) LastEntityId(id uint32, uid uint64) {
	if model.err != nil {
		return
	}
	C.ob_model_last_entity_id(model.model, C.uint(id), C.ulong(uid))
}

func (model *Model) Entity(name string, id uint32, uid uint64) (err error) {
	if model.err != nil {
		return model.err
	}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	rc := C.ob_model_entity(model.model, cname, C.uint(id), C.ulong(uid))
	if rc != 0 {
		err = createError()
	}
	return
}

func (model *Model) EntityLastPropertyId(id uint32, uid uint64) (err error) {
	if model.err != nil {
		return model.err
	}
	rc := C.ob_model_entity_last_property_id(model.model, C.uint(id), C.ulong(uid))
	if rc != 0 {
		err = createError()
	}
	return
}

func (model *Model) Property(name string, propertyType int, id uint32, uid uint64) (err error) {
	if model.err != nil {
		return model.err
	}
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	rc := C.ob_model_property(model.model, cname, C.OBPropertyType(propertyType), C.uint(id), C.ulong(uid))
	if rc != 0 {
		err = createError()
	}
	return
}

func (model *Model) PropertyFlags(propertyFlags int) (err error) {
	if model.err != nil {
		return model.err
	}
	rc := C.ob_model_property_flags(model.model, C.OBPropertyFlags(propertyFlags))
	if rc != 0 {
		err = createError()
	}
	return
}

func NewObjectBox(model *Model, name string) (objectBox *ObjectBox, err error) {
	fmt.Println("Ignoring name %v", name)
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	objectBox = &ObjectBox{}
	objectBox.store = C.ob_store_open(model.model, nil)
	if objectBox.store == nil {
		objectBox = nil
		err = createError()
	}
	return
}

func (ob *ObjectBox) BeginTxn() (txn *Transaction, err error) {
	var ctxn = C.ob_txn_begin(ob.store)
	if ctxn == nil {
		return nil, createError()
	}
	return &Transaction{ctxn}, nil
}

func (ob *ObjectBox) BeginTxnRead() (txn *Transaction, err error) {
	var ctxn = C.ob_txn_begin_read(ob.store)
	if ctxn == nil {
		return nil, createError()
	}
	return &Transaction{ctxn}, nil
}

func (ob *ObjectBox) RunInTxn(readOnly bool, txnFun TxnFun) (err error) {
	runtime.LockOSThread()
	var txn *Transaction
	if readOnly {
		txn, err = ob.BeginTxnRead()
	} else {
		txn, err = ob.BeginTxn()
	}
	if err != nil {
		return
	}

	gid := getGID()
	//fmt.Println(">>> START TX")
	//os.Stdout.Sync()

	err = txnFun(txn)

	gid2 := getGID()

	if gid != gid2 {
		panic(fmt.Sprintf("GID %v vs. %v", gid, gid2))
	}
	//fmt.Println("<<< END TX")
	//os.Stdout.Sync()

	if !readOnly && err == nil {
		err = txn.Commit()
	}
	err2 := txn.Destroy()
	if err == nil {
		err = err2
	}
	runtime.UnlockOSThread()

	//fmt.Println("<<< END TX Destroy")
	//os.Stdout.Sync()

	return
}

func (ob *ObjectBox) RunWithCursor(entityId uint, readOnly bool, cursorFun CursorFun) (err error) {
	return ob.RunInTxn(readOnly, func(txn *Transaction) (err error) {
		cursor, err := txn.Cursor(entityId)
		if err != nil {
			return
		}
		//fmt.Println(">>> START C")
		//os.Stdout.Sync()

		err = cursorFun(cursor)

		//fmt.Println("<<< END C")
		//os.Stdout.Sync()

		err2 := cursor.Destroy()
		if err == nil {
			err = err2
		}
		return
	})
}

func (ob *ObjectBox) SetDebugFlags(flags uint) (err error) {
	rc := C.ob_store_debug_flags(ob.store, C.uint32_t(flags))
	if rc != 0 {
		err = createError()
	}
	return
}

func (ob *ObjectBox) Box(entitySchemaId uint) (*Box, error) {
	cbox := C.ob_box_create(ob.store, C.uint(entitySchemaId))
	if cbox == nil {
		return nil, createError()
	}
	return &Box{cbox, flatbuffers.NewBuilder(512)}, nil
}

func (ob *ObjectBox) Strict() *ObjectBox {
	if C.ob_store_await_async_completion(ob.store) != 0 {
		fmt.Println(createError())
	}
	return ob
}

func (txn *Transaction) Destroy() (err error) {
	rc := C.ob_txn_destroy(txn.txn)
	txn.txn = nil
	if rc != 0 {
		err = createError()
	}
	return
}

func (txn *Transaction) Abort() (err error) {
	rc := C.ob_txn_abort(txn.txn)
	if rc != 0 {
		err = createError()
	}
	return
}

func (txn *Transaction) Commit() (err error) {
	rc := C.ob_txn_commit(txn.txn)
	if rc != 0 {
		err = createError()
	}
	return
}

func (txn *Transaction) Cursor(entitySchemaId uint) (*Cursor, error) {
	ccursor := C.ob_cursor_create(txn.txn, C.uint(entitySchemaId))
	if ccursor == nil {
		return nil, createError()
	}
	return &Cursor{ccursor, flatbuffers.NewBuilder(512)}, nil
}

func (txn *Transaction) CursorForName(entitySchemaName string) (*Cursor, error) {
	cname := C.CString(entitySchemaName)
	defer C.free(unsafe.Pointer(cname))

	ccursor := C.ob_cursor_create2(txn.txn, cname)
	if ccursor == nil {
		return nil, createError()
	}
	return &Cursor{ccursor, flatbuffers.NewBuilder(512)}, nil
}

func (cursor *Cursor) Destroy() (err error) {
	rc := C.ob_cursor_destroy(cursor.cursor)
	cursor.cursor = nil
	if rc != 0 {
		err = createError()
	}
	return
}

func (cursor *Cursor) Get(id uint64) (bytes []byte, err error) {
	var data *C.void
	var dataSize C.size_t
	dataPtr := unsafe.Pointer(data) // Need ptr to an unsafe ptr here
	rc := C.ob_cursor_get(cursor.cursor, C.uint64_t(id), &dataPtr, &dataSize)
	if rc != 0 {
		if rc != 404 {
			err = createError()
		}
		return
	}
	bytes = C.GoBytes(dataPtr, C.int(dataSize))
	return
}

func (cursor *Cursor) First() (bytes []byte, err error) {
	var data *C.void
	var dataSize C.size_t
	dataPtr := unsafe.Pointer(data) // Need ptr to an unsafe ptr here
	rc := C.ob_cursor_first(cursor.cursor, &dataPtr, &dataSize)
	if rc != 0 {
		if rc != 404 {
			err = createError()
		}
		return
	}
	bytes = C.GoBytes(dataPtr, C.int(dataSize))
	return
}

func (cursor *Cursor) Next() (bytes []byte, err error) {
	var data *C.void
	var dataSize C.size_t
	dataPtr := unsafe.Pointer(data) // Need ptr to an unsafe ptr here
	rc := C.ob_cursor_next(cursor.cursor, &dataPtr, &dataSize)
	if rc != 0 {
		if rc != 404 {
			err = createError()
		}
		return
	}
	bytes = C.GoBytes(dataPtr, C.int(dataSize))
	return
}

func (cursor *Cursor) Count() (count uint64, err error) {
	var cCount C.uint64_t
	rc := C.ob_cursor_count(cursor.cursor, &cCount)
	if rc != 0 {
		err = createError()
		return
	}
	return uint64(cCount), nil
}

func (cursor *Cursor) Put(id uint64, checkForPreviousObject bool) (err error) {
	fbb := cursor.fbb
	fbb.Finish(fbb.EndObject())
	bytes := fbb.FinishedBytes()

	cCheckPrevious := 0
	if checkForPreviousObject {
		cCheckPrevious = 1
	}
	rc := C.ob_cursor_put(cursor.cursor, C.uint64_t(id), unsafe.Pointer(&bytes[0]), C.size_t(len(bytes)),
		C.int(cCheckPrevious))
	if rc != 0 {
		err = createError()
	}

	// Reset to have a clear state for the next caller
	fbb.Reset()

	return
}

func (cursor *Cursor) IdForPut(idCandidate uint64) (id uint64, err error) {
	id = uint64(C.ob_cursor_id_for_put(cursor.cursor, C.uint64_t(idCandidate)))
	if id == 0 {
		err = createError()
	}
	return
}

func (cursor *Cursor) RemoveAll() (err error) {
	rc := C.ob_cursor_remove_all(cursor.cursor)
	if rc != 0 {
		err = createError()
	}
	return
}

func (cursor *Cursor) FindByString(propertyId uint, value string) (bytesArray *BytesArray, err error) {
	cvalue := C.CString(value)
	defer C.free(unsafe.Pointer(cvalue))

	cBytesArray := C.ob_query_by_string(cursor.cursor, C.uint32_t(propertyId), cvalue)
	if cBytesArray == nil {
		err = createError()
		return
	}
	size := int(cBytesArray.size)
	plainBytesArray := make([][]byte, size)
	if size > 0 {
		goBytesArray := (*[1 << 30]C.OB_bytes)(unsafe.Pointer(cBytesArray.bytes))[:size:size]
		for i := 0; i < size; i++ {
			cBytes := goBytesArray[i]
			dataBytes := C.GoBytes(cBytes.data, C.int(cBytes.size))
			plainBytesArray[i] = dataBytes
		}
	}

	return &BytesArray{plainBytesArray, cBytesArray}, nil
}

func (bytesArray *BytesArray) Destroy() {
	cBytesArray := bytesArray.cBytesArray
	if cBytesArray != nil {
		bytesArray.cBytesArray = nil
		C.ob_bytes_array_destroy(cBytesArray)
	}
	bytesArray.bytesArray = nil
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

func (box *Box) PutAsync(id uint64, checkForPreviousObject bool) (err error) {
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

func createError() error {
	msg := C.ob_last_error_message()
	return errors.New(C.GoString(msg))
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
