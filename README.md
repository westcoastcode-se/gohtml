A low-latency HTML components renderer using Go functions

This is, somewhat, under development. Mostly for my amusement at the moment

## Examples

* [Hello World ](examples/hello_world/main.go) - A very simple hello world example
* [Using Arrays](examples/arrays/main.go) - A slightly more complex example in which we emit nodes using arrays
* [Using Channels](examples/channels/main.go) - Sometimes using external systems is needed. This shows how we can
  emit HTML nodes from a channel
* [Caching](examples/caching/main.go) - Emitting HTML nodes from external systems is sometimes slow. Caching such nodes
  might be needed
