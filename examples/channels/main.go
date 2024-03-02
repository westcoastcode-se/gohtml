package main

import (
	"github.com/westcoastcode-se/gohtml/a"
	. "github.com/westcoastcode-se/gohtml/h"
	"log"
	"net/http"
	"time"
)

func SlowIO(args int) chan Node {
	// Create a channel that's responsible for emitting Nodes. Each emit is simulated to take 100 milliseconds
	ch := make(chan Node)
	go func() {
		for i := 0; i < args; i++ {
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
				// Emitting SlowIO(10) takes 1 second. Consider looking into the Caching example on
				// how to handle this type of slow IO
				EmitChannel(func() chan Node {
					return SlowIO(10)
				}),
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
