// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "time"

// Book represents the book resources that is returned from the api.
type Book struct {
	// The hypermedia URL of this resource.
	URL string

	// The name of this book.
	Name string

	// The International Standard Book Number that uniquely identifies this book.
	// The format used is ISBN-13.
	ISBN string

	// An array of names of the authors that wrote this book.
	Authors []string

	// The number of pages in this book.
	NumberOfPages int

	// The company that published this book.
	Publisher string

	// The country which this book was published in.
	Country string

	// The type of media this book was released in. Possible values are: Hardback,
	// Hardcover, GraphicNovel and Paperback.
	MediaType string

	// The date, in ISO 8601 format, which this book was released.
	Released time.Time

	// An array of Character ids that has been in this book.
	CharacterIds []int

	// An array of Character ids that has had a POV-chapter in this book.
	PovCharacterIds []int
}

type book struct {
	URL           string         `json:"url"`
	Name          string         `json:"name"`
	ISBN          string         `json:"isbn"`
	Authors       []string       `json:"authors"`
	NumberOfPages int            `json:"numberOfPages"`
	Publisher     string         `json:"publisher"`
	Country       string         `json:"country"`
	MediaType     string         `json:"mediaType"`
	Released      DateTime       `json:"released"`
	Characters    urlStringSlice `json:"characters"`
	PovCharacters urlStringSlice `json:"povCharacters"`
}

func (b book) Convert() Book {
	book := Book{
		URL:             b.URL,
		Name:            b.Name,
		ISBN:            b.ISBN,
		Authors:         b.Authors,
		NumberOfPages:   b.NumberOfPages,
		Publisher:       b.Publisher,
		Country:         b.Country,
		MediaType:       b.MediaType,
		Released:        b.Released.Value(),
		CharacterIds:    b.Characters.ids(),
		PovCharacterIds: b.PovCharacters.ids(),
	}

	return book
}

type booksResponse []book

func (booksResponse booksResponse) Convert() []Book {
	books := []Book{}

	for _, bookResponse := range booksResponse {
		books = append(books, bookResponse.Convert())
	}

	return books
}
