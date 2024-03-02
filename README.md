# gohtml

A functional low-latency HTML components renderer

This is, somewhat, under development. Mostly for my amusement at the moment

## Examples

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

## What isn't this framework?

This framework is a functional library with very generic functions. It allows for combining nodes with other nodes. But
since the framework isn't creating any structures then it can't really know if the same attribute is added twice. It
also can't know if you are mixing node- and attribute nodes. This is because there are no way for Golang to know if a
function is an attribute builder or a node builder. This is why attributes must be added first, and then child nodes

The intended purpose of this is, however, to be built upon to create [extensions](examples/extension/main.go). Those can
be used to enforce types in a structured manner.
