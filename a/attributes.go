package a

import "github.com/westcoastcode-se/gohtml"

func ID(id string) gohtml.Node {
	return gohtml.Attribute("id", id)
}

func Class(classes string) gohtml.Node {
	return gohtml.Attribute("class", classes)
}

func Href(h string) gohtml.Node {
	return gohtml.Attribute("href", h)
}

func Src(src string) gohtml.Node {
	return gohtml.Attribute("src", src)
}

func Role(classes string) gohtml.Node {
	return gohtml.Attribute("role", classes)
}

func Integrity(val string) gohtml.Node {
	return gohtml.Attribute("integrity", val)
}

const RelStylesheet = "stylesheet"
const RelIcon = "icon"

func Rel(rel string) gohtml.Node {
	return gohtml.Attribute("rel", rel)
}

const CrossOriginAnonymous = "anonymous"

func CrossOrigin(val string) gohtml.Node {
	return gohtml.Attribute("crossorigin", val)
}

func Charset(val string) gohtml.Node {
	return gohtml.Attribute("charset", val)
}

func Name(val string) gohtml.Node {
	return gohtml.Attribute("name", val)
}

func Content(val string) gohtml.Node {
	return gohtml.Attribute("content", val)
}

func Scope(val string) gohtml.Node {
	return gohtml.Attribute("scope", val)
}
