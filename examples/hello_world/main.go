package main

import (
	"bytes"
	"fmt"
	"github.com/westcoastcode-se/gohtml/a"
	. "github.com/westcoastcode-se/gohtml/h"
)

func main() {
	b := bytes.Buffer{}

	// generate the actual html
	numBytes, err := Html(a.Lang("en"),
		Head(
			// Add a meta header tag with the attribute charset="UTF-8"
			Meta(a.Charset("UTF-8")),
			Title("My Title"),
		),
		Body(
			H1(
				Text("Hello World"),
			),
		),
	)(&b)

	// write the result in the console
	fmt.Println("written", numBytes, "bytes with error", err)
	fmt.Println(b.String())
}
