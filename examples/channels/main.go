package main

import (
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
	"log"
	"net/http"
	"time"
)

func SimulateSlowIO() chan Node {
	// Create a channel that's responsible for emitting Nodes. We are simulating slow IO by using Sleep
	// Consider looking into the Caching example on how to handle slow IO.
	ch := make(chan Node)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- Tr(
				Td(
					Textf("value: %d", i),
				),
			)
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()
	return ch
}

func SimulateSlowIO2() func() chan Node {
	return func() chan Node {
		return SimulateSlowIO()
	}
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
				EmitChannel(SimulateSlowIO2()),
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
