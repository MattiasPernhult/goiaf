// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "strconv"

// HouseResponse contains the data from the performed request.
//
// HouseResponse supports pagination by having four methods: Next(), Prev(), First() and Last().
// These methods will return a HouseRequest, already formatted like the previous request you used, except
// the request will return a different result set.
//
// Note that, if a result set is not available these methods will return the ErrNoResultSet error.
type HouseResponse struct {
	// Data contains the houses from the request.
	Data []House

	links map[string]string
}

// Next returns a HouseRequest, which can be used to retrieve
// the next result set of houses.
func (response *HouseResponse) Next() (HouseRequest, error) {
	return response.getRequestForURL(response.links["next"])
}

// Prev returns a HouseRequest, which can be used to retrieve
// the previous result set of houses.
func (response *HouseResponse) Prev() (HouseRequest, error) {
	return response.getRequestForURL(response.links["prev"])
}

// First returns a HouseRequest, which can be used to retrieve
// the last result set of houses.
func (response *HouseResponse) First() (HouseRequest, error) {
	return response.getRequestForURL(response.links["first"])
}

// Last returns a HouseRequest, which can be used to retrieve
// the first result set of houses.
func (response *HouseResponse) Last() (HouseRequest, error) {
	return response.getRequestForURL(response.links["last"])
}

func (response *HouseResponse) getRequestForURL(urlStr string) (HouseRequest, error) {
	query, err := getQueryFromURL(urlStr)
	if err != nil {
		return nil, err
	}

	page, pageSize, err := getPageInfo(query)
	if err != nil {
		return nil, err
	}

	request := houseRequest{}
	request.limit = page
	request.page = &pageSize

	if value := query.Get("name"); value != "" {
		request.name = &value
	}
	if value := query.Get("region"); value != "" {
		request.region = &value
	}
	if value := query.Get("words"); value != "" {
		request.words = &value
	}
	if value := query.Get("hasWords"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.hasWords = &boolValue
	}
	if value := query.Get("hasTitles"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.hasTitles = &boolValue
	}
	if value := query.Get("hasSeats"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.hasSeats = &boolValue
	}
	if value := query.Get("hasDiedOut"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.hasDiedOut = &boolValue
	}
	if value := query.Get("hasAncestralWeapons"); value != "" {
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return nil, err
		}
		request.hasAncestralWeapons = &boolValue
	}

	return request, nil
}
