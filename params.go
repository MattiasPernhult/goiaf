// Copyright 2017 Mattias Pernhult. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package goiaf

import "net/url"

// ParamConverter has only one method, which return a url.Values object.
//
// ParamConverter is implemented by all the request types.
// This way these request types only need to implement this interface
// and they can be passed to the method of the client, which performs the
// api request.
type ParamConverter interface {
	Convert() url.Values
}
