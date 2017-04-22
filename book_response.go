// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

type BookResponse struct {
	links map[string]string

	Data []Book
}

func (response *BookResponse) Next() (BookRequest, error) {
	return response.getRequestForURL(response.links["next"])
}

func (response *BookResponse) Prev() (BookRequest, error) {
	return response.getRequestForURL(response.links["prev"])
}

func (response *BookResponse) First() (BookRequest, error) {
	return response.getRequestForURL(response.links["first"])
}

func (response *BookResponse) Last() (BookRequest, error) {
	return response.getRequestForURL(response.links["last"])
}

func (response *BookResponse) getRequestForURL(urlStr string) (BookRequest, error) {
	query, err := getQueryFromURL(urlStr)
	if err != nil {
		return nil, err
	}

	page, pageSize, err := getPageInfo(query)
	if err != nil {
		return nil, err
	}

	request := bookRequest{}
	request.limit = page
	request.page = &pageSize

	if value := query.Get("name"); value != "" {
		request.name = &value
	}
	if value := query.Get("fromReleaseDate"); value != "" {
		request.fromReleaseDate = &value
	}
	if value := query.Get("toReleaseDate"); value != "" {
		request.toReleaseDate = &value
	}

	return request, nil
}
