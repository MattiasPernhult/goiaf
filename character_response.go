// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "strconv"

// CharacterResponse contains the data from the performed request.
//
// CharacterResponse supports pagination by having four methods: Next(), Prev(), First() and Last().
// These methods will return a CharacterRequest, already formatted like the previous request you used, except
// the request will return a different result set.
//
// Note that, if a result set is not available these methods will return the ErrNoResultSet error.
type CharacterResponse struct {
	// Data contains the characters from the request.
	Data []Character

	links map[string]string
}

func (response *CharacterResponse) getRequestForURL(urlStr string) (CharacterRequest, error) {
	query, err := getQueryFromURL(urlStr)
	if err != nil {
		return nil, err
	}

	page, pageSize, err := getPageInfo(query)
	if err != nil {
		return nil, err
	}

	request := characterRequest{}
	request.limit = page
	request.page = &pageSize

	if value := query.Get("name"); value != "" {
		request.name = &value
	}
	if value := query.Get("gender"); value != "" {
		request.gender = &value
	}
	if value := query.Get("culture"); value != "" {
		request.culture = &value
	}
	if value := query.Get("born"); value != "" {
		request.born = &value
	}
	if value := query.Get("died"); value != "" {
		request.died = &value
	}
	if value := query.Get("isAlive"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.isAlive = &boolValue
	}

	return request, nil
}

// Next returns a CharacterRequest, which can be used to retrieve
// the next result set of characters.
func (response *CharacterResponse) Next() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["next"])
}

// Prev returns a CharacterRequest, which can be used to retrieve
// the previous result set of characters.
func (response *CharacterResponse) Prev() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["prev"])
}

// First returns a CharacterRequest, which can be used to retrieve
// the last result set of characters.
func (response *CharacterResponse) First() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["first"])
}

// Last returns a CharacterRequest, which can be used to retrieve
// the first result set of characters.
func (response *CharacterResponse) Last() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["last"])
}
