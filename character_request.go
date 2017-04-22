// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"net/url"
	"strconv"
)

// CharacterRequest contains method which can be used to filter the response.
type CharacterRequest interface {
	ParamConverter

	// Name can be used to filter the returned characters by their name.
	Name(string) CharacterRequest

	// Gender filters the result based on gender. Possible values are Female, Male or Unknown.
	Gender(string) CharacterRequest

	// Culture sets the culture for the request. Only characters with the given
	// culture are included in the response.
	Culture(string) CharacterRequest

	// Born can be used to filter characters that was born the given year.
	Born(string) CharacterRequest

	// Died can be used to filter characters that died the given year.
	Died(string) CharacterRequest

	// IsAlive only returns characters that are alive or dead, depending on the
	// given argument value. Does not have a default value, so the response will
	// include characters that are dead and alive.
	IsAlive(bool) CharacterRequest
}

// NewCharacterRequest returns a new CharacterRequest which can be used to filter books.
func NewCharacterRequest() CharacterRequest {
	return characterRequest{}
}

type characterRequest struct {
	name    *string
	gender  *string
	culture *string
	born    *string
	died    *string
	isAlive *bool
}

func (request characterRequest) Convert() url.Values {
	params := url.Values{}

	if request.name != nil {
		params.Set("name", *request.name)
	}
	if request.gender != nil {
		params.Set("gender", *request.gender)
	}
	if request.culture != nil {
		params.Set("culture", *request.culture)
	}
	if request.born != nil {
		params.Set("born", *request.born)
	}
	if request.died != nil {
		params.Set("died", *request.died)
	}
	if request.isAlive != nil {
		params.Set("isAlive", strconv.FormatBool(*request.isAlive))
	}

	return params
}

func (request characterRequest) Name(value string) CharacterRequest {
	request.name = &value
	return request
}

func (request characterRequest) Gender(value string) CharacterRequest {
	request.gender = &value
	return request
}

func (request characterRequest) Culture(value string) CharacterRequest {
	request.culture = &value
	return request
}

func (request characterRequest) Born(value string) CharacterRequest {
	request.born = &value
	return request
}

func (request characterRequest) Died(value string) CharacterRequest {
	request.died = &value
	return request
}

func (request characterRequest) IsAlive(value bool) CharacterRequest {
	request.isAlive = &value
	return request
}
