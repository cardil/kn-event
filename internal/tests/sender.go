package tests

import (
	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Sender is a mock sender that will record sent events for testing.
type Sender struct {
	Sent []cloudevents.Event
}

// Send will send event to specified target.
func (m *Sender) Send(ce cloudevents.Event) error {
	m.Sent = append(m.Sent, ce)
	return nil
}

// WithSender can be used to wrap invocation with setting a sender
// implementation.
func WithSender(sender event.Sender, body func() error) error {
	old := event.SenderFactory
	defer func() {
		event.SenderFactory = old
	}()
	event.SenderFactory = func(target *event.Target) (event.Sender, error) {
		return sender, nil
	}
	return body()
}
