package h

import (
	"fmt"
	"io"
	"log"
)

type RootNode func(w io.Writer) (int, error)

type Node func(b byte, w io.Writer) byte

// NodeIf evaluates the node if the supplied test is true
func NodeIf(test bool, n Node) Node {
	return func(b byte, w io.Writer) byte {
		if !test {
			return b
		}
		return n(b, w)
	}
}

// Attribs gives you a way of creating an attribute with multiple values using the Value, ValueIf functions
func Attribs(key string, values ...Node) Node {
	return func(b byte, w io.Writer) byte {
		_, _ = w.Write([]byte{' '})
		_, _ = w.Write([]byte(key))
		_, _ = w.Write([]byte("=\""))
		for _, a := range values {
			_, _ = w.Write([]byte{' '})
			b = a(b, w)
		}
		_, _ = w.Write([]byte{'"'})
		return b
	}
}

// Raw simply writes the byte content directly to the io.Writer and does nothing else
func Raw(value string) Node {
	return func(b byte, w io.Writer) byte {
		_, _ = w.Write([]byte(value))
		return b
	}
}

func RawIf(value string, test bool) Node {
	return func(b byte, w io.Writer) byte {
		if test {
			_, _ = w.Write([]byte(value))
		}
		return b
	}
}

// TagEmpty writes a single HTML tag that have 0 or more attributes
func TagEmpty(name string, c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte{'<'})
		_, _ = w.Write([]byte(name))
		b = '>'
		for _, cc := range c {
			b = cc(b, w)
		}
		_, _ = w.Write([]byte("/>"))
		return 0
	}
}

// Tag writes a single HTML tag that have 0 or more attributes and 0 or more child-nodes
func Tag(name string, c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte{'<'})
		_, _ = w.Write([]byte(name))
		b = '>'
		for _, cc := range c {
			b = cc(b, w)
		}
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte("</"))
		_, _ = w.Write([]byte(name))
		_, _ = w.Write([]byte{'>'})
		return 0
	}
}

// Join nodes multiple nodes into a single Node
func Join(c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, cc := range c {
			b = cc(b, w)
		}
		return b
	}
}

// Empty is a way to represent nothing
func Empty() Node {
	return func(b byte, w io.Writer) byte {
		return b
	}
}

func ExpandArray(c []Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, cc := range c {
			b = cc(b, w)
		}
		return b
	}
}

// Bytes writes the supplied bytes as if it is a text html element
func Bytes(value []byte) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write(value)
		return 0
	}
}

// Text creates a text html element
func Text(text string) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte(text))
		return 0
	}
}

// Textf creates a formatted text html element
func Textf(format string, a ...any) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}

		_, _ = w.Write([]byte(fmt.Sprintf(format, a...)))
		return 0
	}
}

// TextIf creates a simple text string if the test is true
func TextIf(text string, test bool) Node {
	return func(b byte, w io.Writer) byte {
		if !test {
			return b
		}
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte(text))
		return 0
	}
}

// TextfIf creates a formatted text string if the supplied test is true
func TextfIf(format string, test bool, a ...any) Node {
	return func(b byte, w io.Writer) byte {
		if !test {
			return b
		}
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte(fmt.Sprintf(format, a...)))
		return 0
	}
}

// Log is logging a message when this node is being processed
func Log(v ...interface{}) Node {
	return func(b byte, _ io.Writer) byte {
		log.Print(v...)
		return b
	}
}

// Logf is logging a message when this node is being processed
func Logf(format string, a ...interface{}) Node {
	return func(b byte, _ io.Writer) byte {
		log.Printf(format, a...)
		return b
	}
}

// Comment creates a comment tag
func Comment(c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte("<!--"))
		for _, cc := range c {
			b = cc(0, w)
		}
		_, _ = w.Write([]byte("-->"))
		return 0
	}
}

func A(c ...Node) Node {
	return Tag("a", c...)
}

func Abbr(c ...Node) Node {
	return Tag("abbr", c...)
}

func Address(c ...Node) Node {
	return Tag("address", c...)
}

func Area(c ...Node) Node {
	return TagEmpty("area", c...)
}

func Article(c ...Node) Node {
	return Tag("article", c...)
}

func Aside(c ...Node) Node {
	return Tag("aside", c...)
}

