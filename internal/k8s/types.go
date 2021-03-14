package k8s

import (
	"errors"
	"net/url"

	batchv1 "k8s.io/api/batch/v1"
	"knative.dev/pkg/tracker"
)

// ErrInvalidReference is returned if given reference is invalid.
var ErrInvalidReference = errors.New("reference is invalid")

// JobRunner will launch a Job and monitor it for completion.
type JobRunner interface {
	Run(batchv1.Job) error
}

// ReferenceAddressResolver will resolve the tracker.Reference to an url.URL, or
// return an error.
type ReferenceAddressResolver interface {
	ResolveAddress(ref *tracker.Reference) (*url.URL, error)
}
