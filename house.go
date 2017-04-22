// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "encoding/json"

type House struct {
	// The hypermedia URL of this resource.
	URL string

	// The name of this house.
	Name string

	// The region that this house resides in.
	Region string

	// Text describing the coat of arms of this house.
	CoatOfArms string

	// The words of this house.
	Words string

	// The titles that this house holds.
	Titles []string

	// The seats that this house holds.
	Seats []string

	// The Character id of this house's current lord.
	CurrentLordId int

	// The Character id of this house's heir.
	HeirId int

	// The Houses id that this house answers to.
	OverlordId int

	// The year that this house was founded.
	Founded string

	// The Character resource URL that founded this house.
	FounderId int

	// The year that this house died out.
	DiedOut string

	// An array of names of the noteworthy weapons that this house owns.
	AncestralWeapons []string

	// An array of Houses ids that was founded from this house.
	CadetBranchesIds []int

	// An array of Character ids that are sworn to this house.
	SwornMembersIds []int
}

type house struct {
	URL              string         `json:"url"`
	Name             string         `json:"name"`
	Region           string         `json:"region"`
	CoatOfArms       string         `json:"coatOfArms"`
	Words            string         `json:"words"`
	Titles           []string       `json:"titles"`
	Seats            []string       `json:"seats"`
	CurrentLord      urlString      `json:"currentLord"`
	Heir             urlString      `json:"heir"`
	Overlord         urlString      `json:"overlord"`
	Founded          string         `json:"founded"`
	Founder          urlString      `json:"founder"`
	DiedOut          string         `json:"diedOut"`
	AncestralWeapons []string       `json:"ancestralWeapons"`
	CadetBranches    urlStringSlice `json:"cadetBranches"`
	SwornMembers     urlStringSlice `json:"swornMembers"`
}

func (h house) Convert() House {
	house := House{
		URL:              h.URL,
		Name:             h.Name,
		Region:           h.Region,
		CoatOfArms:       h.CoatOfArms,
		Words:            h.Words,
		Titles:           h.Titles,
		Seats:            h.Seats,
		CurrentLordId:    h.CurrentLord.id(),
		HeirId:           h.Heir.id(),
		OverlordId:       h.Overlord.id(),
		Founded:          h.Founded,
		FounderId:        h.Founder.id(),
		DiedOut:          h.DiedOut,
		AncestralWeapons: h.AncestralWeapons,
		CadetBranchesIds: h.CadetBranches.ids(),
		SwornMembersIds:  h.SwornMembers.ids(),
	}

	return house
}

type housesResponse struct {
	links map[string]string

	Houses []house
}

func (housesResponse *housesResponse) Link(links map[string]string) {
	housesResponse.links = links
}

func (housesResponse *housesResponse) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &housesResponse.Houses)
}

func (housesResponse housesResponse) Convert() []House {
	houses := []House{}

	for _, houseResponse := range housesResponse.Houses {
		houses = append(houses, houseResponse.Convert())
	}

	return houses
}
