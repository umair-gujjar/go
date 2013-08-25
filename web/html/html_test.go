// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
    "testing"
)

func ExampleNewTag() {
    t := html.NewTag("input") // Create a new Tag
    t.SetAttr("type", "text")
    t.SetAttr("class", "control")
    t.SetVal("MyValue")
    t.SetType(html.TYPE_TXT)                 // MyValue
    t.SetType(html.TYPE_SELF_CLOSED)         // <input type="text" class="control" value="MyValue"/>
    t.SetType(html.TYPE_SELF_CLOSED_NOSLASH) // <input type="text" class="control" value="MyValue">
    t.SetType(html.TYPE_CLOSED)              // <input type="text" class="control">MyValue</input>
    fmt.Println(t.String())
}
