// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "errors"

var (
	// ErrResourceNotFound will be used if the response from the api has 404 as HTTP status.
	ErrResourceNotFound = errors.New("Resource not found")
)
