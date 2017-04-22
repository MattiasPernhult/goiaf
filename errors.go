// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "errors"

var (
	// ErrResourceNotFound will be used if the response from the api has 404 as HTTP status.
	ErrResourceNotFound = errors.New("Resource not found")

	// ErrNoResult will be used if no result set exists. Example if response is on first page
	// and you call Prev() on the response, this error will be returned, as no previous result
	// set exists.
	ErrNoResultSet = errors.New("This result set does not exist")

	// ErrPaginationInfoMissing will be used if the api is returning an invalid url.
	ErrPaginationInfoMissing = errors.New("Pagination info missing from returned url by api")
)
