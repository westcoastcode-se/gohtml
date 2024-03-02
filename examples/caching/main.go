package main

import (
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
	"log"
	"net/http"
	"time"
)

func SlowIO(count int) chan Node {
	ch := make(chan Node)
	go func() {
		for i := 0; i < count; i++ {
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

func SuperSlowNode() Node {
	var count = 10
	return EmitChannel(func() chan Node {
		return SlowIO(count)
	})
}

// We are using a simple in-memory cache storage. This can be a session-associated cache storage instead, if
// you want to cache html based on the actual logged-in user
var cacheStorage = CreateInMemoryCacheStorage()

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
				// Cache the result in the supplied cache storage. The result will be cached for 10 seconds and any
				// child-nodes will not be called
				Cache(cacheStorage, "mykey", 10*time.Second,
					SuperSlowNode(),
				),
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
