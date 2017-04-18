// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

// TODO: add functionality to increase page size
// TODO: add pagination functionality
// TODO: add functionality for cache headers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	baseURL            string = "http://www.anapioficeandfire.com/api"
	booksEndpoint      string = baseURL + "/books"
	charactersEndpoint string = baseURL + "/characters"
	housesEndpoint     string = baseURL + "/houses"
)

type Client interface {
	// Returns all books from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the BookRequest interface.
	Books(BookRequest) ([]Book, error)

	// Returns a specific book based on the id.
	Book(int) (Book, error)

	// Returns all characters from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the CharacterRequest interface.
	Characters(CharacterRequest) (CharacterResponse, error)

	// Return a specific character based on the id.
	Character(int) (Character, error)

	// Returns all houses from the api. There is also the possibility
	// to include filter parameters in your request. The possible filter
	// parameters are exposed by the HouseRequest interface.
	Houses(HouseRequest) (HouseResponse, error)

	// Return a specific house based on the id.
	House(int) (House, error)
}

type client struct {
	httpClient http.Client
}

// NewClient returns a ice and fire client. All endpoints from the api
// are exposed through this client.
func NewClient() Client {
	return &client{
		// TODO: timeouts for the http client
		httpClient: http.Client{},
	}
}

func (c *client) Books(request BookRequest) ([]Book, error) {
	booksResponse := booksResponse{}
	err := c.doRequest(booksEndpoint, request, &booksResponse)
	if err != nil {
		return nil, err
	}

	return booksResponse.Convert(), nil
}

func (c *client) Book(id int) (Book, error) {
	endpoint := fmt.Sprintf("%s/%d", booksEndpoint, id)

	bookResponse := book{}
	err := c.doRequest(endpoint, nil, &bookResponse)
	if err != nil {
		return Book{}, err
	}

	return bookResponse.Convert(), nil
}

func (c *client) Characters(request CharacterRequest) (CharacterResponse, error) {
	response := CharacterResponse{}
	charactersResponse := charactersResponse{}

	err := c.doRequest(charactersEndpoint, request, &charactersResponse)
	if err != nil {
		return response, err
	}

	response.Data = charactersResponse.Convert()

	return response, nil
}

func (c *client) Character(id int) (Character, error) {
	endpoint := fmt.Sprintf("%s/%d", charactersEndpoint, id)

	characterResponse := character{}
	err := c.doRequest(endpoint, nil, &characterResponse)
	if err != nil {
		return Character{}, err
	}

	return characterResponse.Convert(), nil
}

func (c *client) Houses(request HouseRequest) (HouseResponse, error) {
	response := HouseResponse{}
	housesResponse := housesResponse{}

	err := c.doRequest(housesEndpoint, request, &housesResponse)
	if err != nil {
		return response, err
	}

	response.Data = housesResponse.Convert()

	return response, nil
}

func (c *client) House(id int) (House, error) {
	endpoint := fmt.Sprintf("%s/%d", housesEndpoint, id)

	houseResponse := house{}
	err := c.doRequest(endpoint, nil, &houseResponse)
	if err != nil {
		return House{}, err
	}

	return houseResponse.Convert(), nil
}

func (c *client) doRequest(endpoint string, converter ParamConverter, data interface{}) error {
	url := endpoint
	if converter != nil {
		url = fmt.Sprintf("%s?%s", endpoint, converter.Convert().Encode())
	}

	req, err := http.NewRequest("GET", url, nil)
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

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(b, data)
}
