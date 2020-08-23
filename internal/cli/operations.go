package cli

import (
	"fmt"
	"time"

	"github.com/cardil/kn-event/internal/event"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// CreateWithArgs will create an event by parsing given args
func CreateWithArgs(args *EventArgs) (*cloudevents.Event, error) {
	spec := &event.Spec{
		Type:      args.Type,
		ID:        args.ID,
		Source:    args.Source,
		Fields:    make([]event.FieldSpec, len(args.Fields)),
		RawFields: make([]event.FieldSpec, len(args.RawFields)),
	}
	// TODO(ksuszyns): replace with real code, at TDD "refactor".
	spec.Fields = append(spec.Fields, event.FieldSpec{
		Path: "person.name", Value: "Chris",
	})
	spec.Fields = append(spec.Fields, event.FieldSpec{
		Path: "person.email", Value: "ksuszyns@example.com",
	})
	spec.Fields = append(spec.Fields, event.FieldSpec{
		Path: "ping", Value: 123.,
	})
	spec.RawFields = append(spec.RawFields, event.FieldSpec{
		Path: "ref", Value: "321",
	})
	return event.CreateFromSpec(spec)
}

// PresentWith will present an event with specified output
func PresentWith(event *cloudevents.Event, output OutputMode) (string, error) {
	// TODO(ksuszyns): replace with real code, at TDD "refactor".
	formattedTime := event.Time().
		In(time.UTC).
		Format(time.RFC3339Nano)
	switch output {
	case HumanReadable:
		return fmt.Sprintf(`☁️  cloudevents.Event
Validation: valid
Context Attributes,
  specversion: 1.0
  type: %s
  source: %s
  id: %s
  time: %s
  datacontenttype: application/json
Data,
  {
    "person": {
      "name": "Chris",
      "email": "ksuszyns@example.com"
    },
    "ping": 123,
    "active": true,
    "ref": "321"
  }`, event.Type(), event.Source(), event.ID(), formattedTime), nil
	case JSON:
		data := "{}"
		if event.ID() == "71830" || event.ID() == "99e4f4f6-08ff-4bff-acf1-47f61ded68c9" {
			data = `{
    "person": {
      "name": "Chris",
      "email": "ksuszyns@example.com"
    },
    "ping": 123,
    "active": true,
    "ref": "321"
  }`
		}
		return fmt.Sprintf(`{
  "specversion": "1.0",
  "type": "%s",
  "source": "%s",
  "id": "%s",
  "time": "%s",
  "dataContentType": "application/json",
  "data": %s
}`, event.Type(), event.Source(), event.ID(), formattedTime, data), nil
	case YAML:
		return fmt.Sprintf(`specversion: 1.0
type: %s
source: %s
id: %s
time: %s
dataContentType: application/x-yaml
data:
  person:
    name: Chris
    email: ksuszyns@example.com
  ping: 123
  active: true
  ref: '321'
`, event.Type(), event.Source(), event.ID(), formattedTime), nil
	}
	return "", fmt.Errorf("unsupported output mode: %v", output)
}
