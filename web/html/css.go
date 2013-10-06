// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

type Classes []string
type Styles map[string]string

type Styler interface {
    SetClass(string) *css
    HasClass(string) (int, bool)
    SetClasses(Classes) *css
    GetClasses() Classes
    DelClass(string) *css
    ResetClass() *css
    SetStyle(property, val string) *css
    GetStyle(string) string
    SetStyles(Styles) *css
    GetStyles() Styles
    HasStyle(string) bool
    DelStyle(string) *css
    ResetStyle()
}

type css struct {
    classes Classes
    styles  Styles
}

func NewCSS() *css {
    return &css{make(Classes, 0), make(Styles)}
}

// Set a new class
func (c *css) SetClass(name string) *css {
    c.classes = append(c.classes, name)
    return c
}

// Checks if the class exists.
// If the classes are the function returns its index and true,
// otherwise it returns 0 and false.
func (c css) HasClass(name string) (int, bool) {
    for i, clsName := range c.GetClasses() {
        if clsName == name {
            return i, true
        }
    }
    return 0, false
}

// Delete a class
func (c *css) DelClass(name string) *css {
    if i, ok := c.HasClass(name); ok {
        copy(c.classes[i:], c.classes[i+1:])
        c.classes[len(c.classes)-1] = ""
        c.classes = c.classes[:len(c.classes)-1]
    }
    return c
}

// Set some classes
func (c *css) SetClasses(classes []string) *css {
    c.classes = classes
    return c
}

// Returns all classes
func (c css) GetClasses() []string {
    return c.classes
}

// Resets all classes
func (c *css) ResetClass() *css {
    c.classes = make([]string, 0)
    return c
}

// Initialize the rendering classes
func (c *css) classesToString() string {
    if len(c.classes) > 0 {
        classes := ""
        for _, class := range c.GetClasses() {
            classes += class + " "
        }
        return classes[:len(classes)-1]
    }
    return ""
}

// Sets the css style of this tag
func (c *css) SetStyle(property, val string) *css {
    c.styles[property] = val
    return c
}

// Returns the css style of this tag
func (c *css) GetStyle(property string) string {
    if c.HasStyle(property) {
        return c.styles["style"]
    }
    return ""
}

// Set some styles
func (c *css) SetStyles(styles Styles) *css {
    c.styles = styles
    return c
}

// Returns all styles of this tag
func (c css) GetStyles() Styles {
    return c.styles
}

// Deletes a style property
func (c *css) DelStyle(property string) *css {
    if c.HasStyle(property) {
        delete(c.styles, property)
    }
    return c
}

// The style of this tag exist?
func (c css) HasStyle(property string) bool {
    _, ok := c.styles[property]
    return ok
}

// Resets and removes all styles
func (c *css) ResetStyle() *css {
    c.styles = make(map[string]string)
    return c
}

// Initialize the rendering style of this tag
func (c *css) stylesToString() string {
    if len(c.GetStyles()) > 0 {
        style := ""
        for property, val := range c.GetStyles() {
            style += property + ": " + val + "; "
        }
        return style[:len(style)-1]
    }
    return ""
}

// Applies a css element
func (c *css) ApplyTo(e ElementHandler) *css {
    if len(c.classes) > 0 {
        e.SetAttr("class", c.classesToString())
    }
    if len(c.styles) > 0 {
        e.SetAttr("style", c.stylesToString())
    }
    return c
}
