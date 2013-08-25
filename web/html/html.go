// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

var DefaultDoctype Doctype
var DefaultTagType ElemType

func init() {
    DefaultDoctype = HTML5
    DefaultTagType = TYPE_TXT
}

type Stringer interface {
    String() string
}

type tag struct {
    *element
    Doctype
}

// Tag returns a new html.Tag to handle
func NewTag(name string) *tag {
    return &tag{newElement(name, DefaultTagType), HTML5}
}
