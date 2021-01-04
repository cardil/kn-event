package sender

import (
	"errors"
)

var (
	// ErrUnsupportedTargetType is an error if user pass unsupported event target
	// type. Only supporting: reachable or addressable.
	ErrUnsupportedTargetType = errors.New("unsupported target type")

	// ErrCouldntBeSent is an error that will be return in case event that suppose
	// to be sent, couldn't be, for whatever technical reason.
	ErrCouldntBeSent = errors.New("event couldn't be sent")
)
