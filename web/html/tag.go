// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import "fmt"

type tagType uint16

const (
    TYPE_OPEN               tagType = 1 << iota
    TYPE_CLOSE              tagType = 1 << iota
    TYPE_OPEN_CLOSE         tagType = 1 << iota
    TYPE_SELF_CLOSED_STRICT tagType = 1 << iota
)

// Map of tags
var tagsFormat map[tagType]string

func init() {
    tagsFormat = make(map[tagType]string)
    tagsFormat[TYPE_OPEN] = "<%s%s>"
    tagsFormat[TYPE_CLOSE] = "</%s>"
    tagsFormat[TYPE_OPEN_CLOSE] = tagsFormat[TYPE_OPEN] + "%s" + tagsFormat[TYPE_CLOSE]
    tagsFormat[TYPE_SELF_CLOSED_STRICT] = "<%s%s/>"
}

type Displayer interface {
    Render() string
}

type Tagger interface {
    ElementHandler
    SetId(string) *tag
    GetId() string
    SetType(tagType) *tag
    GetType() tagType
    Displayer
}

type tag struct {
    ElementHandler
    Tagtype tagType
}

// NewTag returns a new html.NewTag to handle
func NewTag(name string, tagtype tagType) *tag {
    return &tag{NewElement(name), tagtype}
}

// Sets the id of this tag
func (t *tag) SetId(id string) *tag {
    t.SetAttr("id", id)
    return t
}

// Returns the id of this tag
func (t tag) GetId() string {
    return t.GetAttr("id")
}

// Sets the type of tag
func (t *tag) SetType(tagtype tagType) *tag {
    t.Tagtype = tagtype
    return t
}

// Returns the type of tag.
func (t tag) GetType() tagType {
    return t.Tagtype
}

// Returns the attributes from tagger to string
func getAttrsToString(t Tagger) string {
    var attrs string
    if t.LenAttrs() > 0 {
        for _, attr := range t.GetAttrs() {
            attrs += fmt.Sprintf(`%s="%s" `, attr.GetName(), attr.GetVal())
        }
        attrs = " " + attrs[:(len(attrs)-1)]
    }
    return attrs
}

// Render the current tag
func (t *tag) render() string {
    tagType := t.GetType()
    tagFormat := tagsFormat[tagType]

    switch t.GetType() {
    case TYPE_OPEN:
        return fmt.Sprintf(tagFormat, t.GetName(), getAttrsToString(t))
    case TYPE_CLOSE:
        return fmt.Sprintf(tagFormat, t.GetName())
    case TYPE_OPEN_CLOSE:
        return fmt.Sprintf(tagFormat, t.GetName(), getAttrsToString(t), t.GetVal(), t.GetName())
    case TYPE_SELF_CLOSED_STRICT:
        if t.GetVal() != "" {
            t.SetAttr("value", t.GetVal())
        }
        return fmt.Sprintf(tagFormat, t.GetName(), getAttrsToString(t))
    default:
        return t.GetName()
    }
}

// Render the current tag
func (t tag) Render() string {
    return t.render()
}

// Rendering of a tag
func Render(tag Tagger) string {
    return tag.Render()
}
