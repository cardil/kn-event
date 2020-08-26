package cmd

import (
	"bytes"
	"context"
	"net/http/httptest"
	"testing"

	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/assert"
)

func TestSendToAddress(t *testing.T) {
	var ce *cloudevents.Event
	receive := func(ctx context.Context, event cloudevents.Event) {
		ce = &event
	}
	ctx := context.Background()
	protocol, err := cloudevents.NewHTTP()
	assert.NoError(t, err)
	handler, err := cloudevents.NewHTTPReceiveHandler(ctx, protocol, receive)
	assert.NoError(t, err)
	server := httptest.NewServer(handler)
	defer server.Close()
	rootCmd.SetArgs([]string{
		"send",
		"--to-url", server.URL,
		"--id", "654321",
		"--field", "person.name=Chris",
		"--field", "person.email=ksuszyns@example.com",
		"--field", "ping=123",
	})
	buf := bytes.NewBuffer([]byte{})
	rootCmd.SetOut(buf)
	assert.NoError(t, rootCmd.Execute())
	out := string(buf.Bytes())
	assert.Equal(t, "Event (ID: 654321) have been sent.\n", out)
	assert.NotNil(t, ce)
	assert.Equal(t, "654321", ce.ID())
	payload, err := event.UnmarshalData(ce.Data())
	assert.NoError(t, err)
	assert.EqualValues(t, map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Chris",
			"email": "ksuszyns@example.com",
		},
		"ping": 123,
	}, payload)
}
