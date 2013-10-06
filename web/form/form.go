// Copyright 2013 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package form

var validGlobalAttrs map[string]bool

func init() {
	validGlobalAttrs = make(map[string]bool)
    validGlobalAttrs["accesskey"] = true
    validGlobalAttrs["class"] = true
    validGlobalAttrs["checked"] = true
    validGlobalAttrs["contenteditable"] = true
    validGlobalAttrs["contextmenu"] = true
    validGlobalAttrs["dir"] = true
    validGlobalAttrs["draggable"] = true
    validGlobalAttrs["dropzone"] = true
    validGlobalAttrs["hidden"] = true
    validGlobalAttrs["id"] = true
    validGlobalAttrs["lang"] = true
    validGlobalAttrs["onabort"] = true
    validGlobalAttrs["onblur"] = true
    validGlobalAttrs["oncanplay"] = true
    validGlobalAttrs["oncanplaythrough"] = true
    validGlobalAttrs["onchange"] = true
    validGlobalAttrs["onclick"] = true
    validGlobalAttrs["oncontextmenu"] = true
    validGlobalAttrs["ondblclick"] = true
    validGlobalAttrs["ondragend"] = true
    validGlobalAttrs["ondragenter"] = true
    validGlobalAttrs["ondragleave"] = true
    validGlobalAttrs["ondragover"] = true
    validGlobalAttrs["ondragstart"] = true
    validGlobalAttrs["ondrop"] = true
    validGlobalAttrs["ondurationchange"] = true
    validGlobalAttrs["onemptied"] = true
    validGlobalAttrs["onended"] = true
    validGlobalAttrs["onerror"] = true
    validGlobalAttrs["onfocus"] = true
    validGlobalAttrs["oninput"] = true
    validGlobalAttrs["oninvalid"] = true
    validGlobalAttrs["onkeydown"] = true
    validGlobalAttrs["onkeypress"] = true
    validGlobalAttrs["onkeyup"] = true
    validGlobalAttrs["onload"] = true
    validGlobalAttrs["onloadeddata"] = true
    validGlobalAttrs["onloadedmetadata"] = true
    validGlobalAttrs["onloadstart"] = true
    validGlobalAttrs["onmousedown"] = true
    validGlobalAttrs["onmousemove"] = true
    validGlobalAttrs["onmouseout"] = true
    validGlobalAttrs["onmouseover"] = true
    validGlobalAttrs["onmouseup"] = true
    validGlobalAttrs["onmousewheel"] = true
    validGlobalAttrs["onpause"] = true
    validGlobalAttrs["onplay"] = true
    validGlobalAttrs["onplaying"] = true
    validGlobalAttrs["onprogress"] = true
    validGlobalAttrs["onratechange"] = true
    validGlobalAttrs["onreadystatechange"] = true
    validGlobalAttrs["onreset"] = true
    validGlobalAttrs["onscroll"] = true
    validGlobalAttrs["onseeked"] = true
    validGlobalAttrs["onselect"] = true
    validGlobalAttrs["onshow"] = true
    validGlobalAttrs["onstalled"] = true
    validGlobalAttrs["onsubmit"] = true
    validGlobalAttrs["onsuspend"] = true
    validGlobalAttrs["ontimeupdate"] = true
    validGlobalAttrs["onvolumechange"] = true
    validGlobalAttrs["onwaiting"] = true
    validGlobalAttrs["selected"] = true
    validGlobalAttrs["spellcheck"] = true
    validGlobalAttrs["style"] = true
    validGlobalAttrs["tabindex"] = true
    validGlobalAttrs["title"] = true
    validGlobalAttrs["xml:base"] = true
    validGlobalAttrs["xml:lang"] = true
    validGlobalAttrs["xml:space"] = true
}
