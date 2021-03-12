package event

import (
	"errors"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var (
	// ErrSenderFactoryUnset will be returned if sender factory isnt set.
	ErrSenderFactoryUnset = errors.New("sender factory is not set")
)

// NewSender will create a sender that can send event to cluster.
func (b Binding) NewSender(target *Target) (Sender, error) {
	sender, err := b.CreateSender(target)
	if err != nil {
		return nil, err
	}
	return &sendLogic{Sender: sender, Properties: target.Properties}, nil
}

type sendLogic struct {
	Sender
	*Properties
}

func (l *sendLogic) Send(ce cloudevents.Event) error {
	err := l.Sender.Send(ce)
	if err == nil {
		l.Log.Infof("Event (ID: %s) have been sent.", ce.ID())
	}
	return err
}
