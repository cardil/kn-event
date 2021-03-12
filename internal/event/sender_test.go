package event_test

import (
	"errors"
	"testing"

	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var errTestError = errors.New("test error")

func TestSendingAnEvent(t *testing.T) {
	tests := []testCase{
		passingCase(),
		unconfiguredBinding(),
		failingSend(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			binding := event.Binding{CreateSender: tt.CreateSender}
			s, err := binding.NewSender(tt.target)
			if err != nil {
				if !errors.Is(err, tt.want) {
					t.Errorf("want: %#v\n got: %#v", tt.want, err)
				}
				return
			}
			got := s.Send(tt.ce)
			if !errors.Is(got, tt.want) {
				t.Errorf("want: %#v\n got: %#v", tt.want, got)
			}
			if tt.bufTest != nil {
				tt.bufTest(t)
			}
		})
	}
}

func passingCase() testCase {
	var buf zaptest.Buffer
	cfg := zap.NewDevelopmentConfig()
	enc := zapcore.NewJSONEncoder(cfg.EncoderConfig)
	log := zap.New(zapcore.NewCore(enc, &buf, cfg.Level))
	ce := cloudevents.NewEvent("1.0")
	ce.SetID("123456")
	return testCase{
		props: &event.Properties{
			Log: log.Sugar(),
		},
		bufTest: func(t *testing.T) {
			assert.Contains(t, buf.String(), "Event (ID: 123456) have been sent.")
		},
		name:         "passing",
		ce:           ce,
		CreateSender: stubSenderFactory,
	}
}

func unconfiguredBinding() testCase {
	return testCase{
		name: "unconfiguredBinding",
		want: event.ErrSenderFactoryUnset,
	}
}

func failingSend() testCase {
	return testCase{
		name: "failingSend",
		want: errTestError,
		CreateSender: func(target *event.Target) (event.Sender, error) {
			return nil, errTestError
		},
	}
}

type stubSender struct{}

func (m *stubSender) Send(_ cloudevents.Event) error {
	return nil
}

var stubSenderFactory = func(*event.Target) (event.Sender, error) {
	return &stubSender{}, nil
}

type testCase struct {
	name    string
	bufTest func(t *testing.T)
	target  *event.Target
	props   *event.Properties
	ce      cloudevents.Event
	want    error
	event.CreateSender
}
