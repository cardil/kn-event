package configuration

import (
	"github.com/cardil/kn-event/internal/event"
	"github.com/cardil/kn-event/internal/k8s"
	"github.com/cardil/kn-event/internal/sender"
)

func senderBinding() sender.Binding {
	return sender.Binding{
		CreateJobRunner: k8s.CreateJobRunner,
	}
}

func eventsBinding(binding sender.Binding) event.Binding {
	return event.Binding{
		CreateSender: binding.New,
	}
}
