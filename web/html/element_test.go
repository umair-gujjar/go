// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
    "reflect"
    "testing"
)

func TestNewElement(t *testing.T) {
    e := newElement("div", TYPE_TXT)

    t.Errorf("%s", reflect.TypeOf(e).String())

    if e.GetName() == "div" {
        t.Errorf("")
    }
}
