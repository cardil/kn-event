package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/cardil/kn-event/internal"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuildSubCommandWithNoOptions(t *testing.T) {
	performTestsOnBuildSubCommand(t, newCmdArgs("build"))
}

func TestBuildSubCommandWithComplexOptions(t *testing.T) {
	performTestsOnBuildSubCommand(
		t, newCmdArgs("build",
			"--type", "org.example.ping",
			"--id", "71830",
			"--source", "/api/v1/ping",
			"--field", "person.name=Chris",
			"--field", "person.email=ksuszyns@example.com",
			"--field", "ping=123",
			"--raw-field", "ref=321",
		),
		func(e *cloudevents.Event) {
			e.SetType("org.example.ping")
			e.SetID("71830")
			e.SetSource("/api/v1/ping")
			assert.NoError(t, e.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
				"person": map[string]interface{}{
					"name":  "Chris",
					"email": "ksuszyns@example.com",
				},
				"ping": 123,
				"ref":  "321",
			}))
		},
	)
}

type eventPreparer func(*cloudevents.Event)

type cmdArgs struct {
	args []string
}

func newCmdArgs(args ...string) cmdArgs {
	return cmdArgs{
		args: args,
	}
}

func performTestsOnBuildSubCommand(t *testing.T, cmd cmdArgs, preparers ...eventPreparer) {
	rootCmd.SetArgs(cmd.args)
	buf := bytes.NewBuffer([]byte{})
	rootCmd.SetOut(buf)
	assert.NoError(t, rootCmd.Execute())
	output := buf.Bytes()
	ec := newEventChecks(t)
	for _, preparer := range preparers {
		preparer(ec.event)
	}
	ec.performEventChecks(output)
}

func (ec eventChecks) performEventChecks(out []byte) {
	actual := cloudevents.NewEvent()
	expected := ec.event
	t := ec.t
	assert.NoError(ec.t, json.Unmarshal(out, &actual))

	assert.NoError(t, actual.Validate())
	assert.Equal(t, expected.Type(), actual.Type())
	assert.Equal(t, expected.DataContentType(), actual.DataContentType())
	assert.Equal(t, ec.unmarshalData(expected.Data()), ec.unmarshalData(actual.Data()))
	assert.Equal(t, expected.Source(), actual.Source())
	delta := 1_000_000.
	assert.InDelta(t, expected.Time().UnixNano(), actual.Time().UnixNano(), delta)
}

func (ec eventChecks) unmarshalData(bytes []byte) map[string]interface{} {
	m := map[string]interface{}{}
	assert.NoError(ec.t, json.Unmarshal(bytes, &m))
	return m
}

func newEventChecks(t *testing.T) eventChecks {
	e := cloudevents.NewEvent()
	e.SetType("dev.knative.cli.plugin.event.generic")
	e.SetID(uuid.New().String())
	assert.NoError(t, e.SetData(cloudevents.ApplicationJSON, map[string]string{}))
	e.SetSource(fmt.Sprintf("%s/%s", internal.PluginName, internal.Version))
	e.SetTime(time.Now())
	assert.NoError(t, e.Validate())
	return eventChecks{
		t:     t,
		event: &e,
	}
}

type eventChecks struct {
	t     *testing.T
	event *cloudevents.Event
}
