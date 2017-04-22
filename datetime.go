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

// Date is a wrapper around time.Time.
type DateTime struct {
	time.Time
}

// UnmarshalJSON makes the DateTime type implement the json.Unmarshaller interface.
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

func (dt DateTime) Value() time.Time {
	return dt.Time
}
