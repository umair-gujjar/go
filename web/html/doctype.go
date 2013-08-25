// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

type Doctype string

const (
    HTML5		    Doctype = "HTML 5"
    HTML401_STRICT 	Doctype = "HTML 4.01 Strict"
    HTML401_TRANS	Doctype = "HTML 4.01 Transitional"
    HTML401_FRAME	Doctype = "HTML 4.01 Frameset"
	HTML20 		    Doctype = "HTML 2.0"
	HTML32 		    Doctype = "HTML 2.2"
	XHTML_10_STRICT Doctype = "XHTML 1.0 Strict"
	XHTML_10_TANS   Doctype = "XHTML 1.0 Transitional"
	XHTML_10_FRAME  Doctype = "XHTML 1.0 Frameset"
	XHTML_11	    Doctype = "XHTML 1.1"
	XHTML_BASIC_11  Doctype = "XHTML Basic 1.1"
	XHTML_BASIC_10  Doctype = "XHTML Basic 1.0"
    MATH_ML20	    Doctype = "MathML 2.0"
    MATH_ML_101		Doctype = "MathML 1.01"
)
