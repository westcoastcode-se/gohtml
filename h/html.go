package h

import (
	"fmt"
	"io"
	"log"
)

type RootNode func(w io.Writer) (int, error)

type Node func(b byte, w io.Writer) byte

// NodeIf gives you a way of creating a Node only if the supplied test is true
func NodeIf(test bool, f func() Node) Node {
	return func(b byte, w io.Writer) byte {
		if !test {
			return b
		}
		return f()(b, w)
	}
}

// AttribIf gives you a way of creating a Node only if the supplied test is true. This function
// is identical  as NodeIf
func AttribIf(test bool, f func() Node) Node {
	return NodeIf(test, f)
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

func Attrib(key string, value string) Node {
	return func(b byte, w io.Writer) byte {
		_, _ = w.Write([]byte{' '})
		_, _ = w.Write([]byte(key))
		_, _ = w.Write([]byte("=\""))
		_, _ = w.Write([]byte(value))
		_, _ = w.Write([]byte{'"'})
		return b
	}
}

func ElemEmpty(name string, c ...Node) Node {
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

// Elem writes a single HTML tag that might have 0 or more child-nodes
func Elem(name string, c ...Node) Node {
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

func Expand(c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, cc := range c {
			b = cc(b, w)
		}
		return b
	}
}

// Empty is a way to represent nothing. This is useful when
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
	return Elem("a", c...)
}

func Abbr(c ...Node) Node {
	return Elem("abbr", c...)
}

func Address(c ...Node) Node {
	return Elem("address", c...)
}

func Area(c ...Node) Node {
	return ElemEmpty("area", c...)
}

func Article(c ...Node) Node {
	return Elem("article", c...)
}

func Aside(c ...Node) Node {
	return Elem("aside", c...)
}

func Audio(c ...Node) Node {
	return Elem("audio", c...)
}

func B(c ...Node) Node {
	return Elem("b", c...)
}

func Base(c ...Node) Node {
	return ElemEmpty("base", c...)
}

func Bdi(c ...Node) Node {
	return Elem("bdi", c...)
}

func Bdo(c ...Node) Node {
	return Elem("bdo", c...)
}

func Blockquote(c ...Node) Node {
	return Elem("blockquote", c...)
}

func Body(c ...Node) Node {
	return Elem("body", c...)
}

func Br(c ...Node) Node {
	return ElemEmpty("br", c...)
}

func Button(c ...Node) Node {
	return Elem("button", c...)
}

func Canvas(c ...Node) Node {
	return Elem("canvas", c...)
}

func Caption(c ...Node) Node {
	return Elem("caption", c...)
}

func Cite(c ...Node) Node {
	return Elem("cite", c...)
}

func Code(c ...Node) Node {
	return Elem("code", c...)
}

func Col(c ...Node) Node {
	return Elem("col", c...)
}

func Colgroup(c ...Node) Node {
	return Elem("colgroup", c...)
}

func Data(c ...Node) Node {
	return Elem("data", c...)
}

func Datalist(c ...Node) Node {
	return Elem("datalist", c...)
}

func Dd(c ...Node) Node {
	return Elem("dd", c...)
}

func Del(c ...Node) Node {
	return Elem("del", c...)
}

func Details(c ...Node) Node {
	return Elem("details", c...)
}

func Dfn(c ...Node) Node {
	return Elem("dfn", c...)
}

func Dialog(c ...Node) Node {
	return Elem("dialog", c...)
}

func Div(c ...Node) Node {
	return Elem("div", c...)
}

func Dl(c ...Node) Node {
	return Elem("dl", c...)
}

func Dt(c ...Node) Node {
	return Elem("dt", c...)
}

func Em(c ...Node) Node {
	return Elem("em", c...)
}

func Embed(c ...Node) Node {
	return ElemEmpty("embed", c...)
}

func Fieldset(c ...Node) Node {
	return Elem("fieldset", c...)
}

func Figcaption(c ...Node) Node {
	return Elem("figcaption", c...)
}

func Figure(c ...Node) Node {
	return Elem("figure", c...)
}

func Footer(c ...Node) Node {
	return Elem("footer", c...)
}

func Form(c ...Node) Node {
	return Elem("form", c...)
}

func H1(c ...Node) Node {
	return Elem("h1", c...)
}

func H2(c ...Node) Node {
	return Elem("h2", c...)
}

func H3(c ...Node) Node {
	return Elem("h3", c...)
}

func H4(c ...Node) Node {
	return Elem("h4", c...)
}

func H5(c ...Node) Node {
	return Elem("h5", c...)
}

func H6(c ...Node) Node {
	return Elem("h6", c...)
}

func Head(c ...Node) Node {
	return Elem("head", c...)
}

func Header(c ...Node) Node {
	return Elem("header", c...)
}

func Hgroup(c ...Node) Node {
	return Elem("hgroup", c...)
}

func Hr(c ...Node) Node {
	return ElemEmpty("hr", c...)
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
	return Elem("i", c...)
}

func Iframe(c ...Node) Node {
	return ElemEmpty("iframe", c...)
}

func Img(c ...Node) Node {
	return ElemEmpty("img", c...)
}

func Input(c ...Node) Node {
	return ElemEmpty("input", c...)
}

func Ins(c ...Node) Node {
	return Elem("ins", c...)
}

func Kbd(c ...Node) Node {
	return Elem("kbd", c...)
}

func Label(c ...Node) Node {
	return Elem("label", c...)
}

func Legend(c ...Node) Node {
	return Elem("legend", c...)
}

func Li(c ...Node) Node {
	return Elem("li", c...)
}

func Link(c ...Node) Node {
	return ElemEmpty("link", c...)
}

func Main(c ...Node) Node {
	return Elem("main", c...)
}

func Map(c ...Node) Node {
	return Elem("map", c...)
}

func Mark(c ...Node) Node {
	return Elem("mark", c...)
}

func Menu(c ...Node) Node {
	return Elem("menu", c...)
}

func Meta(c ...Node) Node {
	return ElemEmpty("meta", c...)
}

func Meter(c ...Node) Node {
	return Elem("meter", c...)
}

func Nav(c ...Node) Node {
	return Elem("nav", c...)
}

func Noscript(c ...Node) Node {
	return Elem("noscript", c...)
}

func Object(c ...Node) Node {
	return Elem("object", c...)
}

func Ol(c ...Node) Node {
	return Elem("ol", c...)
}

func Optgroup(c ...Node) Node {
	return Elem("optgroup", c...)
}

func Option(c ...Node) Node {
	return Elem("option", c...)
}

func Output(c ...Node) Node {
	return Elem("output", c...)
}

func P(c ...Node) Node {
	return Elem("p", c...)
}

func Param(c ...Node) Node {
	return ElemEmpty("param", c...)
}

func Picture(c ...Node) Node {
	return Elem("picture", c...)
}

func Pre(c ...Node) Node {
	return Elem("pre", c...)
}

func Progress(c ...Node) Node {
	return Elem("progress", c...)
}

func Q(c ...Node) Node {
	return Elem("q", c...)
}

func Rp(c ...Node) Node {
	return Elem("rp", c...)
}

func Rt(c ...Node) Node {
	return Elem("rt", c...)
}

func Ruby(c ...Node) Node {
	return Elem("ruby", c...)
}

func S(c ...Node) Node {
	return Elem("s", c...)
}

func Samp(c ...Node) Node {
	return Elem("samp", c...)
}

func Script(c ...Node) Node {
	return Elem("script", c...)
}

func Search(c ...Node) Node {
	return Elem("search", c...)
}

func Section(c ...Node) Node {
	return Elem("section", c...)
}

func Select(c ...Node) Node {
	return Elem("select", c...)
}

func Small(c ...Node) Node {
	return Elem("small", c...)
}

func Source(c ...Node) Node {
	return ElemEmpty("source", c...)
}

func Span(c ...Node) Node {
	return Elem("span", c...)
}

func Strong(c ...Node) Node {
	return Elem("strong", c...)
}

func Style(c ...Node) Node {
	return Elem("style", c...)
}

func Sub(c ...Node) Node {
	return Elem("sub", c...)
}

func Summary(c ...Node) Node {
	return Elem("summary", c...)
}

func Sup(c ...Node) Node {
	return Elem("sup", c...)
}

func Svg(c ...Node) Node {
	return Elem("svg", c...)
}

func Table(c ...Node) Node {
	return Elem("table", c...)
}

func Tbody(c ...Node) Node {
	return Elem("tbody", c...)
}

func Td(c ...Node) Node {
	return Elem("td", c...)
}

func Template(c ...Node) Node {
	return Elem("template", c...)
}

func Textarea(c ...Node) Node {
	return Elem("textarea", c...)
}

func Tfoot(c ...Node) Node {
	return Elem("tfoot", c...)
}

func Th(c ...Node) Node {
	return Elem("th", c...)
}

func Thead(c ...Node) Node {
	return Elem("thead", c...)
}

func Time(c ...Node) Node {
	return Elem("time", c...)
}

func Title(title string) Node {
	return Elem("title", Text(title))
}

func Tr(c ...Node) Node {
	return Elem("tr", c...)
}

func Track(c ...Node) Node {
	return ElemEmpty("track", c...)
}

func U(c ...Node) Node {
	return Elem("u", c...)
}

func Ul(c ...Node) Node {
	return Elem("ul", c...)
}

func Var(c ...Node) Node {
	return Elem("var", c...)
}

func Video(c ...Node) Node {
	return Elem("video", c...)
}

func Wbr(c ...Node) Node {
	return Elem("wbr", c...)
}
