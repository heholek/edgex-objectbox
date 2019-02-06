//
// Copyright (c) 2018 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package test

import (
	scheduler "github.com/edgexfoundry/edgex-go/internal/support/scheduler/interfaces"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"testing"
)

// TODO proper tests
func TestSchedulerDB(t *testing.T, db scheduler.DBClient) {
	var err error

	_, err = db.ScrubAllIntervals()
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.ScrubAllIntervalActions()
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddInterval(models.Interval{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddIntervalAction(models.IntervalAction{})
	if err != nil {
		t.Fatal(err)
	}

	intervals, err := db.Intervals()
	if err != nil {
		t.Fatal(err)
	}

	intervals, err = db.IntervalsWithLimit(1)
	if err != nil {
		t.Fatal(err)
	}

	actions, err := db.IntervalActions()
	if err != nil {
		t.Fatal(err)
	}

	actions, err = db.IntervalActionsWithLimit(1)
	if err != nil {
		t.Fatal(err)
	}

	intervals[0].Name = "foo"
	err = db.UpdateInterval(intervals[0])
	if err != nil {
		t.Fatal(err)
	}

	actions[0].Name = "bar"
	err = db.UpdateIntervalAction(actions[0])
	if err != nil {
		t.Fatal(err)
	}

	intervals[0], err = db.IntervalById(intervals[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	actions[0], err = db.IntervalActionById(actions[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	intervals[0], err = db.IntervalByName(intervals[0].Name)
	if err != nil {
		t.Fatal(err)
	}

	actions[0], err = db.IntervalActionByName(actions[0].Name)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.IntervalActionsByTarget("")
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.IntervalActionsByIntervalName("")
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteIntervalById(intervals[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteIntervalActionById(actions[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	db.CloseSession()
}
