package cli

import (
	"io"

	"github.com/thediveo/enumflag"
)

// Options holds a general args for all commands.
type Options struct {
	// KnConfig holds kn configuration file (default: ~/.config/kn/config.yaml)
	KnConfig string

	// Kubeconfig holds kubectl configuration file (default: ~/.kube/config)
	Kubeconfig string

	// Output define type of output commands should be producing.
	Output OutputMode

	// Verbose tells does commands should display additional information about
	// what's happening? Verbose information is printed on stderr.
	Verbose bool

	// LogHTTP tells if kn-event plugin should log HTTP requests it makes
	LogHTTP bool

	OutWriter io.Writer
	ErrWriter io.Writer
}

// EventArgs holds args of event to be created with.
type EventArgs struct {
	Type      string
	ID        string
	Source    string
	Fields    []string
	RawFields []string
}

// TargetArgs holds args specific for even sending.
type TargetArgs struct {
	URL         string
	Addressable string
	Namespace   string
}

// OutputMode is type of output to produce.
type OutputMode enumflag.Flag

// OutputMode enumeration values.
const (
	HumanReadable OutputMode = iota
	JSON
	YAML
)
