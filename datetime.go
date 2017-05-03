// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

var unmarshalDTLayout = "2006-01-02T15:04:05"

// DateTime is a wrapper around time.Time to be able to do proper format for the value.
type DateTime struct {
	time.Time
}

// UnmarshalJSON makes the DateTime type implement the json.Unmarshaller interface.
// It will parse the date according to the ISO 8601 standard.
func (dt *DateTime) UnmarshalJSON(data []byte) error {
	var t string
	err := json.Unmarshal(data, &t)
	if err != nil {
		return errors.New("Invalid Datetime")
	}

	rawTime, err := time.Parse(unmarshalDTLayout, t)
	if err != nil {
		return err
	}

	*dt = DateTime{rawTime}
	return nil
}

// Value returns the underlying wrapped time object
func (dt DateTime) Value() time.Time {
	return dt.Time
}
