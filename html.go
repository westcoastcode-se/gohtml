package gohtml

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

func Html(c ...Node) RootNode {
	return func(w io.Writer) (int, error) {
		we := &errorAwareWriter{
			w: w,
		}
		if _, err := we.Write([]byte("<!doctype html><html")); err != nil {
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

func Body(c ...Node) Node {
	return Elem("body", c...)
}

func Head(c ...Node) Node {
	return Elem("head", c...)
}

func Title(title string) Node {
	return Elem("title", Text(title))
}

func Expand(c ...Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, cc := range c {
			b = cc(b, w)
		}
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

func Meta(c ...Node) Node {
	return ElemEmpty("meta", c...)
}

func Link(c ...Node) Node {
	return ElemEmpty("link", c...)
}

func Script(c ...Node) Node {
	return Elem("script", c...)
}

func Style(c ...Node) Node {
	return Elem("style", c...)
}

func Div(c ...Node) Node {
	return Elem("div", c...)
}

func Br(c ...Node) Node {
	return ElemEmpty("br", c...)
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

func Table(c ...Node) Node {
	return Elem("table", c...)
}

func Tr(c ...Node) Node {
	return Elem("tr", c...)
}

func Td(c ...Node) Node {
	return Elem("td", c...)
}

func Th(c ...Node) Node {
	return Elem("th", c...)
}

func THead(c ...Node) Node {
	return Elem("thead", c...)
}

func TBody(c ...Node) Node {
	return Elem("tbody", c...)
}

func Header(c ...Node) Node {
	return Elem("header", c...)
}

func Main(c ...Node) Node {
	return Elem("main", c...)
}

func Button(c ...Node) Node {
	return Elem("button", c...)
}

func Ul(c ...Node) Node {
	return Elem("ul", c...)
}

func Li(c ...Node) Node {
	return Elem("li", c...)
}

func A(c ...Node) Node {
	return Elem("a", c...)
}

func P(c ...Node) Node {
	return Elem("p", c...)
}

// Bytes writes the supplied bytes as if it's a single text-block
func Bytes(value []byte) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write(value)
		return 0
	}
}

// Text creates a simple text string
func Text(text string) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte(text))
		return 0
	}
}

// Textf creates a formatted text string
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

func Span(c ...Node) Node {
	return Elem("span", c...)
}
