// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

type Validater interface {
	Valid(val interface{}) bool
}

type Filter interface {
	Filtered(val interface{}) interface{}
}



