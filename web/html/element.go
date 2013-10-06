// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

type options map[string]interface{}

// Represents the interface of a html element
type ElementHandler interface {
    SetName(name string) *element
    GetName() string
    SetLabel(label string) *element
    GetLabel() string
    SetOption(key string, val interface{}) *element
    GetOption(key string) interface{}
    RemoveOption(key string) *element
    SetOptions(opt options) *element
    GetOptions() options
    SetAttr(name string, val string) *element
    SetAttrs(map[string]string) *element
    GetAttr(name string) string
    GetAttrs() []*attr
    HasAttr(name string) bool
    DelAttr(name string) *element
    ResetAttrs() *element
    LenAttrs() int
    SetVal(val string) *element
    GetVal() string
    SetMessages(messages []string) *element
    GetMessages() []string
}

type element struct {
    name     string
    label    string
    opts     options
    attrs    map[string]*attr
    val      string
    messages []string
}

// Creating a new element
func NewElement(name string) *element {
    return &element{
        name:  name,
        label: "",
        opts:  make(options),
        attrs: make(map[string]*attr),
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

// Set the label of this element
func (e *element) SetLabel(label string) *element {
    e.label = label
    return e
}

// Get the label of this element
func (e element) GetLabel() string {
    return e.label
}

// Set an option of this element
func (e *element) SetOption(key string, val interface{}) *element {
    e.opts[key] = val
    return e
}

// Get an option of this element
func (e element) GetOption(key string) interface{} {
    if val, ok := e.opts[key]; ok {
        return val
    }
    return nil
}

// Remove an option of this element
func (e *element) RemoveOption(key string) *element {
    if _, ok := e.opts[key]; ok {
        delete(e.opts, key)
    }
    return e
}

// Set some options of this element
func (e *element) SetOptions(opts options) *element {
    e.opts = opts
    return e
}

// Get all options of this element
func (e element) GetOptions() options {
    return e.opts
}

// Set an attribute to the current element
func (e *element) SetAttr(name string, val string) *element {
    attr := NewAttr(name, val)
    e.attrs[attr.GetName()] = attr
    return e
}

// Sets some attributes from map
func (e *element) SetAttrs(attrs map[string]string) *element {
    for name, val := range attrs {
        e.SetAttr(name, val)
    }
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
func (e element) LenAttrs() int {
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

// Set a list of messages to report when validation fails
func (e *element) SetMessages(messages []string) *element {
    e.messages = messages
    return e
}

// Get validation error message, if any
// Returns a list of validation failure messages, if any
func (e element) GetMessages() []string {
    return e.messages
}
