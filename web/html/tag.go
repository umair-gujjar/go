// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import "fmt"

type TagType uint16
type ElemTag string

const (
    TYPE_TXT                TagType = 1 << iota
    TYPE_CLOSED             TagType = 1 << iota
    TYPE_SELF_CLOSED        TagType = 1 << iota
    TYPE_SELF_CLOSED_STRICT TagType = 1 << iota
)

const (
    TAG_TXT                ElemTag = "%s"
    TAG_CLOSED             ElemTag = "<%s%s>%s</%s>"
    TAG_SELF_CLOSED        ElemTag = "<%s%s>"
    TAG_SELF_CLOSED_STRICT ElemTag = "<%s%s/>"
)

type Displayer interface {
    Render() string
}

// Map of tags
var Tags map[TagType]ElemTag

func init() {
    Tags = make(map[TagType]ElemTag)
    Tags[TYPE_TXT] = TAG_TXT
    Tags[TYPE_CLOSED] = TAG_CLOSED
    Tags[TYPE_SELF_CLOSED] = TAG_SELF_CLOSED
    Tags[TYPE_SELF_CLOSED_STRICT] = TAG_SELF_CLOSED_STRICT
}

type Tagger interface {
    ElementHandler
    Displayer
    SetId(string) *tag
    GetId() string
    SetType(TagType) *tag
    GetType() TagType
}

type tag struct {
    ElementHandler
    doctype Doctype
    tagType TagType
}

// NewTag returns a new html.NewTag to handle
func NewTag(name string) *tag {
    return &tag{
        NewElement(name),
        HTML5,
        DefaultTagType,
    }
}

// Set the tag type of this element
func (t *tag) SetType(tagType TagType) *tag {
    t.tagType = tagType
    return t
}

// Return the tag type of this element
func (t tag) GetType() TagType {
    return t.tagType
}

// Sets the type of the html document
func (t *tag) SetDoctype(d Doctype) *tag {
    t.doctype = d
    return t
}

// Returns the type of the html document
func (t tag) GetDoctype() Doctype {
    return t.doctype
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

// Returns the attributes from tagger to string
func GetAttrsToString(t Tagger) string {
    var attrs string
    if t.LenAttrs() > 0 {
        for _, attr := range t.GetAttrs() {
            attrs += fmt.Sprintf(`%s="%s" `, attr.GetName(), attr.GetVal())
        }
        attrs = " " + attrs[:(len(attrs)-1)]
    }
    return attrs
}

func (t *tag) render() string {
    tagType := t.GetType()

    initAttribs := func() {
        if t.GetVal() != "" {
            t.SetAttr("value", t.GetVal())
        }
    }

    switch tag := string(Tags[tagType]); tagType {
    case TYPE_TXT:
        return fmt.Sprintf(tag, t.GetVal())
    case TYPE_CLOSED:
        return fmt.Sprintf(tag, t.GetName(), GetAttrsToString(t), t.GetVal(), t.GetName())
    case TYPE_SELF_CLOSED, TYPE_SELF_CLOSED_STRICT:
        initAttribs()
        return fmt.Sprintf(tag, t.GetName(), GetAttrsToString(t))
    default:
        initAttribs()
        return fmt.Sprintf(string(TAG_SELF_CLOSED), t.GetName(), GetAttrsToString(t))
    }
}

func Render(tag Tagger) string {
    return tag.Render()
}

func (t tag) Render() string {
    return t.render()
}
