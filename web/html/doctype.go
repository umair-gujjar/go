// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

var doctypes map[Doctype]string

type Doctype string

const (
    HTML5          Doctype = "HTML 5"
    HTML401_STRICT Doctype = "HTML 4.01 Strict"
    HTML401_TRANS  Doctype = "HTML 4.01 Transitional"
    HTML401_FRAME  Doctype = "HTML 4.01 Frameset"
    HTML20         Doctype = "HTML 2.0"
    HTML32         Doctype = "HTML 3.2"
    XHTML10_STRICT Doctype = "XHTML 1.0 Strict"
    XHTML10_TRANS  Doctype = "XHTML 1.0 Transitional"
    XHTML10_FRAME  Doctype = "XHTML 1.0 Frameset"
    XHTML11        Doctype = "XHTML 1.1"
    XHTML_BASIC11  Doctype = "XHTML Basic 1.1"
    XHTML_BASIC10  Doctype = "XHTML Basic 1.0"
    MATH_ML20      Doctype = "MathML 2.0"
    MATH_ML101     Doctype = "MathML 1.01"
    SVG11_FULL     Doctype = "SVG 1.1 Full"
    SVG10          Doctype = "SVG 1.0"
    SVG11_BASIC    Doctype = "SVG 1.1 Basic"
    SVG11_TINY     Doctype = "SVG 1.1 Tiny"
)

func init() {
    doctypes = make(map[Doctype]string)

    eol := "\n"
    tab := "\t"

    // HTML 5
    doctypes[HTML5] = `<!DOCTYPE HTML>` + eol

    // HTML 4.01 Strict
    doctypes[HTML401_STRICT] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN"` + eol
    doctypes[HTML401_STRICT] += tab + `"http://www.w3.org/TR/html4/strict.dtd">` + eol

    // HTML 4.01 Transitional
    doctypes[HTML401_TRANS] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"` + eol
    doctypes[HTML401_TRANS] += tab + `"http://www.w3.org/TR/html4/loose.dtd">` + eol

    // HTML 4.01 Frameset
    doctypes[HTML401_FRAME] = `<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Frameset//EN"` + eol
    doctypes[HTML401_FRAME] += tab + `"http://www.w3.org/TR/html4/frameset.dtd">` + eol

    // HTML 2.0 - DTD
    doctypes[HTML20] = `<!DOCTYPE html PUBLIC "-//IETF//DTD HTML 2.0//EN">` + eol

    // HTML 3.2 - DTD
    doctypes[HTML32] = `<!DOCTYPE html PUBLIC "-//W3C//DTD HTML 3.2 Final//EN">` + eol

    // XHTML 1.0 Strict
    doctypes[XHTML10_STRICT] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Strict//EN"` + eol
    doctypes[XHTML10_STRICT] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-strict.dtd">` + eol

    // XHTML 1.0 Transitional
    doctypes[XHTML10_TRANS] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"` + eol
    doctypes[XHTML10_TRANS] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">` + eol

    // XHTML 1.0 Frameset
    doctypes[XHTML10_FRAME] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Frameset//EN"` + eol
    doctypes[XHTML10_FRAME] += tab + `"http://www.w3.org/TR/xhtml1/DTD/xhtml1-frameset.dtd">` + eol

    // XHTML 1.1 DTD
    doctypes[XHTML11] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.1//EN"` + eol
    doctypes[XHTML11] += tab + `"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">` + eol

    // XHTML Basic 1.1
    doctypes[XHTML_BASIC11] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.1//EN"` + eol
    doctypes[XHTML_BASIC11] += tab + `"http://www.w3.org/TR/xhtml-basic/xhtml-basic11.dtd">` + eol

    // XHTML Basic 1.0
    doctypes[XHTML_BASIC10] = `<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML Basic 1.0//EN"` + eol
    doctypes[XHTML_BASIC10] += tab + `"http://www.w3.org/TR/xhtml-basic/xhtml-basic10.dtd">` + eol

    // MathML 2.0 - DTD
    doctypes[MATH_ML20] = `<!DOCTYPE math PUBLIC "-//W3C//DTD MathML 2.0//EN"` + eol
    doctypes[MATH_ML20] += tab + `"http://www.w3.org/Math/DTD/mathml2/mathml2.dtd">` + eol

    // MathML 1.01 - DTD
    doctypes[MATH_ML101] = `<!DOCTYPE math SYSTEM ` + eol
    doctypes[MATH_ML101] += tab + `"http://www.w3.org/Math/DTD/mathml1/mathml.dtd">` + eol

    // SVG 1.1 Full
    doctypes[SVG11_FULL] = `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN"` + eol
    doctypes[SVG11_FULL] += tab + `"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">` + eol

    // SVG 1.0
    doctypes[SVG10] = `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.0//EN"` + eol
    doctypes[SVG10] += tab + `"http://www.w3.org/TR/2001/REC-SVG-20010904/DTD/svg10.dtd">` + eol

    // SVG 1.1 Basic
    doctypes[SVG11_BASIC] = `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Basic//EN"` + eol
    doctypes[SVG11_BASIC] += tab + `"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-basic.dtd">` + eol

    // SVG 1.1 Tiny
    doctypes[SVG11_TINY] = `<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1 Tiny//EN"` + eol
    doctypes[SVG11_TINY] += tab + `"http://www.w3.org/Graphics/SVG/1.1/DTD/svg11-tiny.dtd">` + eol
}

// Returns the doctype by name
func GetDoctype(name Doctype) string {
    return doctypes[name]
}
