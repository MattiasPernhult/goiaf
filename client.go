// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

// TODO: add functionality for cache headers

import (
	"fmt"
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
		// TODO: timeouts for the http client
		httpClient: http.Client{},
	}
}

func (c *client) Books(request BookRequest) (BookResponse, error) {
	response := BookResponse{}
	booksResponse := booksResponse{}
	err := c.get(booksEndpoint, request, &booksResponse)
	if err != nil {
		return response, err
	}

	response.Data = booksResponse.Convert()
	response.links = booksResponse.links

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
	response := CharacterResponse{}
	charactersResponse := charactersResponse{}

	err := c.get(charactersEndpoint, request, &charactersResponse)
	if err != nil {
		return response, err
	}

	response.Data = charactersResponse.Convert()
	response.links = charactersResponse.links

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
	response := HouseResponse{}
	housesResponse := housesResponse{}

	err := c.get(housesEndpoint, request, &housesResponse)
	if err != nil {
		return response, err
	}

	response.Data = housesResponse.Convert()
	response.links = housesResponse.links

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
