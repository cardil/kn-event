package cli

import (
	"context"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Send will send CloudEvent to target.
func Send(ce *cloudevents.Event, target *TargetArgs, options *Options) error {
	c, err := cloudevents.NewDefaultClient()
	if err != nil {
		return err
	}

	// Set a target.
	ctx := cloudevents.ContextWithTarget(context.Background(), target.URL)

	// Send that Event.
	if err := c.Send(ctx, *ce); cloudevents.IsUndelivered(err) {
		return err
	}

	_, err = fmt.Fprintf(options.OutWriter,
		"Event (ID: %s) have been sent.\n", ce.ID())
	return err
}
