// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"net/url"
	"strconv"
)

// HouseRequest contains method which can be used to filter the response.
type HouseRequest interface {
	ParamConverter

	// Limit sets the maximum houses to return.
	Limit(int) HouseRequest

	// Name can be used to filter the returned houses by their name.
	Name(string) HouseRequest

	// Region sets the value for the region parameter.
	// Only houses that belong in the given region are included in the response.
	Region(string) HouseRequest

	// Words sets the value for the words parameter.
	// Only houses that has the given words are included in the response.
	Words(string) HouseRequest

	// HasWords sets the value for the hasWords parameter.
	// Only houses that have words are included in the response.
	HasWords(bool) HouseRequest

	// HasTitles sets the value for the hasTitles parameter.
	// Only houses that have titles are included in the response.
	HasTitles(bool) HouseRequest

	// HasSeats sets the value for the hasSeats parameter.
	// Only houses that have seats are included in the response.
	HasSeats(bool) HouseRequest

	// HasDiedOut sets the value for the hasDiedOut parameter.
	// Only houses that are extinct are included in the response.
	HasDiedOut(bool) HouseRequest

	// HasAncestralWeapons sets the value for the hasAncestralWeapons parameter.
	// Only houses that have ancestral weapons are included in the response.
	HasAncestralWeapons(bool) HouseRequest
}

// NewHouseRequest returns a new HouseRequest which can be used to filter houses.
func NewHouseRequest() HouseRequest {
	h := houseRequest{}
	h.limit = 10
	return h
}

type houseRequest struct {
	request

	name                *string
	region              *string
	words               *string
	hasWords            *bool
	hasTitles           *bool
	hasSeats            *bool
	hasDiedOut          *bool
	hasAncestralWeapons *bool
}

func (request houseRequest) Limit(value int) HouseRequest {
	request.limit = value
	return request
}

func (request houseRequest) Convert() url.Values {
	params := url.Values{}
	params.Add("page", strconv.Itoa(request.limit))

	if request.page != nil {
		params.Set("pageSize", strconv.Itoa(*request.page))
	}
	if request.name != nil {
		params.Set("name", *request.name)
	}
	if request.region != nil {
		params.Set("region", *request.region)
	}
	if request.words != nil {
		params.Set("words", *request.words)
	}
	if request.hasWords != nil {
		params.Set("hasWords", strconv.FormatBool(*request.hasWords))
	}
	if request.hasTitles != nil {
		params.Set("hasTitles", strconv.FormatBool(*request.hasTitles))
	}
	if request.hasSeats != nil {
		params.Set("hasSeats", strconv.FormatBool(*request.hasSeats))
	}
	if request.hasDiedOut != nil {
		params.Set("hasDiedOut", strconv.FormatBool(*request.hasDiedOut))
	}
	if request.hasAncestralWeapons != nil {
		params.Set("hasAncestralWeapons", strconv.FormatBool(*request.hasAncestralWeapons))
	}

	return params
}

func (request houseRequest) Name(value string) HouseRequest {
	request.name = &value
	return request
}

func (request houseRequest) Region(value string) HouseRequest {
	request.region = &value
	return request
}

func (request houseRequest) Words(value string) HouseRequest {
	request.words = &value
	return request
}

func (request houseRequest) HasWords(value bool) HouseRequest {
	request.hasWords = &value
	return request
}

func (request houseRequest) HasTitles(value bool) HouseRequest {
	request.hasTitles = &value
	return request
}

func (request houseRequest) HasSeats(value bool) HouseRequest {
	request.hasSeats = &value
	return request
}

func (request houseRequest) HasDiedOut(value bool) HouseRequest {
	request.hasDiedOut = &value
	return request
}

func (request houseRequest) HasAncestralWeapons(value bool) HouseRequest {
	request.hasAncestralWeapons = &value
	return request
}
