package objectbox

// implements export-client service contract

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db"
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type schedulerClient struct {
	objectBox *objectbox.ObjectBox

	intervalBox       *obx.IntervalBox
	intervalActionBox *obx.IntervalActionBox

	queries schedulerQueries
}

//region Queries
type schedulerQueries struct {
	interval struct {
		all  intervalQuery
		name intervalQuery
	}
	intervalAction struct {
		all      intervalActionQuery
		interval intervalActionQuery
		name     intervalActionQuery
		target   intervalActionQuery
	}
}

type intervalQuery struct {
	*obx.IntervalQuery
	sync.Mutex
}

type intervalActionQuery struct {
	*obx.IntervalActionQuery
	sync.Mutex
}

//endregion

func newSchedulerClient(objectBox *objectbox.ObjectBox) (*schedulerClient, error) {
	var client = &schedulerClient{objectBox: objectBox}
	var err error

	client.intervalBox = obx.BoxForInterval(objectBox)
	client.intervalActionBox = obx.BoxForIntervalAction(objectBox)

	//region Interval
	if err == nil {
		client.queries.interval.all.IntervalQuery, err = client.intervalBox.QueryOrError()
	}
	if err == nil {
		client.queries.interval.name.IntervalQuery, err =
			client.intervalBox.QueryOrError(obx.Interval_.Name.Equals("", true))
	}
	//endregion

	//region IntervalAction
	if err == nil {
		client.queries.intervalAction.all.IntervalActionQuery, err = client.intervalActionBox.QueryOrError()
	}
	if err == nil {
		client.queries.intervalAction.interval.IntervalActionQuery, err =
			client.intervalActionBox.QueryOrError(obx.IntervalAction_.Interval.Equals("", true))
	}
	if err == nil {
		client.queries.intervalAction.name.IntervalActionQuery, err =
			client.intervalActionBox.QueryOrError(obx.IntervalAction_.Name.Equals("", true))
	}
	if err == nil {
		client.queries.intervalAction.target.IntervalActionQuery, err =
			client.intervalActionBox.QueryOrError(obx.IntervalAction_.Target.Equals("", true))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *schedulerClient) Intervals() ([]contract.Interval, error) {
	return client.intervalBox.GetAll()
}

func (client *schedulerClient) IntervalsWithLimit(limit int) ([]contract.Interval, error) {
	var query = &client.queries.interval.all

	query.Lock()
	defer query.Unlock()

	return query.Limit(uint64(limit)).Find()
}

func (client *schedulerClient) IntervalByName(name string) (contract.Interval, error) {
	var query = &client.queries.interval.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.Interval_.Name, name); err != nil {
		return contract.Interval{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.Interval{}, err
	} else if len(list) == 0 {
		return contract.Interval{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *schedulerClient) IntervalById(id string) (contract.Interval, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.Interval{}, err
	} else if object, err := client.intervalBox.Get(id); err != nil {
		return contract.Interval{}, err
	} else if object == nil {
		return contract.Interval{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *schedulerClient) AddInterval(interval contract.Interval) (string, error) {
	// NOTE this is done instead of onCreate because there is no reg.BaseObject
	if interval.Created == 0 {
		interval.Created = db.MakeTimestamp()
	}

	id, err := client.intervalBox.Put(&interval)
	return obx.IdToString(id), err
}

func (client *schedulerClient) UpdateInterval(interval contract.Interval) error {
	// NOTE this is done instead of onUpdate because there is no reg.BaseObject
	interval.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(interval.ID); err != nil {
		return err
	} else if exists, err := client.intervalBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.intervalBox.Put(&interval)
	return err
}

func (client *schedulerClient) DeleteIntervalById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.intervalBox.Box.Remove(id)
	}
}

func (client *schedulerClient) IntervalActions() ([]contract.IntervalAction, error) {
	return client.intervalActionBox.GetAll()
}

func (client *schedulerClient) IntervalActionsWithLimit(limit int) ([]contract.IntervalAction, error) {
	var query = &client.queries.intervalAction.all

	query.Lock()
	defer query.Unlock()

	return query.Limit(uint64(limit)).Find()
}

func (client *schedulerClient) IntervalActionsByIntervalName(name string) ([]contract.IntervalAction, error) {
	var query = &client.queries.intervalAction.interval

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.IntervalAction_.Interval, name); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *schedulerClient) IntervalActionsByTarget(name string) ([]contract.IntervalAction, error) {
	var query = &client.queries.intervalAction.target

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.IntervalAction_.Target, name); err != nil {
		return nil, err
	}

	return query.Find()
}

func (client *schedulerClient) IntervalActionById(id string) (contract.IntervalAction, error) {
	if id, err := obx.IdFromString(id); err != nil {
		return contract.IntervalAction{}, err
	} else if object, err := client.intervalActionBox.Get(id); err != nil {
		return contract.IntervalAction{}, err
	} else if object == nil {
		return contract.IntervalAction{}, db.ErrNotFound
	} else {
		return *object, nil
	}
}

func (client *schedulerClient) IntervalActionByName(name string) (contract.IntervalAction, error) {
	var query = &client.queries.intervalAction.name

	query.Lock()
	defer query.Unlock()

	if err := query.SetStringParams(obx.IntervalAction_.Name, name); err != nil {
		return contract.IntervalAction{}, err
	}

	if list, err := query.Limit(1).Find(); err != nil {
		return contract.IntervalAction{}, err
	} else if len(list) == 0 {
		return contract.IntervalAction{}, db.ErrNotFound
	} else {
		return list[0], nil
	}
}

func (client *schedulerClient) AddIntervalAction(intervalAction contract.IntervalAction) (string, error) {
	// NOTE this is done instead of onCreate because there is no reg.BaseObject
	if intervalAction.Created == 0 {
		intervalAction.Created = db.MakeTimestamp()
	}

	id, err := client.intervalActionBox.Put(&intervalAction)
	return obx.IdToString(id), err
}

func (client *schedulerClient) UpdateIntervalAction(intervalAction contract.IntervalAction) error {
	// NOTE this is done instead of onUpdate because there is no reg.BaseObject
	intervalAction.Modified = db.MakeTimestamp()

	if id, err := obx.IdFromString(intervalAction.ID); err != nil {
		return err
	} else if exists, err := client.intervalActionBox.Contains(id); err != nil {
		return err
	} else if !exists {
		return db.ErrNotFound
	}

	_, err := client.intervalActionBox.Put(&intervalAction)
	return err
}

func (client *schedulerClient) DeleteIntervalActionById(id string) error {
	if id, err := obx.IdFromString(id); err != nil {
		return err
	} else {
		return client.intervalActionBox.Box.Remove(id)
	}
}

func (client *schedulerClient) ScrubAllIntervalActions() (int, error) {
	var query = &client.queries.intervalAction.all

	query.Lock()
	defer query.Unlock()

	if count, err := query.Remove(); err != nil {
		return 0, err
	} else {
		return int(count), nil
	}
}

func (client *schedulerClient) ScrubAllIntervals() (int, error) {
	var query = &client.queries.interval.all

	query.Lock()
	defer query.Unlock()

	if count, err := query.Remove(); err != nil {
		return 0, err
	} else {
		return int(count), nil
	}
}
