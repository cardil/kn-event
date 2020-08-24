package event

import (
	"encoding/json"
	"time"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/wavesoftware/go-ensure"
)

// NewDefault creates a default CloudEvent
func NewDefault() *cloudevents.Event {
	e := cloudevents.NewEvent()
	e.SetType(DefaultType)
	e.SetID(NewID())
	ensure.NoError(e.SetData(cloudevents.ApplicationJSON, map[string]string{}))
	e.SetSource(DefaultSource)
	e.SetTime(time.Now())
	ensure.NoError(e.Validate())
	return &e
}

// CreateFromSpec will create an event by parsing given args
func CreateFromSpec(args *Spec) (*cloudevents.Event, error) {
	e := NewDefault()
	e.SetID(args.ID)
	e.SetSource(args.Source)
	e.SetType(args.Type)
	// TODO(ksuszyns): replace with real code, at TDD "refactor".
	_ = e.SetData(cloudevents.ApplicationJSON, map[string]interface{}{
		"person": map[string]interface{}{
			"name":  "Chris",
			"email": "ksuszyns@example.com",
		},
		"ping":   123.,
		"ref":    "321",
		"active": true,
	})
	return e, nil
}

// UnmarshalData will take bytes and unmarshall it as JSON to map structure
func UnmarshalData(bytes []byte) (map[string]interface{}, error) {
	m := map[string]interface{}{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// AddField will add a field to the spec
func (s *Spec) AddField(path string, val interface{}) {
	s.Fields = append(s.Fields, FieldSpec{
		Path: path, Value: val,
	})
}
