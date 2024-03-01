package gohtml

import (
	"fmt"
	"io"
	"log"
)

type RootNode func(w io.Writer) (int, error)
type Node func(b byte, w io.Writer) byte

func Attribute(key string, value string) Node {
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

func Html(lang string, c ...Node) RootNode {
	return func(w io.Writer) (int, error) {
		we := &ErrorAwareWriter{
			w: w,
		}
		if _, err := we.Write([]byte("<!doctype html><html lang=\"" + lang + "\">")); err != nil {
			return 0, err
		}
		for _, cc := range c {
			cc(0, we)
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

func Text(text string) Node {
	return func(b byte, w io.Writer) byte {
		if b != 0 {
			_, _ = w.Write([]byte{b})
		}
		_, _ = w.Write([]byte(text))
		return 0
	}
}

func Textf(format string, a ...any) Node {
	return func(b byte, w io.Writer) byte {
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

// EmitArray emits an array of items and converts them into nodes to be written
func EmitArray[T any](arr []T, emit func(t T) Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, item := range arr {
			b = emit(item)(b, w)
		}
		return b
	}
}
