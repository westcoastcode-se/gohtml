package gohtml

import (
	"io"
	"time"
)

// EmitChannel calls a function be called in which a channel of nodes is consumed until the channel is closed
func EmitChannel(f func() chan Node) Node {
	return func(b byte, w io.Writer) byte {
		ch := f()
		for n := range ch {
			b = n(b, w)
		}
		return b
	}
}

type ChannelProps struct {
	// Timeout until we give up on reading data from the channel
	Timeout time.Duration
}

// EmitChannelEx is an extended version of EmitChannel in which you can configure how the framework should
// handle the processing of the channel
func EmitChannelEx(f func() chan Node, props ChannelProps) Node {
	// Normal emit if no timeout is required
	if props.Timeout == time.Duration(0) {
		return EmitChannel(f)
	}

	timeout := time.After(props.Timeout)
	return func(b byte, w io.Writer) byte {
		// call the supplied function so that we can get a channel to consume
		ch := f()
	rangeLoop:
		for {
			select {
			case <-timeout:
				break rangeLoop
			case out, ok := <-ch:
				if !ok {
					break rangeLoop
				}
				b = out(b, w)
			}
		}
		return b
	}
}

// EmitArray emits an array of items and converts them into nodes to be written
func EmitArray[T any](arr []T, emit func(t T) Node) Node {
	return func(b byte, w io.Writer) byte {
		for _, item := range arr {
			b = emit(item)(b, w)
		}
		return b
	}
}
