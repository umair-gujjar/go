// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

import "fmt"

const (
    TAG_TXT                ElemTag = "%s"
    TAG_CLOSED             ElemTag = "<%s%s>%s</%s>"
    TAG_SELF_CLOSED        ElemTag = "<%s%s>"
    TAG_SELF_CLOSED_STRICT ElemTag = "<%s%s/>"
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
    styles map[string]string
    classes []string
}

// NewTag returns a new html.NewTag to handle
func NewTag(name string) *tag {
    return &tag{
        NewElement(name, DefaultTagType),
        HTML5,
        make(map[string]string),
        make([]string, 0),
    }
}

func (t *tag) SetDoctype(d Doctype) *tag {
    t.doctype = d
    return t
}

func (t tag) GetDoctype() Doctype {
    return t.doctype
}

func (t *tag) SetId(name string) *tag {
	t.SetAttr("id", name)
	return t
}

func (t tag) GetId() string {
	return t.GetAttr("id")
}

func (t *tag) SetStyle(propertie, val string) *tag {
	t.styles[propertie] = val
	return t
}

func (t *tag) GetStyle(propertie string) string {
	if t.HasStyle(propertie) {
		return t.GetAttr("style")
	}
	return ""
}

func (t tag) GetStyles() map[string]string {
	return t.styles
}

func (t tag) HasStyle(propertie string) bool {
	_, isOk := t.styles[propertie]
	return isOk
}

func (t *tag) ResetStyle() *tag {
	t.styles = make(map[string]string)
	return t
}

func (t *tag) styleRender() *tag {
	if len(t.GetStyles()) > 0 {
		style := ""
		for propertie, val := range t.GetStyles() {
			style += fmt.Sprintf("%s: %s; ", propertie, val)
		}
		t.SetAttr("style", style[:len(style)-1])
	}
	return t
}

func (t *tag) SetClass(name string) *tag {
	t.classes = append(t.classes, name)
	return t
}

func (t tag) GetClasses() []string {
	return t.classes
}

func (t *tag) ResetClass() *tag {
	t.classes = make([]string, 0)
	return t
}

func (t *tag) classesRender() *tag {
	if len(t.classes) > 0 {
		classes := ""
		for _, class := range t.GetClasses() {
			classes += fmt.Sprintf("%s ", class)
		}
		t.SetAttr("class", classes[:len(classes)-1])
	}
	return t
}

func (t tag) render() string {
    typeElem := t.GetType()

	t.styleRender().classesRender()

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
        initAttribs()
        return fmt.Sprintf(string(TAG_SELF_CLOSED), t.GetName(), getAttrsToString())
    }
}

func Render(tag Tagger) string {
    return tag.Render()
}

func (t tag) Render() string {
    return t.render()
}
