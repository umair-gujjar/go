// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

import (
   "validate"
   "web/html"
)

type ElementType string

const (
    TYPE_BUTTON   ElementType = "button"
    TYPE_CHECKBOX ElementType = "checkbox"
    TYPE_DATE	  ElementType = "date"
    TYPE_EMAIL    ElementType = "email"
    TYPE_FILE     ElementType = "file"
    TYPE_HIDDEN   ElementType = "hidden"
    TYPE_IMAGE    ElementType = "image"
    TYPE_NUMBER   ElementType = "number"
    TYPE_PASSWORD ElementType = "password"
    TYPE_RADIO    ElementType = "radio"
    TYPE_SUBMIT   ElementType = "submit"
    TYPE_SELECT   ElementType = "select"
    TYPE_TEXT     ElementType = "text"
    TYPE_TEXTAREA ElementType = "textarea"
    TYPE_TIME     ElementType = "time"
    TYPE_URL      ElementType = "url"
)

type element struct {
    Type ElementType
	html.Tagger
    validators []validate.Validater
}

func newElement(elemType ElementType, name string, tag html.TagType) *element {
    return &element{elemType, html.NewTag(name, tag), make([]validate.Validater, 0)}
}

func newInput(inputType ElementType) *element {
	return newElement(inputType, "input", html.TYPE_SELF_CLOSED_STRICT)
}

func NewButton(name string) *element {
    e := newElement(TYPE_BUTTON, "button", html.TYPE_OPEN_CLOSE)
    e.SetAttr("name", name)
    return e
}

func NewText(name string) *element {
	e := newInput(TYPE_TEXT)
	e.SetAttr("name", name)
    return e
}



