// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

var DefaultDoctype Doctype
var DefaultTagType TagType

func init() {
    DefaultDoctype = HTML5
    DefaultTagType = TYPE_TXT
}
