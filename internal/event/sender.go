package event

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var (
	// ErrUnsupportedTargetType is an error if user pass unsupported event target
	// type. Only supporting: reachable or addressable.
	ErrUnsupportedTargetType = errors.New("unsupported target type")
	// ErrNotYetImplemented is an error for not yet implemented code.
	ErrNotYetImplemented = errors.New("not yet implemented")
)

// NewSender creates a new Sender.
func NewSender(target *Target, options *Properties) Sender {
	switch target.Type {
	case TargetTypeReachable:
		return &directSender{
			url:        *target.URLVal,
			Properties: options,
		}
	case TargetTypeAddressable:
		return &inClusterSender{}
	}
	panic(fmt.Errorf("%w: %v", ErrUnsupportedTargetType, target.Type))
}

type directSender struct {
	url url.URL
	*Properties
}

func (d *directSender) Send(ce cloudevents.Event) error {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		return err
	}

	// Set a target.
	ctx := cloudevents.ContextWithTarget(context.Background(), d.url.String())

	// Send that Event.
	if err := c.Send(ctx, ce); cloudevents.IsUndelivered(err) {
		return err
	}

	d.Log.Infof("Event (ID: %s) have been sent.", ce.ID())

	return nil
}

type inClusterSender struct {
}

func (i *inClusterSender) Send(ce cloudevents.Event) error {
	return ErrNotYetImplemented
}
