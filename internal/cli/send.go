package cli

import (
	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Send will send CloudEvent to target.
func Send(ce cloudevents.Event, target *TargetArgs, options *OptionsArgs) error {
	t, err := createTarget(target)
	if err != nil {
		return err
	}
	o := options.WithLogger()
	sender := event.NewSender(t, o)
	return sender.Send(ce)
}
