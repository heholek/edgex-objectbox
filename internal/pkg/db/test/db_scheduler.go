/*
 * Copyright 2019 ObjectBox Ltd. All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
