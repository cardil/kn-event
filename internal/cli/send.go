package cli

import (
	"net/url"

	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Send will send CloudEvent to target.
func Send(ce cloudevents.Event, target *TargetArgs, options *Options) error {
	t, err := createTarget(target)
	if err != nil {
		return err
	}
	o := createOptions(options)
	sender := event.NewSender(t, o)
	return sender.Send(ce)
}

func createTarget(args *TargetArgs) (*event.Target, error) {
	if args.URL != "" {
		u, err := url.Parse(args.URL)
		if err != nil {
			return nil, err
		}
		return &event.Target{
			Type:   event.TargetTypeReachable,
			URLVal: u,
		}, nil
	}
	return nil, event.ErrNotYetImplemented
}
