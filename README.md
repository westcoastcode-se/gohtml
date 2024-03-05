# gohtml

A functional low-latency HTML components renderer. All HTML tags are prefixed with `h` and all attributes are prefixed
with `a`. If a HTML tag is missing, then you can use the generic `h. 

## What isn't this framework?

This framework is a functional library built using generic functions. It allows for combining nodes in a tree-like
structure using function calls. Though, as a side effect of the design of this framework, if you want to put attributes
on the HTML tag then you have to add those attributes before the first child-node.

## Examples

### Hello World

This example prints out the html content of a simple Hello World HTML into the console.

```go
package main

import (
	"bytes"
	"fmt"
	"github.com/westcoastcode-se/gohtml/a"
	"github.com/westcoastcode-se/gohtml/h"
)

func main() {
	b := bytes.Buffer{}

	numBytes, err := h.Html(a.Lang("en"),
		h.Head(
			h.Meta(a.Charset("UTF-8")),
			h.Title("My Title"),
		),
		h.Body(
			h.H1(
				h.Text("Hello World"),
			),
		),
	)(&b)

	fmt.Println("written", numBytes, "bytes with error", err)
	fmt.Println(b.String())
}
```

* [Hello World ](examples/hello_world/main.go) - A very simple hello world example
* [Using Arrays](examples/arrays/main.go) - A slightly more complex example in which we emit nodes using arrays
* [Using Channels](examples/channels/main.go) - Sometimes using external systems is needed. This shows how we can
  emit HTML nodes from a channel
* [Caching](examples/caching/main.go) - Emitting HTML nodes from external systems is sometimes slow. Caching such nodes
  might be needed
* [CDN](examples/cdn/main.go) - Another caching example that simulates a more accurate real-world scenario. This
  use-case downloads javascript and css files from a CDN and injects it directly in the HTML. We are, then, caching the
  javascript- and css content
* [Extension](examples/extension/main.go) - If you want to create complex html structures, then it sometimes makes sense
  to create structs that defines how html elements are connected together. This example shows how we can use structures
  in order to build a form with multiple input fields in a more controlled manner
