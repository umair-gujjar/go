// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

type ElemType uint16

const (
    TYPE_TXT                 ElemType = 1 << iota
    TYPE_CLOSED              ElemType = 1 << iota
    TYPE_SELF_CLOSED         ElemType = 1 << iota
    TYPE_SELF_CLOSED_STRICT  ElemType = 1 << iota
)

// Represents the interface of a html element
type ElementHandler interface {
    SetName(name string) *element
    GetName() string
    SetAttr(name string, val string) *element
    GetAttr(name string) string
    GetAttrs() []*attr
    HasAttr(name string) bool
    DelAttr(name string) *element
    ResetAttrs() *element
    SizeAttrs() int
    SetVal(val string) *element
    GetVal() string
    SetType(elemType ElemType) *element
    GetType() ElemType
}

type element struct {
    name    string
    attrs   map[string]*attr
    val     string
    tagType ElemType
}

// Creating a new element
func NewElement(name string, tagType ElemType) *element {
    return &element{
        name:    name,
        attrs:   make(map[string]*attr),
        tagType: tagType,
    }
}

// Set the name of this element
func (e *element) SetName(name string) *element {
    e.name = name
    return e
}

// Get the name of this element
func (e element) GetName() string {
    return e.name
}

// Set an attribute to the current element
func (e *element) SetAttr(name string, val string) *element {
    attr := NewAttr(name, val)
    e.attrs[attr.GetName()] = attr
    return e
}

// Return the value of an attribute by its key
func (e element) GetAttr(name string) string {
    if e.HasAttr(name) {
        return e.attrs[name].GetVal()
    } else {
        return ""
    }
}

// Return a map of all attributes of this item
func (e element) GetAttrs() []*attr {
    attrs := make([]*attr, 0)
    for _, attr := range e.attrs {
        attrs = append(attrs, attr)
    }
    return attrs
}

// Checks if attribute exists
func (e element) HasAttr(name string) bool {
    if _, ok := e.attrs[name]; ok {
        return true
    }
    return false
}

// Delele an attribute of this element
func (e *element) DelAttr(name string) *element {
    if e.HasAttr(name) {
        delete(e.attrs, name)
    }
    return e
}

// Removes all attributes
func (e *element) ResetAttrs() *element {
    e.attrs = make(map[string]*attr)
    return e
}

// Return the number of attributes of this element
func (e element) SizeAttrs() int {
    return len(e.attrs)
}

// Set the value of this element
func (e *element) SetVal(val string) *element {
    e.val = val
    return e
}

// Return the value of this element
func (e element) GetVal() string {
    return e.val
}

// Set the tag type of this element
func (e *element) SetType(tagType ElemType) *element {
    e.tagType = tagType
    return e
}

// Return the tag type of this element
func (e element) GetType() ElemType {
    return e.tagType
}
