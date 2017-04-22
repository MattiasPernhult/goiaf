// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

// BookResponse contains the data from the performed request.
//
// BookResponse supports pagination by having four methods: Next(), Prev(), First() and Last().
// These methods will return a BookRequest, already formatted like the previous request you used, except
// the request will return a different result set.
//
// Note that, if a result set is not available these methods will return the ErrNoResultSet error.
type BookResponse struct {
	// Data contains the books from the request.
	Data []Book

	links map[string]string
}

// Next returns a BookRequest, which can be used to retrieve
// the next result set of books.
func (response *BookResponse) Next() (BookRequest, error) {
	return response.getRequestForURL(response.links["next"])
}

// Prev returns a BookRequest, which can be used to retrieve
// the previous result set of books.
func (response *BookResponse) Prev() (BookRequest, error) {
	return response.getRequestForURL(response.links["prev"])
}

// First returns a BookRequest, which can be used to retrieve
// the last result set of books.
func (response *BookResponse) First() (BookRequest, error) {
	return response.getRequestForURL(response.links["first"])
}

// Last returns a BookRequest, which can be used to retrieve
// the first result set of books.
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
