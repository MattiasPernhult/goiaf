// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

type Character struct {
	// The hypermedia URL of this resource.
	URL string

	// The name of this character.
	Name string

	// The gender of this character. Possible values are:
	// Female, Male and Unknown.
	Gender string

	// The culture that this character belongs to.
	Culture string

	// The year that this person was born.
	Born string

	// The year that this person died.
	Died string

	// The titles that this character holds.
	Titles []string

	// The aliases that this character goes by.
	Aliases []string

	// The character id of this character's father.
	FatherId int

	// The character id of this character's mother.
	MotherId int

	// The character id of this character's spouse.
	SpouseId int

	// An array of House ids that this character is loyal to.
	AllegianceIds []int

	// An array of Book ids that this character has been in.
	BookIds []int

	// An array of Book ids that this character has had a POV-chapter in.
	PovBookIds []int

	// An array of names of the seasons of Game of Thrones that this
	// character has been in.
	TvSeries []string

	// An array of actor names that has played this character in
	// the TV show Game Of Thrones.
	PlayedBy []string
}

type CharacterResponse struct {
	// TODO: add links functionality to next, prev, first, last
	Data []Character
}

type character struct {
	URL         string         `json:"url"`
	Name        string         `json:"name"`
	Gender      string         `json:"gender"`
	Culture     string         `json:"culture"`
	Born        string         `json:"born"`
	Died        string         `json:"died"`
	Titles      []string       `json:"titles"`
	Aliases     []string       `json:"aliases"`
	Father      urlString      `json:"father"`
	Mother      urlString      `json:"mother"`
	Spouse      urlString      `json:"spouse"`
	Allegiances urlStringSlice `json:"allegiances"`
	Books       urlStringSlice `json:"books"`
	PovBooks    urlStringSlice `json:"povBooks"`
	TvSeries    []string       `json:"tvSeries"`
	PlayedBy    []string       `json:"playedBy"`
}

func (c character) Convert() Character {
	character := Character{
		URL:           c.URL,
		Name:          c.Name,
		Gender:        c.Gender,
		Culture:       c.Culture,
		Born:          c.Born,
		Died:          c.Died,
		Titles:        c.Titles,
		Aliases:       c.Aliases,
		FatherId:      c.Father.id(),
		MotherId:      c.Mother.id(),
		SpouseId:      c.Spouse.id(),
		AllegianceIds: c.Allegiances.ids(),
		BookIds:       c.Books.ids(),
		PovBookIds:    c.PovBooks.ids(),
		TvSeries:      c.TvSeries,
		PlayedBy:      c.PlayedBy,
	}

	return character
}

type charactersResponse []character

func (charactersResponse charactersResponse) Convert() []Character {
	characters := []Character{}

	for _, characterResponse := range charactersResponse {
		characters = append(characters, characterResponse.Convert())
	}

	return characters
}
