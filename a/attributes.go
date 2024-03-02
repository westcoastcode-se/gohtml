package a

import (
	"github.com/westcoastcode-se/gohtml"
)

func ID(id string) gohtml.Node {
	return gohtml.Attrib("id", id)
}

func Class(classes string) gohtml.Node {
	return gohtml.Attrib("class", classes)
}

func ClassIf(classes string, optional bool) gohtml.Node {
	return gohtml.AttribIf(optional, func() gohtml.Node {
		return Class(classes)
	})
}

// ClassesIf emits a class attribute with zero or many values. Each value is expected to be emitted from a gohtml.Value functions
func ClassesIf(values ...gohtml.Node) gohtml.Node {
	return gohtml.Attribs("class", values...)
}

func Href(h string) gohtml.Node {
	return gohtml.Attrib("href", h)
}

func Src(src string) gohtml.Node {
	return gohtml.Attrib("src", src)
}

func Role(classes string) gohtml.Node {
	return gohtml.Attrib("role", classes)
}

func Integrity(val string) gohtml.Node {
	return gohtml.Attrib("integrity", val)
}

const RelStylesheet = "stylesheet"
const RelIcon = "icon"

func Rel(rel string) gohtml.Node {
	return gohtml.Attrib("rel", rel)
}

const CrossOriginAnonymous = "anonymous"

func CrossOrigin(val string) gohtml.Node {
	return gohtml.Attrib("crossorigin", val)
}

func Charset(val string) gohtml.Node {
	return gohtml.Attrib("charset", val)
}

func Name(val string) gohtml.Node {
	return gohtml.Attrib("name", val)
}

func Content(val string) gohtml.Node {
	return gohtml.Attrib("content", val)
}

func Scope(val string) gohtml.Node {
	return gohtml.Attrib("scope", val)
}

func Style(val string) gohtml.Node {
	return gohtml.Attrib("style", val)
}

func Lang(val string) gohtml.Node {
	return gohtml.Attrib("lang", val)
}
