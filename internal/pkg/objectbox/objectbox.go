package objectbox

/*
#cgo LDFLAGS: -L ${SRCDIR}/libs -lobjectboxc
#include <stdlib.h>
#include <string.h>
#include "objectbox.h"
*/
import "C"

import (
	"errors"
	"fmt"
	"github.com/google/flatbuffers/go"
	"runtime"
	"strconv"
	"strings"
)

const Unavailable = flatbuffers.UOffsetT(0)

//noinspection GoUnusedConst
const (
	DebugFlags_LOG_TRANSACTIONS_READ  = 1
	DebugFlags_LOG_TRANSACTIONS_WRITE = 2
	DebugFlags_LOG_QUERIES            = 4
	DebugFlags_LOG_QUERY_PARAMETERS   = 8
	DebugFlags_LOG_ASYNC_QUEUE        = 16
)

type TypeId uint32

type ObjectBinding interface {
	GetTypeId() TypeId
	GetTypeName() string
	GetId(object interface{}) (id uint64, err error)
	Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64)
	ToObject(bytes []byte) interface{}
	AppendToSlice(slice interface{}, object interface{}) (sliceNew interface{})
}

type ObjectBox struct {
	store          *C.OB_store
	bindingsById   map[TypeId]ObjectBinding
	bindingsByName map[string]ObjectBinding
}

type TableArray struct {
	tableArray *C.OB_table_array
}

type BytesArray struct {
	BytesArray  [][]byte
	cBytesArray *C.OB_bytes_array
}

type TxnFun func(transaction *Transaction) (err error)
type CursorFun func(cursor *Cursor) (err error)

func (ob *ObjectBox) RegisterBinding(binding ObjectBinding) {
	id := binding.GetTypeId()
	name := binding.GetTypeName()
	existingBinding := ob.bindingsById[id]
	if existingBinding != nil {
		panic("Already registered a binding for ID " + strconv.Itoa(int(id)) + ": " + binding.GetTypeName())
	}
	existingBinding = ob.bindingsByName[name]
	if existingBinding != nil {
		panic("Already registered a binding for name " + name + ": ID " + strconv.Itoa(int(binding.GetTypeId())))
	}
	ob.bindingsById[id] = binding
	ob.bindingsByName[name] = binding
}

func (ob *ObjectBox) BeginTxn() (txn *Transaction, err error) {
	var ctxn = C.ob_txn_begin(ob.store)
	if ctxn == nil {
		return nil, createError()
	}
	return &Transaction{ctxn, ob}, nil
}

func (ob *ObjectBox) BeginTxnRead() (txn *Transaction, err error) {
	var ctxn = C.ob_txn_begin_read(ob.store)
	if ctxn == nil {
		return nil, createError()
	}
	return &Transaction{ctxn, ob}, nil
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

	//fmt.Println(">>> START TX")
	//os.Stdout.Sync()

	err = txnFun(txn)

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

func (ob ObjectBox) getBindingById(typeId TypeId) ObjectBinding {
	binding := ob.bindingsById[typeId]
	if binding == nil {
		// Configuration error by the dev, OK to panic
		panic("Configuration error; no binding registered for type ID " + strconv.Itoa(int(typeId)))
	}
	return binding
}

func (ob ObjectBox) getBindingByName(typeName string) ObjectBinding {
	binding := ob.bindingsByName[strings.ToLower(typeName)]
	if binding == nil {
		// Configuration error by the dev, OK to panic
		panic("Configuration error; no binding registered for type name " + typeName)
	}
	return binding
}

func (ob *ObjectBox) RunWithCursor(typeId TypeId, readOnly bool, cursorFun CursorFun) (err error) {
	binding := ob.getBindingById(typeId)
	return ob.RunInTxn(readOnly, func(txn *Transaction) (err error) {
		cursor, err := txn.Cursor(binding)
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

func (ob *ObjectBox) Box(entitySchemaId TypeId) (*Box, error) {
	binding := ob.getBindingById(entitySchemaId)
	cbox := C.ob_box_create(ob.store, C.uint(entitySchemaId))
	if cbox == nil {
		return nil, createError()
	}
	return &Box{cbox, binding, flatbuffers.NewBuilder(512)}, nil
}

func (ob *ObjectBox) Strict() *ObjectBox {
	if C.ob_store_await_async_completion(ob.store) != 0 {
		fmt.Println(createError())
	}
	return ob
}

func (bytesArray *BytesArray) Destroy() {
	cBytesArray := bytesArray.cBytesArray
	if cBytesArray != nil {
		bytesArray.cBytesArray = nil
		C.ob_bytes_array_destroy(cBytesArray)
	}
	bytesArray.BytesArray = nil
}

func createError() error {
	msg := C.ob_last_error_message()
	if msg == nil {
		return errors.New("no error info available; please report")
	} else {
		return errors.New(C.GoString(msg))
	}
}
