// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func (c *client) get(endpoint string, converter ParamConverter, data interface{}) error {
	if converter != nil {
		endpoint = fmt.Sprintf("%s?%s", endpoint, converter.Convert().Encode())
	}

	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}
	req.Close = true

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return ErrResourceNotFound
	}

	if t, ok := data.(linker); ok {
		t.Link(c.getLinks(resp.Header.Get("link")))
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, data)
}

func (c *client) getLinks(linkHeader string) map[string]string {
	result := map[string]string{}
	if linkHeader == "" {
		return result
	}

	links := strings.Split(linkHeader, ",")
	for _, link := range links {
		link = strings.TrimSpace(link)
		linkPair := strings.Split(link, ";")

		urlStr := strings.TrimSpace(linkPair[0])
		rel := strings.Replace(strings.TrimSpace(linkPair[1]), "rel=", "", 1)

		urlStr = urlStr[1 : len(urlStr)-1]
		rel = rel[1 : len(rel)-1]

		result[rel] = urlStr
	}

	return result
}

func getQueryFromURL(urlStr string) (url.Values, error) {
	if urlStr == "" {
		return nil, errors.New("REPLACE THIS LATER, NO RELATION")
	}

	u, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, err
	}

	return u.Query(), nil
}

func getPageInfo(query url.Values) (int, int, error) {
	pageStr, pageSizeStr := query.Get("page"), query.Get("pageSize")
	if pageStr == "" || pageSizeStr == "" {
		return 0, 0, errors.New("PAGE SIZE AND PAGE ARE MISSING")
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		return 0, 0, err
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		return 0, 0, err
	}

	return page, pageSize, nil
}
