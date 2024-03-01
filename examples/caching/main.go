package main

import (
	"bytes"
	"fmt"
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
	"time"
)

func SlowIO(count int) chan Node {
	ch := make(chan Node)
	go func() {
		for i := 0; i < count; i++ {
			ch <- Div(Text(fmt.Sprintf("value: %d", i)))
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

func PrintHtml() {
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
				// Cache the result in the supplied cache storage. The result will be cached for 10 seconds and any
				// child-nodes will not be called
				Cache(cacheStorage, "mykey", 10*time.Second,
					SuperSlowNode(),
				),
			),
		),
	)(&b)

	// write the result in the console
	fmt.Println("written", numBytes, "bytes with error", err)
	fmt.Println(b.String())
}

func main() {
	// Call the print function
	PrintHtml()

	// The second call is much faster because of the caching mechanism
	PrintHtml()
}
