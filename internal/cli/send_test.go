package cli_test

import (
	"bytes"
	"fmt"
	"net/url"
	"testing"

	"github.com/cardil/kn-event/internal/cli"
	"github.com/cardil/kn-event/internal/tests"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func TestSendInCli(t *testing.T) {
	outputModes := []cli.OutputMode{cli.HumanReadable, cli.JSON, cli.YAML}
	for _, mode := range outputModes {
		t.Run(fmt.Sprint("OutputMode==", mode), func(t *testing.T) {
			assertWithOutputMode(t, mode)
		})
	}
}

func assertWithOutputMode(t *testing.T, mode cli.OutputMode) {
	var buf bytes.Buffer
	ce, err := tests.WithCloudEventsServer(func(serverURL *url.URL) error {
		ce := cloudevents.NewEvent(cloudevents.VersionV1)
		ce.SetID("543")
		ce.SetType("type")
		ce.SetSource("source")
		target := &cli.TargetArgs{
			URL: serverURL.String(),
		}
		opts := &cli.OptionsArgs{
			Output:    mode,
			OutWriter: &buf,
		}
		return cli.Send(ce, target, opts)
	})
	assert.NoError(t, err)
	assert.NotNil(t, ce)
	assert.Equal(t, "543", ce.ID())
	out := buf.String()

	assert.Contains(t, out, "Event (ID: 543) have been sent.")
}
