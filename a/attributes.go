package a

import (
	"github.com/westcoastcode-se/gohtml/h"
	"strconv"
)

func ID(id string) h.Node {
	return h.Attrib("id", id)
}

func Class(classes string) h.Node {
	return h.Attrib("class", classes)
}

func ClassIf(classes string, optional bool) h.Node {
	return h.AttribIf(optional, func() h.Node {
		return Class(classes)
	})
}

// ClassesIf emits a class attribute with zero or many values. Each value is expected to be emitted from a h.Raw functions
func ClassesIf(values ...h.Node) h.Node {
	return h.Attribs("class", values...)
}

func Href(val string) h.Node {
	return h.Attrib("href", val)
}

func Src(src string) h.Node {
	return h.Attrib("src", src)
}

func Role(classes string) h.Node {
	return h.Attrib("role", classes)
}

func Integrity(val string) h.Node {
	return h.Attrib("integrity", val)
}

const RelStylesheet = "stylesheet"
const RelIcon = "icon"

func Rel(rel string) h.Node {
	return h.Attrib("rel", rel)
}

const CrossOriginAnonymous = "anonymous"

func CrossOrigin(val string) h.Node {
	return h.Attrib("crossorigin", val)
}

func Charset(val string) h.Node {
	return h.Attrib("charset", val)
}

func Name(val string) h.Node {
	return h.Attrib("name", val)
}

func Content(val string) h.Node {
	return h.Attrib("content", val)
}

func Scope(val string) h.Node {
	return h.Attrib("scope", val)
}

func Style(val string) h.Node {
	return h.Attrib("style", val)
}

func Lang(val string) h.Node {
	return h.Attrib("lang", val)
}

func Type(t string) h.Node {
	return h.Attrib("type", t)
}

func OnClick(val string) h.Node {
	return h.Attrib("onclick", val)
}

func Method(val string) h.Node {
	return h.Attrib("method", val)
}

func Action(val string) h.Node {
	return h.Attrib("action", val)
}

func For(val string) h.Node {
	return h.Attrib("for", val)
}

func Colspan(n int) h.Node {
	return h.Attrib("colspan", strconv.Itoa(n))
}

func Rowspan(n int) h.Node {
	return h.Attrib("rowspan", strconv.Itoa(n))
}

func Value(v string) h.Node {
	return h.Attrib("value", v)
}
