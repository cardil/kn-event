package ics

import (
	"fmt"
	"net/url"

	"github.com/cardil/kn-event/internal/event"
	"github.com/kelseyhightower/envconfig"
)

// SendFromEnv will send an event based on a values stored in environmental
// variables.
func SendFromEnv() error {
	args := &Args{
		Sink: "localhost",
	}
	err := envconfig.Process("K", args)
	if err != nil {
		return fmt.Errorf("20210104:200412: %w", err)
	}
	u, err := url.Parse(args.Sink)
	if err != nil {
		return fmt.Errorf("20210104:200434: %w", err)
	}
	target := &event.Target{
		Type:   event.TargetTypeReachable,
		URLVal: u,
	}
	s, err := event.SenderFactory(target)
	if err != nil {
		return fmt.Errorf("20210104:200451: %w", err)
	}
	ce, err := Decode(args.Event)
	if err != nil {
		return fmt.Errorf("20210104:200521: %w", err)
	}
	return s.Send(*ce)
}
