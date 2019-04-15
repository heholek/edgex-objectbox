// Code generated by ObjectBox; DO NOT EDIT.

package obx

import (
	"github.com/objectbox/objectbox-go/objectbox"
)

// ObjectBoxModel declares and builds the model from all the entities in the package.
// It is usually used when setting-up ObjectBox as an argument to the Builder.Model() function.
func ObjectBoxModel() *objectbox.Model {
	model := objectbox.NewModel()
	model.GeneratorVersion(2)

	model.RegisterBinding(AddressableBinding)
	model.RegisterBinding(CommandBinding)
	model.RegisterBinding(DeviceServiceBinding)
	model.RegisterBinding(ReadingBinding)
	model.RegisterBinding(DeviceProfileBinding)
	model.RegisterBinding(DeviceBinding)
	model.RegisterBinding(DeviceReportBinding)
	model.RegisterBinding(EventBinding)
	model.RegisterBinding(IntervalActionBinding)
	model.RegisterBinding(IntervalBinding)
	model.RegisterBinding(NotificationBinding)
	model.RegisterBinding(ProvisionWatcherBinding)
	model.RegisterBinding(RegistrationBinding)
	model.RegisterBinding(SubscriptionBinding)
	model.RegisterBinding(TransmissionBinding)
	model.RegisterBinding(ValueDescriptorBinding)
	model.LastEntityId(16, 2903250960703756959)
	model.LastIndexId(19, 8558604861306073232)
	model.LastRelationId(2, 6583600503460504451)

	return model
}
