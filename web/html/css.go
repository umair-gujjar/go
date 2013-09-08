// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

type Classer interface {
    SetClass(string) *Classer
    GetClasses() []string
    ResetClass() *Classer
}

type InlineStyler interface {
    SetStyle(propertie, val string)
    GetStyle(string) string
    GetStyles() map[string]string
    HasStyle(string) bool
    ResetStyle() *InlineStyler
}
type Css interface {
    Classer
    InlineStyler
}

type css struct {
    styles  map[string]string
    classes []string
}

func NewCSS() *css {
    return &css{make(map[string]string), make([]string, 0)}
}

// Sets the css style of this tag
func (c *css) SetStyle(propertie, val string) *css {
    c.styles[propertie] = val
    return c
}

// Returns the css style of this tag
func (c *css) GetStyle(propertie string) string {
    if c.HasStyle(propertie) {
        return c.styles["style"]
    }
    return ""
}

// Returns all styles of this tag
func (c css) GetStyles() map[string]string {
    return c.styles
}

// The style of this tag exist?
func (c css) HasStyle(propertie string) bool {
    _, isOk := c.styles[propertie]
    return isOk
}

// Resets and removes all styles
func (c *css) ResetStyle() *css {
    c.styles = make(map[string]string)
    return c
}

// Initialize the rendering style of this tag
func (c *css) styleRender() string {
    if len(c.GetStyles()) > 0 {
        style := ""
        for propertie, val := range c.GetStyles() {
            style += propertie + ":" + val + ";"
        }
        return style[:len(style)-1]
    }
    return ""
}

// The a class of this tag
func (c *css) SetClass(name string) *css {
    c.classes = append(c.classes, name)
    return c
}

// Returns all classes of this tag
func (c css) GetClasses() []string {
    return c.classes
}

// Resets and removes all classes of this tag
func (c *css) ResetClass() *css {
    c.classes = make([]string, 0)
    return c
}

// Initialize the rendering classes of this tag
func (c *css) classesRender() string {
    if len(c.classes) > 0 {
        classes := ""
        for _, class := range c.GetClasses() {
            classes += class + " "
        }
        return classes[:len(classes)-1]
    }
    return ""
}