func Audio(c ...Node) Node {
	return Tag("audio", c...)
}

func B(c ...Node) Node {
	return Tag("b", c...)
}

func Base(c ...Node) Node {
	return TagEmpty("base", c...)
}

func Bdi(c ...Node) Node {
	return Tag("bdi", c...)
}

func Bdo(c ...Node) Node {
	return Tag("bdo", c...)
}

func Blockquote(c ...Node) Node {
	return Tag("blockquote", c...)
}

func Body(c ...Node) Node {
	return Tag("body", c...)
}

func Br(c ...Node) Node {
	return TagEmpty("br", c...)
}

func Button(c ...Node) Node {
	return Tag("button", c...)
}

func Canvas(c ...Node) Node {
	return Tag("canvas", c...)
}

func Caption(c ...Node) Node {
	return Tag("caption", c...)
}

func Cite(c ...Node) Node {
	return Tag("cite", c...)
}

func Code(c ...Node) Node {
	return Tag("code", c...)
}

func Col(c ...Node) Node {
	return Tag("col", c...)
}

func Colgroup(c ...Node) Node {
	return Tag("colgroup", c...)
}

func Data(c ...Node) Node {
	return Tag("data", c...)
}

func Datalist(c ...Node) Node {
	return Tag("datalist", c...)
}

func Dd(c ...Node) Node {
	return Tag("dd", c...)
}

func Del(c ...Node) Node {
	return Tag("del", c...)
}

func Details(c ...Node) Node {
	return Tag("details", c...)
}

func Dfn(c ...Node) Node {
	return Tag("dfn", c...)
}

func Dialog(c ...Node) Node {
	return Tag("dialog", c...)
}

func Div(c ...Node) Node {
	return Tag("div", c...)
}

func Dl(c ...Node) Node {
	return Tag("dl", c...)
}

func Dt(c ...Node) Node {
	return Tag("dt", c...)
}

func Em(c ...Node) Node {
	return Tag("em", c...)
}

func Embed(c ...Node) Node {
	return TagEmpty("embed", c...)
}

func Fieldset(c ...Node) Node {
	return Tag("fieldset", c...)
}

func Figcaption(c ...Node) Node {
	return Tag("figcaption", c...)
}

func Figure(c ...Node) Node {
	return Tag("figure", c...)
}

func Footer(c ...Node) Node {
	return Tag("footer", c...)
}

func Form(c ...Node) Node {
	return Tag("form", c...)
}

func H1(c ...Node) Node {
	return Tag("h1", c...)
}

func H2(c ...Node) Node {
	return Tag("h2", c...)
}

func H3(c ...Node) Node {
	return Tag("h3", c...)
}

func H4(c ...Node) Node {
	return Tag("h4", c...)
}

func H5(c ...Node) Node {
	return Tag("h5", c...)
}

func H6(c ...Node) Node {
	return Tag("h6", c...)
}

func Head(c ...Node) Node {
	return Tag("head", c...)
}

func Header(c ...Node) Node {
	return Tag("header", c...)
}

func Hgroup(c ...Node) Node {
	return Tag("hgroup", c...)
}

func Hr(c ...Node) Node {
	return TagEmpty("hr", c...)
}

func Html(c ...Node) RootNode {
	return func(w io.Writer) (int, error) {
		we := &errorAwareWriter{
			w: w,
		}
		if _, err := we.Write([]byte("<!DOCTYPE html><html")); err != nil {
			return 0, err
		}
		b := byte('>')
		for _, cc := range c {
			b = cc(b, we)
		}
		if b != 0 {
			_, _ = we.Write([]byte{b})
		}
		_, _ = we.Write([]byte("</html>"))
		return we.len, we.err
	}
}

func I(c ...Node) Node {
	return Tag("i", c...)
}

func Iframe(c ...Node) Node {
	return TagEmpty("iframe", c...)
}

func Img(c ...Node) Node {
	return TagEmpty("img", c...)
}

func Input(c ...Node) Node {
	return TagEmpty("input", c...)
}

func Ins(c ...Node) Node {
	return Tag("ins", c...)
}

func Kbd(c ...Node) Node {
	return Tag("kbd", c...)
}

func Label(c ...Node) Node {
	return Tag("label", c...)
}

