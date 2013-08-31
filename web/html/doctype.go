// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

var doctype map[Doctype]string

type Doctype string

const (
    HTML5           Doctype = "HTML 5"
    HTML401_STRICT  Doctype = "HTML 4.01 Strict"
    HTML401_TRANS   Doctype = "HTML 4.01 Transitional"
    HTML401_FRAME   Doctype = "HTML 4.01 Frameset"
    HTML20          Doctype = "HTML 2.0"
    HTML32          Doctype = "HTML 2.2"
    XHTML_10_STRICT Doctype = "XHTML 1.0 Strict"
    XHTML_10_TANS   Doctype = "XHTML 1.0 Transitional"
    XHTML_10_FRAME  Doctype = "XHTML 1.0 Frameset"
    XHTML_11        Doctype = "XHTML 1.1"
    XHTML_BASIC_11  Doctype = "XHTML Basic 1.1"
    XHTML_BASIC_10  Doctype = "XHTML Basic 1.0"
    MATH_ML20       Doctype = "MathML 2.0"
    MATH_ML_101     Doctype = "MathML 1.01"
)

func init() {
    doctype = make(map[Doctype]string)

    eol := "\n"
    tab := "\t"

    // HTML 5
    doctype[HTML5] = `<!DOCTYPE HTML>` + eol

    // HTML 4.01 Strict
    doctype[HTML401_STRICT] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"` + eol
    doctype[HTML401_STRICT] += tab + `"http://www.w3.org/TR/html4/strict.dtd">` + eol

    // HTML 4.01 Transitional
    doctype[HTML401_TRANS] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"` + eol
    doctype[HTML401_TRANS] += tab + `"http://www.w3.org/TR/html4/loose.dtd">` + eol

    // HTML 4.01 Frameset
    doctype[HTML401_FRAME] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"` + eol
    doctype[HTML401_FRAME] += tab + `"http://www.w3.org/TR/html4/frameset.dtd">` + eol

    // XHTML 1.0 Strict
    doctype[XHTML_10_STRICT] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"` + eol
    doctype[XHTML_10_STRICT] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">` + eol

    // XHTML 1.0 Transitional
    doctype[XHTML_10_TANS] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"` + eol
    doctype[XHTML_10_TANS] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">` + eol

    // XHTML 1.0 Frameset
    doctype[XHTML_10_FRAME] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"` + eol
    doctype[XHTML_10_FRAME] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">` + eol

    // XHTML 1.1 DTD
    doctype[XHTML_11] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"` + eol
    doctype[XHTML_11] += tab + `"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">` + eol

    // XHTML Basic 1.1
    doctype[XHTML_BASIC_11] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"` + eol
    doctype[XHTML_BASIC_11] += `"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">` + eol
}

// Returns the doctype by name
func GetDoctype(name Doctype) string {
    return doctype[name]
}
