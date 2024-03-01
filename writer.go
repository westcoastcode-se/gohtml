package gohtml

import "io"

type ErrorAwareWriter struct {
	w   io.Writer
	len int
	err error
}

func (e *ErrorAwareWriter) Write(b []byte) (int, error) {
	if e.err != nil {
		return 0, e.err
	}
	var i int
	i, e.err = e.w.Write(b)
	e.len += i
	return i, e.err
}
