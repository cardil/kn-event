package sender

import (
	"github.com/cardil/kn-event/internal/event"
	"github.com/cardil/kn-event/internal/k8s"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type inClusterSender struct {
	addressable *event.AddressableSpec
	jobRunner   k8s.JobRunner
}

func (i *inClusterSender) Send(ce cloudevents.Event) error {
	return event.ErrNotYetImplemented
}
