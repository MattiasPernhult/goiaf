// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"net/url"
	"time"
)

// BookRequest contains method which can be used to filter the response.
type BookRequest interface {
	ParamConverter

	// Name can be used to filter the books by name.
	Name(string) BookRequest

	// FromReleaseDate can be used to only return books released
	// after the argument date. This method will format the date
	// after the time.RFC3339.
	FromReleaseDate(time.Time) BookRequest

	// FromReleaseDate can be used to only return books released
	// before the argument date. This method will format the date
	// after the time.RFC3339.
	ToReleaseDate(time.Time) BookRequest
}

// NewBookRequest returns a new BookRequest which can be used to filter books.
func NewBookRequest() BookRequest {
	return bookRequest{}
}

type bookRequest struct {
	name            *string
	fromReleaseDate *string
	toReleaseDate   *string
}

func (request bookRequest) Name(value string) BookRequest {
	request.name = &value
	return request
}

func (request bookRequest) FromReleaseDate(value time.Time) BookRequest {
	v := value.Format(time.RFC3339)
	request.fromReleaseDate = &v
	return request
}

func (request bookRequest) ToReleaseDate(value time.Time) BookRequest {
	v := value.Format(time.RFC3339)
	request.toReleaseDate = &v
	return request
}

func (request bookRequest) Convert() url.Values {
	params := url.Values{}

	if request.name != nil {
		params.Set("name", *request.name)
	}
	if request.fromReleaseDate != nil {
		params.Set("fromReleaseDate", *request.fromReleaseDate)
	}
	if request.toReleaseDate != nil {
		params.Set("toReleaseDate", *request.toReleaseDate)
	}

	return params
}
