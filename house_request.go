// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"net/url"
	"strconv"
)

type HouseRequest interface {
	ParamConverter

	Name(string) HouseRequest
	Region(string) HouseRequest
	Words(string) HouseRequest
	HasWords(bool) HouseRequest
	HasTitles(bool) HouseRequest
	HasSeats(bool) HouseRequest
	HasDiedOut(bool) HouseRequest
	HasAncestralWeapons(bool) HouseRequest
}

func NewHouseRequest() HouseRequest {
	return houseRequest{}
}

type houseRequest struct {
	name                *string
	region              *string
	words               *string
	hasWords            *bool
	hasTitles           *bool
	hasSeats            *bool
	hasDiedOut          *bool
	hasAncestralWeapons *bool
}

func (request houseRequest) Convert() url.Values {
	params := url.Values{}

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
