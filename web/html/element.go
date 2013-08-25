// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import (
    "fmt"
    "strings"
)

type ElemTag string
type ElemType uint16

const (
    TYPE_TXT                 ElemType = 1 << iota
    TYPE_CLOSED              ElemType = 1 << iota
    TYPE_SELF_CLOSED         ElemType = 1 << iota
    TYPE_SELF_CLOSED_NOSLASH ElemType = 1 << iota
    TAG_TXT                  ElemTag  = "%s"
    TAG_CLOSED               ElemTag  = "<%s%s>%s</%s>"
    TAG_SELF_CLOSED          ElemTag  = "<%s%s/>"
    TAG_SELF_CLOSED_NOSLASH  ElemTag  = "<%s%s>"
)

var tags map[ElemType]ElemTag

type Element interface {
    SetName(name string) *element
    GetName() string
    SetAttr(name string, val string) *element
    GetAttr(name string) string
    GetAttrs() []*attr
    HasAttr(name string) bool
    DelAttr(name string) *element
    ClearAttrs() *element
    SizeAttrs() int
    SetVal(val string) *element
    GetVal() string
    stringer
}

type element struct {
    name    string
    attrs   map[string]*attr
    val     string
    tagType ElemType
}

func init() {
    tags = make(map[ElemType]ElemTag)
    tags[TYPE_TXT] = TAG_TXT
    tags[TYPE_CLOSED] = TAG_CLOSED
    tags[TYPE_SELF_CLOSED] = TAG_SELF_CLOSED
    tags[TYPE_SELF_CLOSED_NOSLASH] = TAG_SELF_CLOSED_NOSLASH
}

// Creating a new element
func newElement(name string, tagType ElemType) *element {
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
func (e *element) setAttr(name string, val string) *element {
    attr := NewAttr(name, val)
    e.attrs[attr.GetName()] = attr
    return e
}

// Set an attribute to the current element
func (e *element) SetAttr(name string, val string) *element {
    return e.setAttr(name, val)
}

// Return the value of an attribute by its key
func (e element) GetAttr(name string) string {
    return e.attrs[name].GetVal()
}

// Return a map of all attributes of this item
func (e element) GetAttrs() []*attr {
    attrs := make([]*attr, 0)
    for _, attr := range e.attrs {
        attrs = append(attrs, attr)
    }
    return attrs
}

// Return all attributes to string
func (e element) getAttrsToString() string {
    var attrs string
    if e.SizeAttrs() > 0 {
        for _, attr := range e.GetAttrs() {
            attrs += attr.String() + " "
        }
        attrs = " " + strings.TrimRight(attrs, " ")
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

func (e *element) ClearAttrs() *element {
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
func (e *element) SetTagType(tagType ElemType) *element {
    e.tagType = tagType
    return e
}

// Return the tag type of this element
func (e element) GetTagType() ElemType {
    return e.tagType
}

// Return this element to string
func (e *element) String() string {
    attrs := e.getAttrsToString()
    typeElem := e.GetTagType()

    initValAttr := func() {
        if e.GetVal() != "" {
            e.SetAttr("value", e.GetVal())
        }
    }

    switch tag := string(tags[typeElem]); typeElem {
    case TYPE_TXT:
        return fmt.Sprintf(tag, e.GetVal())
    case TYPE_CLOSED:
        return fmt.Sprintf(tag, e.GetName(), attrs, e.GetVal(), e.GetName())
    case TYPE_SELF_CLOSED, TYPE_SELF_CLOSED_NOSLASH:
        initValAttr()
        return fmt.Sprintf(tag, e.GetName(), attrs)
    default:
        initValAttr()
        return fmt.Sprintf(string(TAG_SELF_CLOSED), e.GetName(), attrs)
    }
}
