// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

// The struct type attr reprensent an html attribute
type attr struct {
    name string
    val  string
}

// Creating a new attribute
func NewAttr(name, val string) *attr {
    return &attr{name, val}
}

// Set the name of this attribute
func (a *attr) SetName(name string) *attr {
    a.name = name
    return a
}

// Return the name of this attribute
func (a attr) GetName() string {
    return a.name
}

// Set the value of this attribute
func (a *attr) SetVal(val string) *attr {
    a.val = val
    return a
}

// Return the value of this attribute
func (a attr) GetVal() string {
    return a.val
}

// Returns a string attribute
func (a attr) ToString() string {
    return a.GetName() + `="` + a.GetVal() + `"`
}
