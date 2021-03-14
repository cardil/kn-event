package k8s

import "github.com/cardil/kn-event/internal/event"

// CreateJobRunner will create a JobRunner, or return an error.
func CreateJobRunner(props *event.Properties) (JobRunner, error) {
	return nil, event.ErrNotYetImplemented
}

// CreateAddressResolver will create ReferenceAddressResolver, or return an
// error.
func CreateAddressResolver(props *event.Properties) (ReferenceAddressResolver, error) {
	return nil, event.ErrNotYetImplemented
}
