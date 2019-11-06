package objectbox

import "C"
import (
	"github.com/objectbox/edgex-objectbox/internal/pkg/db"
	. "github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

// AsyncView is an interface to AsyncBox, keeping track of inserted/removed IDs
type AsyncView struct {
	box   *Box
	async *AsyncBox

	// ids keep track of the last action with each ID. Using pointer to allow three states
	// 0 = unknown, 1 = inserted, -1 = deleted
	ids   map[uint64]int8
	mutex sync.Mutex
}

func newAsyncView(box *Box) *AsyncView {
	return &AsyncView{
		box:   box,
		async: box.Async(),
		ids:   make(map[uint64]int8, 0),
	}
}

func (view *AsyncView) Clear() {
	view.mutex.Lock()
	view.ids = make(map[uint64]int8, 0)
	view.mutex.Unlock()
}

func (view *AsyncView) markIfNoErr(err error, id uint64, state int8) {
	if err != nil {
		return
	}
	view.mutex.Lock()
	view.ids[id] = state
	view.mutex.Unlock()
}

func (view *AsyncView) Contains(id uint64) (bool, error) {
	view.mutex.Lock()
	defer view.mutex.Unlock()
	if view.ids[id] == 0 {
		return view.box.Contains(id)
	}
	var isPresent = view.ids[id] == 1
	return isPresent, nil
}

func (view *AsyncView) Put(id uint64, object interface{}) (uint64, error) {
	_, err := view.async.Put(object)
	view.markIfNoErr(err, id, 1)
	return id, err
}

func (view *AsyncView) Insert(object interface{}) (uint64, error) {
	id, err := view.async.Insert(object)
	view.markIfNoErr(err, id, 1)
	return id, err
}

func (view *AsyncView) Update(id uint64, object interface{}) error {
	if isPresent, err := view.Contains(id); err != nil {
		return err
	} else if !isPresent {
		return db.ErrNotFound
	}
	return view.async.Update(object)
}

func (view *AsyncView) RemoveId(id uint64) error {
	if isPresent, err := view.Contains(id); err != nil {
		return err
	} else if !isPresent {
		return db.ErrNotFound
	}
	err := view.async.RemoveId(id)
	view.markIfNoErr(err, id, -1)
	return err
}
