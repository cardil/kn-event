package ics_test

import (
	"testing"
	"time"

	"github.com/cardil/kn-event/internal/cli/ics"
	"github.com/cardil/kn-event/internal/tests"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func TestSendFromEnv(t *testing.T) {
	want := cloudevents.NewEvent()
	want.SetID("123=456=123")
	want.SetTime(time.Now().UTC())
	want.SetType("example")
	want.SetSource("tests://example")
	kevent, err := ics.Encode(want)
	assert.NoError(t, err)
	sender := &tests.Sender{}
	err = tests.WithSender(sender, func() error {
		env := map[string]string{
			"K_SINK":  "http://cosmos.custer.local",
			"K_EVENT": kevent,
		}
		return tests.WithEnviron(env, ics.SendFromEnv)
	})
	assert.NoError(t, err)
	assert.Equal(t, 1, len(sender.Sent))
	got := sender.Sent[0]
	assert.EqualValues(t, want, got)
}
