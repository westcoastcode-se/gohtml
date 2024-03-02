package main

import (
	"github.com/westcoastcode-se/gohtml/a"
	. "github.com/westcoastcode-se/gohtml/h"
	"log"
	"net/http"
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

func index(w http.ResponseWriter, r *http.Request) {
	// generate the actual html
	_, _ = Html(a.Lang("en"),
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
	)(w)
}

func main() {
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
