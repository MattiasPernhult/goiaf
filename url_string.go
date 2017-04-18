// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import (
	"strconv"
	"strings"
)

type urlString string

func (us urlString) id() int {
	strValue := string(us)
	if strValue == "" {
		return -1
	}

	idStr := strValue[strings.LastIndex(strValue, "/")+1 : len(strValue)]

	id, _ := strconv.Atoi(idStr)
	return id
}

type urlStringSlice []urlString

func (uss urlStringSlice) ids() []int {
	ids := []int{}
	for _, us := range uss {
		ids = append(ids, us.id())
	}

	return ids
}
