// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

// TODO: add functionality for cache headers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	baseURL            string = "http://www.anapioficeandfire.com/api"
	booksEndpoint      string = baseURL + "/books"
	charactersEndpoint string = baseURL + "/characters"
	housesEndpoint     string = baseURL + "/houses"
)

// Client interface which reflects the endpoint for the api.
// All endpoints will be exposed with this client.
type Client interface {
	// Returns all books from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the BookRequest interface.
	Books(BookRequest) (BookResponse, error)

	// Returns a specific book based on the given id.
	Book(int) (Book, error)

	// Returns all characters from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the CharacterRequest interface.
	Characters(CharacterRequest) (CharacterResponse, error)

	// Return a specific character based on the given id.
	Character(int) (Character, error)

	// Returns all houses from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the HouseRequest interface.
	Houses(HouseRequest) (HouseResponse, error)

	// Return a specific house based on the given id.
	House(int) (House, error)
}

type client struct {
	httpClient http.Client
}

// NewClient returns a ice and fire client. All endpoints from the api
// are exposed through this client.
func NewClient() Client {
	return &client{
		httpClient: http.Client{
			Timeout: time.Second * 15,
		},
	}
}

func (c *client) Books(request BookRequest) (BookResponse, error) {
	booksResponse := booksResponse{}
	err := c.get(booksEndpoint, request, &booksResponse)
	if err != nil {
		return BookResponse{}, err
	}

	response := BookResponse{
		Data:  booksResponse.Convert(),
		links: booksResponse.links,
	}

	return response, nil
}

func (c *client) Book(id int) (Book, error) {
	endpoint := fmt.Sprintf("%s/%d", booksEndpoint, id)

	book := book{}
	err := c.get(endpoint, nil, &book)
	if err != nil {
		return Book{}, err
	}

	return book.Convert(), nil
}

func (c *client) Characters(request CharacterRequest) (CharacterResponse, error) {
	charactersResponse := charactersResponse{}

	err := c.get(charactersEndpoint, request, &charactersResponse)
	if err != nil {
		return CharacterResponse{}, err
	}

	response := CharacterResponse{
		Data:  charactersResponse.Convert(),
		links: charactersResponse.links,
	}

	return response, nil
}

func (c *client) Character(id int) (Character, error) {
	endpoint := fmt.Sprintf("%s/%d", charactersEndpoint, id)

	character := character{}
	err := c.get(endpoint, nil, &character)
	if err != nil {
		return Character{}, err
	}

	return character.Convert(), nil
}

func (c *client) Houses(request HouseRequest) (HouseResponse, error) {
	housesResponse := housesResponse{}

	err := c.get(housesEndpoint, request, &housesResponse)
	if err != nil {
		return HouseResponse{}, err
	}

	response := HouseResponse{
		Data:  housesResponse.Convert(),
		links: housesResponse.links,
	}

	return response, nil
}

func (c *client) House(id int) (House, error) {
	endpoint := fmt.Sprintf("%s/%d", housesEndpoint, id)

	house := house{}
	err := c.get(endpoint, nil, &house)
	if err != nil {
		return House{}, err
	}

	return house.Convert(), nil
}

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
		return nil, ErrNoResultSet
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
		return 0, 0, ErrPaginationInfoMissing
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
