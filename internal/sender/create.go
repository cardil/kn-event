package sender

import (
	"fmt"

	"github.com/cardil/kn-event/internal/event"
)

// New creates a new Sender.
func (b *Binding) New(target *event.Target) (event.Sender, error) {
	switch target.Type {
	case event.TargetTypeReachable:
		return &directSender{
			url: *target.URLVal,
		}, nil
	case event.TargetTypeAddressable:
		jr, err := b.CreateJobRunner(target.Properties)
		if err != nil {
			return nil, err
		}
		return &inClusterSender{
			addressable: target.AddressableVal,
			jobRunner:   jr,
		}, nil
	}
	return nil, fmt.Errorf("%w: %v", ErrUnsupportedTargetType, target.Type)
}
