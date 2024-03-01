package main

import (
	"bytes"
	"fmt"
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
)

func Emit5Rows() Node {
	// the array containing the actual data
	items := []int{1, 2, 3, 4, 5}

	// emit array will take the supplied array call the supplied function
	// for each value in the array. The function itself is expected to return a Node
	return EmitArray(items[:], func(val int) Node {
		return Tr(
			Td(Textf("%d", val)),
		)
	})
}

func ArrayOf5Rows() Node {
	// the array containing the actual data
	items := []int{1, 2, 3, 4, 5}

	// creates an array of nodes to be returned
	var nodes []Node
	for _, i := range items {
		nodes = append(nodes, Tr(
			Td(Textf("%d", i)),
		))
	}

	// expand the array into a single Node return statement
	return ExpandArray(nodes)
}

func main() {
	b := bytes.Buffer{}

	// generate the actual html
	numBytes, err := Html("sv",
		Head(
			// Add a meta header tag with the attribute charset="UTF-8"
			Meta(a.Charset("UTF-8")),
			Title("My Title"),
		),
		Body(
			H1(
				Text("Table using emit"),
			),
			Table(
				Emit5Rows(),
			),
			H1(
				Text("Table with arrays"),
			),
			Table(
				ArrayOf5Rows(),
			),
		),
	)(&b)

	// write the result in the console
	fmt.Println("written", numBytes, "bytes with error", err)
	fmt.Println(b.String())
}
