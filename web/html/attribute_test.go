// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import ("testing")

func TestNewAttr(t *testing.T) {
	a := NewAttr("name", "value")
	if a.GetName() != "name" || a.GetVal() != "value" {
		t.Errorf("Error creating an attribute")
	}
}

func TestSetName(t *testing.T) {
	a := NewAttr("name", "value")
	a.SetName("name_test")
	if a.GetName() != "name_test" {
		t.Errorf("Name set")
	} 
}

func TestGetName(t *testing.T) {
	a := NewAttr("name_test", "value")
	if a.GetName() != "name_test" {
		t.Errorf("Name get")
	}
}

func TestSetVal(t *testing.T) {
	a := NewAttr("name", "value")
	a.SetVal("val_test")
	if a.GetVal() != "val_test" {
		t.Errorf("Value set")
	}
}

func TestGetVal(t *testing.T) {
	a := NewAttr("name", "value")
	a.SetVal("val_test")
	if a.GetVal() != "val_test" {
		t.Errorf("No value found")
	}
}

func TestString(t *testing.T) {
	a := NewAttr("name", "value")
	if a.String() != `name="value"` {
		t.Errorf("The string attribute is not valid")
	}
}
