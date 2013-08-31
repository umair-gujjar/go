// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import "fmt"

const (
    TAG_TXT                 ElemTag = "%s"
    TAG_CLOSED              ElemTag = "<%s%s>%s</%s>"
    TAG_SELF_CLOSED 	    ElemTag = "<%s%s>"
    TAG_SELF_CLOSED_STRICT  ElemTag = "<%s%s/>"
)

type Displayer interface {
    Render() string
}

type ElemTag string

// Map of tags
var Tags map[ElemType]ElemTag

func init() {
    Tags = make(map[ElemType]ElemTag)
    Tags[TYPE_TXT] = TAG_TXT
    Tags[TYPE_CLOSED] = TAG_CLOSED
    Tags[TYPE_SELF_CLOSED] = TAG_SELF_CLOSED
    Tags[TYPE_SELF_CLOSED_STRICT] = TAG_SELF_CLOSED_STRICT
}

type Tagger interface {
	ElementHandler
	Displayer
}

type tag struct {
    ElementHandler
    doctype Doctype
}

// NewTag returns a new html.NewTag to handle
func NewTag(name string) *tag {
    return &tag{NewElement(name, DefaultTagType), HTML5}
}

func (t *tag) SetDoctype(d Doctype) *tag {
	t.doctype = d
	return t
}

func (t tag) GetDoctype() Doctype {
	return t.doctype
}

func (t tag) render() string {
    typeElem := t.GetType()

	initAttribs := func() {
		if t.GetVal() != "" {
			t.SetAttr("value", t.GetVal())
		}
	}

	getAttrsToString := func() string {
	    var attrs string
	    if t.SizeAttrs() > 0 {
	        for _, attr := range t.GetAttrs() {
	            attrs += fmt.Sprintf(`%s="%s" `, attr.GetName(), attr.GetVal())
	        }
	        attrs = " " + attrs[:(len(attrs)-1)]
	    }
	    return attrs
	}

    switch tag := string(Tags[typeElem]); typeElem {
    case TYPE_TXT:
        return fmt.Sprintf(tag, t.GetVal())
    case TYPE_CLOSED:
        return fmt.Sprintf(tag, t.GetName(), getAttrsToString(), t.GetVal(), t.GetName())
    case TYPE_SELF_CLOSED, TYPE_SELF_CLOSED_STRICT:
    	initAttribs()
        return fmt.Sprintf(tag, t.GetName(), getAttrsToString())
    default:
    	initAttribs();
        return fmt.Sprintf(string(TAG_SELF_CLOSED), t.GetName(), getAttrsToString())
    }
}

func Render(tag Tagger) string {
	return tag.Render()
}

func (t tag) Render() string {
	return t.render()
}