func Legend(c ...Node) Node {
	return Tag("legend", c...)
}

func Li(c ...Node) Node {
	return Tag("li", c...)
}

func Link(c ...Node) Node {
	return TagEmpty("link", c...)
}

func Main(c ...Node) Node {
	return Tag("main", c...)
}

func Map(c ...Node) Node {
	return Tag("map", c...)
}

func Mark(c ...Node) Node {
	return Tag("mark", c...)
}

func Menu(c ...Node) Node {
	return Tag("menu", c...)
}

func Meta(c ...Node) Node {
	return TagEmpty("meta", c...)
}

func Meter(c ...Node) Node {
	return Tag("meter", c...)
}

func Nav(c ...Node) Node {
	return Tag("nav", c...)
}

func Noscript(c ...Node) Node {
	return Tag("noscript", c...)
}

func Object(c ...Node) Node {
	return Tag("object", c...)
}

func Ol(c ...Node) Node {
	return Tag("ol", c...)
}

func Optgroup(c ...Node) Node {
	return Tag("optgroup", c...)
}

func Option(c ...Node) Node {
	return Tag("option", c...)
}

func Output(c ...Node) Node {
	return Tag("output", c...)
}

func P(c ...Node) Node {
	return Tag("p", c...)
}

func Param(c ...Node) Node {
	return TagEmpty("param", c...)
}

func Picture(c ...Node) Node {
	return Tag("picture", c...)
}

func Pre(c ...Node) Node {
	return Tag("pre", c...)
}

func Progress(c ...Node) Node {
	return Tag("progress", c...)
}

func Q(c ...Node) Node {
	return Tag("q", c...)
}

func Rp(c ...Node) Node {
	return Tag("rp", c...)
}

func Rt(c ...Node) Node {
	return Tag("rt", c...)
}

func Ruby(c ...Node) Node {
	return Tag("ruby", c...)
}

func S(c ...Node) Node {
	return Tag("s", c...)
}

func Samp(c ...Node) Node {
	return Tag("samp", c...)
}

func Script(c ...Node) Node {
	return Tag("script", c...)
}

func Search(c ...Node) Node {
	return Tag("search", c...)
}

func Section(c ...Node) Node {
	return Tag("section", c...)
}

func Select(c ...Node) Node {
	return Tag("select", c...)
}

func Small(c ...Node) Node {
	return Tag("small", c...)
}

func Source(c ...Node) Node {
	return TagEmpty("source", c...)
}

func Span(c ...Node) Node {
	return Tag("span", c...)
}

func Strong(c ...Node) Node {
	return Tag("strong", c...)
}

func Style(c ...Node) Node {
	return Tag("style", c...)
}

func Sub(c ...Node) Node {
	return Tag("sub", c...)
}

func Summary(c ...Node) Node {
	return Tag("summary", c...)
}

func Sup(c ...Node) Node {
	return Tag("sup", c...)
}

func Svg(c ...Node) Node {
	return Tag("svg", c...)
}

func Table(c ...Node) Node {
	return Tag("table", c...)
}

func Tbody(c ...Node) Node {
	return Tag("tbody", c...)
}

func Td(c ...Node) Node {
	return Tag("td", c...)
}

func Template(c ...Node) Node {
	return Tag("template", c...)
}

func Textarea(c ...Node) Node {
	return Tag("textarea", c...)
}

func Tfoot(c ...Node) Node {
	return Tag("tfoot", c...)
}

func Th(c ...Node) Node {
	return Tag("th", c...)
}

func Thead(c ...Node) Node {
	return Tag("thead", c...)
}

func Time(c ...Node) Node {
	return Tag("time", c...)
}

func Title(title string) Node {
	return Tag("title", Text(title))
}

func Tr(c ...Node) Node {
	return Tag("tr", c...)
}

func Track(c ...Node) Node {
	return TagEmpty("track", c...)
}

func U(c ...Node) Node {
	return Tag("u", c...)
}

func Ul(c ...Node) Node {
	return Tag("ul", c...)
}

func Var(c ...Node) Node {
	return Tag("var", c...)
}

func Video(c ...Node) Node {
	return Tag("video", c...)
}

func Wbr(c ...Node) Node {
	return Tag("wbr", c...)
}
