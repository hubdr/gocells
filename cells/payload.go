// Tideland Go Cells - Payload
//
// Copyright (C) 2010-2017 Frank Mueller / Tideland / Oldenburg / Germany
//
// All rights reserved. Use of this source code is governed
// by the new BSD license.

package cells

//--------------------
// IMPORTS
//--------------------

import (
	"encoding/json"
	"fmt"

	"github.com/tideland/golib/errors"
)

//--------------------
// PAYLOAD
//--------------------

// Payload is a write-once/read-multiple container for the
// transport of additional information with events.
type Payload interface {
	fmt.Stringer

	// Bytes returns the raw payload bytes.
	Bytes() []byte

	// Unmarshal parses the JSON-encoded payload bytes and
	// stores the result in the value pointed to by v.
	Unmarshal(v interface{}) error
}

// payload implements the Payload interface.
type payload struct {
	data []byte
}

// NewPayload creates a new payload based on the passed value. In
// case of a byte slice this is taken directly, otherwise it is
// marshalled into JSON.
func NewPayload(v interface{}) (Payload, error) {
	var data []byte
	var err error
	switch tv := v.(type) {
	case []byte:
		data = make([]byte, len(tv))
		copy(data, tv)
	case Payload:
		return tv, nil
	default:
		data, err = json.Marshal(v)
		if err != nil {
			return nil, errors.Annotate(err, ErrMarshal, errorMessages)
		}
	}
	return &payload{
		data: data,
	}, nil
}

// Bytes implements Payload.
func (p *payload) Bytes() []byte {
	data := make([]byte, len(p.data))
	copy(data, p.data)
	return data
}

// Unmarshal implements Payload.
func (p *payload) Unmarshal(v interface{}) error {
	err := json.Unmarshal(p.data, v)
	if err != nil {
		return errors.Annotate(err, ErrUnmarshal, errorMessages)
	}
	return nil
}

// String implements fmt.Stringer.
func (p *payload) String() string {
	return string(p.data)
}

// EOF
