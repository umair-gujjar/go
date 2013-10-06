// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import "fmt"

type TagType uint16

const (
    TYPE_OPEN               TagType = 1 << iota
    TYPE_CLOSE              TagType = 1 << iota
    TYPE_OPEN_CLOSE         TagType = 1 << iota
    TYPE_SELF_CLOSED_STRICT TagType = 1 << iota
)

// Map of tags
var tagsFormat map[TagType]string

func init() {
    tagsFormat = make(map[TagType]string)
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
    SetType(TagType) *tag
    GetType() TagType
    Displayer
}

type tag struct {
    ElementHandler
    Tagtype TagType
}

// NewTag returns a new html.NewTag to handle
func NewTag(name string, tagtype TagType) *tag {
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
func (t *tag) SetType(tagtype TagType) *tag {
    t.Tagtype = tagtype
    return t
}

// Returns the type of tag.
func (t tag) GetType() TagType {
    return t.Tagtype
}

// Returns the attributes from tagger to string
func getAttrsToString(e ElementHandler) string {
    var attrs string
    if e.LenAttrs() > 0 {
        for _, attr := range e.GetAttrs() {
            attrs += attr.ToString() + " "
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
