package main

import (
	"bytes"
	"fmt"
	. "github.com/westcoastcode-se/gohtml"
	"github.com/westcoastcode-se/gohtml/a"
	"time"
)

func SimulateSlowIO() chan Node {
	// Create a channel that's responsible for emitting Nodes. We are simulating slow IO by using Sleep
	ch := make(chan Node)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- Div(Text(fmt.Sprintf("value: %d", i)))
			time.Sleep(100 * time.Millisecond)
		}
		close(ch)
	}()
	return ch
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
				EmitChannel(SimulateSlowIO()),
			),
		),
	)(&b)

	// write the result in the console
	fmt.Println("written", numBytes, "bytes with error", err)
	fmt.Println(b.String())
}
