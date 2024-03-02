package gohtml

import "io"

// errorAwareWriter gives us a way of keeping track of all memory being written to the writer and
// a simple way of centralizing the handling of any error that might occur while writing the HTML content
type errorAwareWriter struct {
	w   io.Writer
	len int
	err error
}

func (e *errorAwareWriter) Write(b []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var i int
	i, e.err = e.w.Write(b)
	e.len += i
	return i, e.err
}
