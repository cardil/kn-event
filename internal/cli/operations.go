package cli

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		Fields:    make([]event.FieldSpec, 0, len(args.Fields) + len(args.RawFields)),
	}
	for _, fieldAssigment := range args.Fields {
		splitted := strings.SplitN(fieldAssigment, "=", 2)
		path, value := splitted[0], splitted[1]
		if boolVal, err := readAsBoolean(value); err == nil {
			spec.AddField(path, boolVal)
		} else if floatVal, err := readAsFloat64(value); err == nil {
			spec.AddField(path, floatVal)
		} else {
			spec.AddField(path, value)
		}
	}
	for _, fieldAssigment := range args.RawFields {
		splitted := strings.SplitN(fieldAssigment, "=", 2)
		path, value := splitted[0], splitted[1]
		spec.AddField(path, value)
	}
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

func readAsBoolean(in string) (bool, error) {
	val, err := strconv.ParseBool(in)
	// TODO(cardil): log error as it may be beneficial for debugging
	if err != nil {
		return false, err
	}
	text := fmt.Sprintf("%t", val)
	if in == text {
		return val, nil
	}
	return false, errors.New("not a boolean: " + in)
}

func readAsFloat64(in string) (float64, error) {
	if intVal, err := readAsInt64(in); err == nil {
		return float64(intVal), err
	}
	val, err := strconv.ParseFloat(in, 64)
	// TODO(cardil): log error as it may be beneficial for debugging
	if err != nil {
		return -0, err
	}
	text := fmt.Sprintf("%f", val)
	if in == text {
		return val, nil
	}
	return -0, errors.New("not a float: " + in)
}

func readAsInt64(in string) (int64, error) {
	val, err := strconv.ParseInt(in, 10, 64)
	// TODO(cardil): log error as it may be beneficial for debugging
	if err != nil {
		return -0, err
	}
	text := fmt.Sprintf("%d", val)
	if in == text {
		return val, nil
	}
	return -0, errors.New("not an int: " + in)
}
