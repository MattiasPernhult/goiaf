// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "strconv"

type CharacterResponse struct {
	links map[string]string

	Data []Character
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

func (response *CharacterResponse) Next() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["next"])
}

func (response *CharacterResponse) Prev() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["prev"])
}

func (response *CharacterResponse) First() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["first"])
}

func (response *CharacterResponse) Last() (CharacterRequest, error) {
	return response.getRequestForURL(response.links["last"])
}
