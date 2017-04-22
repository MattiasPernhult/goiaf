// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

type Request interface {
	Limit(int) Request
	Page(int) Request
}

type request struct {
	limit int
	page  *int
}

func (r request) Limit(value int) Request {
	r.limit = value
	return r
}

func (r request) Page(value int) Request {
	r.page = &value
	return r
}
