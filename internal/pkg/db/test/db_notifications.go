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
	notifications "github.com/edgexfoundry/edgex-go/internal/support/notifications/interfaces"
	"github.com/edgexfoundry/edgex-go/pkg/models"
	"testing"
)

// TODO proper tests
func TestNotificationsDB(t *testing.T, db notifications.DBClient) {
	var err error

	err = db.Cleanup()
	if err != nil {
		t.Fatal(err)
	}

	err = db.CleanupOld(1)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddNotification(models.Notification{
		Status: models.New,
		Sender: "Foo",
		Slug:   "bar",
		Labels: []string{"abc"}})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddSubscription(models.Subscription{
		Slug:                 "Foo",
		Receiver:             "BarB",
		SubscribedCategories: []models.NotificationsCategory{"cat"},
		SubscribedLabels:     []string{"label"}})
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddTransmission(models.Transmission{Status: models.New})
	if err != nil {
		t.Fatal(err)
	}

	nots, err := db.GetNotifications()
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNewNormalNotifications(1)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNewNotifications(1)
	if err != nil {
		t.Fatal(err)
	}

	nots[0], err = db.GetNotificationById(nots[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNotificationBySender(nots[0].Sender, 1)
	if err != nil {
		t.Fatal(err)
	}

	nots[0], err = db.GetNotificationBySlug(nots[0].Slug)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNotificationsByEnd(nots[0].Created+1, 1)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNotificationsByLabels([]string{nots[0].Labels[0]}, 1)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNotificationsByStart(nots[0].Created-1, 1)
	if err != nil {
		t.Fatal(err)
	}

	nots, err = db.GetNotificationsByStartEnd(nots[0].Created, nots[0].Created, 1)
	if err != nil {
		t.Fatal(err)
	}

	subs, err := db.GetSubscriptions()
	if err != nil {
		t.Fatal(err)
	}

	subs, err = db.GetSubscriptionByCategories([]string{string(subs[0].SubscribedCategories[0])})
	if err != nil {
		t.Fatal(err)
	}

	subs, err = db.GetSubscriptionByCategoriesLabels([]string{string(subs[0].SubscribedCategories[0])}, []string{subs[0].SubscribedLabels[0]})
	if err != nil {
		t.Fatal(err)
	}

	subs[0], err = db.GetSubscriptionById(subs[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	subs, err = db.GetSubscriptionByLabels([]string{subs[0].SubscribedLabels[0]})
	if err != nil {
		t.Fatal(err)
	}

	subs, err = db.GetSubscriptionByReceiver(subs[0].Receiver)
	if err != nil {
		t.Fatal(err)
	}

	subs[0], err = db.GetSubscriptionBySlug(subs[0].Slug)
	if err != nil {
		t.Fatal(err)
	}

	trans, err := db.GetTransmissionsByEnd(0, 1)
	if err != nil {
		t.Fatal(err)
	}

	trans, err = db.GetTransmissionsByNotificationSlug(nots[0].Slug, trans[0].ResendCount+1)
	if err != nil {
		t.Fatal(err)
	}

	trans, err = db.GetTransmissionsByStart(trans[0].Created-1, trans[0].ResendCount+1)
	if err != nil {
		t.Fatal(err)
	}

	trans, err = db.GetTransmissionsByStartEnd(trans[0].Created-1, trans[0].Created+1, trans[0].ResendCount+1)
	if err != nil {
		t.Fatal(err)
	}

	trans, err = db.GetTransmissionsByStatus(trans[0].ResendCount+1, trans[0].Status)
	if err != nil {
		t.Fatal(err)
	}

	err = db.MarkNotificationProcessed(nots[0])
	if err != nil {
		t.Fatal(err)
	}

	nots[0].Content = "foo"
	err = db.UpdateNotification(nots[0])
	if err != nil {
		t.Fatal(err)
	}

	subs[0].Description = "bar"
	err = db.UpdateSubscription(subs[0])
	if err != nil {
		t.Fatal(err)
	}

	trans[0].Notification = nots[0]
	err = db.UpdateTransmission(trans[0])
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteNotificationById(nots[0].ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = db.AddNotification(nots[0])
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteNotificationBySlug(nots[0].Slug)
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteNotificationsOld(10)
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteSubscriptionBySlug(subs[0].Slug)
	if err != nil {
		t.Fatal(err)
	}

	err = db.DeleteTransmission(10, trans[0].Status)
	if err != nil {
		t.Fatal(err)
	}

	err = db.Cleanup()
	if err != nil {
		t.Fatal(err)
	}

	err = db.CleanupOld(1)
	if err != nil {
		t.Fatal(err)
	}

	db.CloseSession()
}
