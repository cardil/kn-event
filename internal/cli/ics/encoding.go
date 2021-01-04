package ics

import (
	"bytes"
	"compress/zlib"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

// Encode will encode a cloud event to ICS encoding form:
// Base64(zlib(minimal JSON)).
func Encode(ce *cloudevents.Event) (string, error) {
	bb, err := json.Marshal(ce)
	if err != nil {
		return "", fmt.Errorf("20210104:203217: %w", err)
	}
	var b bytes.Buffer
	encoder := base64.NewEncoder(base64.RawURLEncoding, &b)
	w := zlib.NewWriter(encoder)
	_, err = w.Write(bb)
	if err != nil {
		return "", fmt.Errorf("20210104:203302: %w", err)
	}
	err = w.Close()
	if err != nil {
		return "", fmt.Errorf("20210104:203321: %w", err)
	}
	err = encoder.Close()
	if err != nil {
		return "", fmt.Errorf("20210104:203334: %w", err)
	}
	return b.String(), nil
}

// Decode will decode an event from ICS encoding.
func Decode(encoded string) (*cloudevents.Event, error) {
	r := strings.NewReader(encoded)
	decoder := base64.NewDecoder(base64.RawURLEncoding, r)
	reader, err := zlib.NewReader(decoder)
	if err != nil {
		return nil, fmt.Errorf("20210104:221412: %w", err)
	}
	bb, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("20210104:221831: %w", err)
	}
	ce := &cloudevents.Event{}
	err = json.Unmarshal(bb, ce)
	if err != nil {
		return nil, fmt.Errorf("20210104:222046: %w", err)
	}
	return ce, nil
}
