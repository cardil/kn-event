package k8s

import (
	"context"

	"github.com/cardil/kn-event/internal/event"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// CreateKubeClient creates kubernetes.Interface.
func CreateKubeClient(_ *event.Properties) (Clients, error) {
	return nil, event.ErrNotYetImplemented
}

// Clients holds available Kubernetes clients.
type Clients interface {
	Typed() kubernetes.Interface
	Dynamic() dynamic.Interface
	Context() context.Context
}
