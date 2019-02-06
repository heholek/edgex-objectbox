package objectbox

// implements core-data service contract
// TODO indexes

import (
	"github.com/edgexfoundry/edgex-go/internal/pkg/db/objectbox/obx"
	contract "github.com/edgexfoundry/edgex-go/pkg/models"
	"github.com/objectbox/objectbox-go/objectbox"
	"sync"
)

type notificationsClient struct {
	objectBox *objectbox.ObjectBox

	notificationBox *obx.NotificationBox
	subscriptionBox *obx.SubscriptionBox
	transmissionBox *obx.TransmissionBox

	queries notificationsQueries
}

//region Queries
type notificationsQueries struct {
	notification struct {
		created notificationQuery
		labels  notificationQuery
		sender  notificationQuery
		slug    notificationQuery
	}
	subscription struct {
		categories          subscriptionQuery
		categoriesAndLabels subscriptionQuery
		labels              subscriptionQuery
		receiver            subscriptionQuery
		slug                subscriptionQuery
	}
	transmission struct {
		created transmissionQuery
		status  transmissionQuery
	}
}

type notificationQuery struct {
	*obx.NotificationQuery
	sync.Mutex
}

type subscriptionQuery struct {
	*obx.SubscriptionQuery
	sync.Mutex
}

type transmissionQuery struct {
	*obx.TransmissionQuery
	sync.Mutex
}

//endregion

func newNotificationsClient(objectBox *objectbox.ObjectBox) (*notificationsClient, error) {
	var client = &notificationsClient{objectBox: objectBox}
	var err error

	client.notificationBox = obx.BoxForNotification(client.objectBox)
	client.subscriptionBox = obx.BoxForSubscription(client.objectBox)
	client.transmissionBox = obx.BoxForTransmission(client.objectBox)

	//region Notification
	if err == nil {
		client.queries.notification.created.NotificationQuery, err =
			client.notificationBox.QueryOrError(obx.Notification_.Created.Equals(0))
	}

	if err == nil {
		client.queries.notification.labels.NotificationQuery, err =
			client.notificationBox.QueryOrError(obx.Notification_.Labels.Contains("", true))
	}

	if err == nil {
		client.queries.notification.sender.NotificationQuery, err =
			client.notificationBox.QueryOrError(obx.Notification_.Sender.Equals("", true))
	}

	if err == nil {
		client.queries.notification.slug.NotificationQuery, err =
			client.notificationBox.QueryOrError(obx.Notification_.Slug.Equals("", true))
	}
	//endregion

	//region Subscription
	//if err == nil {
	//	client.queries.subscription.categories.SubscriptionQuery, err =
	//		client.subscriptionBox.QueryOrError(obx.Subscription_.SubscribedCategories.Contains("", true))
	//}
	//
	//if err == nil {
	//	client.queries.subscription.categoriesAndLabels.SubscriptionQuery, err =
	//		client.subscriptionBox.QueryOrError(obx.Reading_.Device.Equals("", true), obx.Reading_.Name.Equals("", true))
	//}

	if err == nil {
		client.queries.subscription.labels.SubscriptionQuery, err =
			client.subscriptionBox.QueryOrError(obx.Subscription_.SubscribedLabels.Contains("", true))
	}

	if err == nil {
		client.queries.subscription.slug.SubscriptionQuery, err =
			client.subscriptionBox.QueryOrError(obx.Subscription_.Slug.Equals("", true))
	}

	if err == nil {
		client.queries.subscription.receiver.SubscriptionQuery, err =
			client.subscriptionBox.QueryOrError(obx.Subscription_.Receiver.Equals("", true))
	}
	//endregion

	//region Transmission
	if err == nil {
		client.queries.transmission.created.TransmissionQuery, err =
			client.transmissionBox.QueryOrError(obx.Transmission_.Created.Equals(0))
	}

	if err == nil {
		client.queries.transmission.status.TransmissionQuery, err =
			client.transmissionBox.QueryOrError(obx.Transmission_.Status.Equals("", true))
	}
	//endregion

	if err == nil {
		return client, nil
	} else {
		return nil, err
	}
}

func (client *notificationsClient) GetNotifications() ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationById(id string) (contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationBySlug(slug string) (contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationBySender(sender string, limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationsByLabels(labels []string, limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationsByStartEnd(start int64, end int64, limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationsByStart(start int64, limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNotificationsByEnd(end int64, limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNewNotifications(limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetNewNormalNotifications(limit int) ([]contract.Notification, error) {
	panic(notImplemented())
}

func (client *notificationsClient) AddNotification(n contract.Notification) (string, error) {
	panic(notImplemented())
}

func (client *notificationsClient) UpdateNotification(n contract.Notification) error {
	panic(notImplemented())
}

func (client *notificationsClient) MarkNotificationProcessed(n contract.Notification) error {
	panic(notImplemented())
}

func (client *notificationsClient) DeleteNotificationById(id string) error { panic(notImplemented()) }

func (client *notificationsClient) DeleteNotificationBySlug(id string) error { panic(notImplemented()) }

func (client *notificationsClient) DeleteNotificationsOld(age int) error { panic(notImplemented()) }

func (client *notificationsClient) GetSubscriptions() ([]contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionById(id string) (contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionBySlug(slug string) (contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionByReceiver(receiver string) ([]contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionByCategories(categories []string) ([]contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionByLabels(labels []string) ([]contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetSubscriptionByCategoriesLabels(categories []string, labels []string) ([]contract.Subscription, error) {
	panic(notImplemented())
}

func (client *notificationsClient) AddSubscription(s contract.Subscription) (string, error) {
	panic(notImplemented())
}

func (client *notificationsClient) UpdateSubscription(s contract.Subscription) error {
	panic(notImplemented())
}

func (client *notificationsClient) DeleteSubscriptionBySlug(id string) error { panic(notImplemented()) }

func (client *notificationsClient) GetTransmissionsByNotificationSlug(slug string, resendLimit int) ([]contract.Transmission, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetTransmissionsByStartEnd(start int64, end int64, resendLimit int) ([]contract.Transmission, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetTransmissionsByStart(start int64, resendLimit int) ([]contract.Transmission, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetTransmissionsByEnd(end int64, resendLimit int) ([]contract.Transmission, error) {
	panic(notImplemented())
}

func (client *notificationsClient) GetTransmissionsByStatus(resendLimit int, status contract.TransmissionStatus) ([]contract.Transmission, error) {
	panic(notImplemented())
}

func (client *notificationsClient) AddTransmission(t contract.Transmission) (string, error) {
	panic(notImplemented())
}

func (client *notificationsClient) UpdateTransmission(t contract.Transmission) error {
	panic(notImplemented())
}

func (client *notificationsClient) DeleteTransmission(age int64, status contract.TransmissionStatus) error {
	panic(notImplemented())
}

func (client *notificationsClient) Cleanup() error { panic(notImplemented()) }

func (client *notificationsClient) CleanupOld(age int) error { panic(notImplemented()) }
