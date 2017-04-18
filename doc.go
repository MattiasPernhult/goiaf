// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package goiaf provides an HTTP wrapper around An API Of Ice And Fire (https://anapioficeandfire.com).


This API includes all the data from the universe of Ice And Fire you've ever wanted, created by https://github.com/joakimskoog

	client := goiaf.NewClient()

	//
	//  Books
	//
	book, err := client.Book(1)
	checkErr(err)
	fmt.Printf("%+v\n", book)

	books, err := client.Books(goiaf.NewBookRequest())
	checkErr(err)
	fmt.Printf("%+v\n", books)

	//
	//  Characters
	//
	characters, err := client.Characters(goiaf.NewCharacterRequest())
	checkErr(err)
	fmt.Printf("%+v\n", characters)

	characters, err = client.Characters(goiaf.NewCharacterRequest().Name("Jon Snow"))
	checkErr(err)
	fmt.Printf("%+v\n", characters)

	character, err := client.Character(583)
	checkErr(err)
	fmt.Printf("%+v\n", character)

	//
	//  Houses
	//
	houses, err := client.Houses(goiaf.NewHouseRequest())
	checkErr(err)
	fmt.Printf("%+v\n", houses)

	houses, err = client.Houses(goiaf.NewHouseRequest().HasDiedOut(false))
	checkErr(err)
	fmt.Printf("%+v\n", houses)

	house, err := client.House(378)
	checkErr(err)
	fmt.Printf("%+v\n", house)
*/
package goiaf
