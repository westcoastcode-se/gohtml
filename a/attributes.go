package a

import (
	"github.com/westcoastcode-se/gohtml/h"
	"io"
	"strconv"
)

func Attrib(key string, value string) h.Node {
	return func(b byte, w io.Writer) byte {
		_, _ = w.Write([]byte{' '})
		_, _ = w.Write([]byte(key))
		_, _ = w.Write([]byte("=\""))
		_, _ = w.Write([]byte(value))
		_, _ = w.Write([]byte{'"'})
		return b
	}
}

func ID(id string) h.Node {
	return Attrib("id", id)
}

func Class(classes string) h.Node {
	return Attrib("class", classes)
}

func ClassIf(classes string, optional bool) h.Node {
	return h.NodeIf(optional, Class(classes))
}

// ClassesIf emits a class attribute with zero or many values. Each value is expected to be emitted from a h.Raw functions
func ClassesIf(values ...h.Node) h.Node {
	return h.Attribs("class", values...)
}

func Href(val string) h.Node {
	return Attrib("href", val)
}

func Src(src string) h.Node {
	return Attrib("src", src)
}

func Role(classes string) h.Node {
	return Attrib("role", classes)
}

func Integrity(val string) h.Node {
	return Attrib("integrity", val)
}

const RelStylesheet = "stylesheet"
const RelIcon = "icon"

func Rel(rel string) h.Node {
	return Attrib("rel", rel)
}

const CrossOriginAnonymous = "anonymous"

func CrossOrigin(val string) h.Node {
	return Attrib("crossorigin", val)
}

func Charset(val string) h.Node {
	return Attrib("charset", val)
}

func Name(val string) h.Node {
	return Attrib("name", val)
}

func Content(val string) h.Node {
	return Attrib("content", val)
}

func Scope(val string) h.Node {
	return Attrib("scope", val)
}

func Style(val string) h.Node {
	return Attrib("style", val)
}

func Lang(val string) h.Node {
	return Attrib("lang", val)
}

func Type(t string) h.Node {
	return Attrib("type", t)
}

func OnClick(val string) h.Node {
	return Attrib("onclick", val)
}

func Method(val string) h.Node {
	return Attrib("method", val)
}

func Action(val string) h.Node {
	return Attrib("action", val)
}

func For(val string) h.Node {
	return Attrib("for", val)
}

func Colspan(n int) h.Node {
	return Attrib("colspan", strconv.Itoa(n))
}

func Rowspan(n int) h.Node {
	return Attrib("rowspan", strconv.Itoa(n))
}

func Value(v string) h.Node {
	return Attrib("value", v)
}
