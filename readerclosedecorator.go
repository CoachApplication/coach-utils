package utils

import "io"

/**
 * We use io.ReadCloser extensively in the config Connector system, but often we have just an io.Reader to work with.
 * This struct just allows decorating the Reader with a `func Close() error` and turning it into an io.ReadCloser
 *
 * This process is pretty simple, so I didn't write any tests for it.
 */

// ReaderCloserDecorator A simple struct that wraps an io.Reader and adds a Close() to make it a io.ReadCloser
type ReaderCloserDecorator struct {
	io.Reader
	closeFunc func() error
}

// DecorateReader Decorate a Reader with a Close Method
func CloseDecorateReader(r io.Reader, f func() error) io.ReadCloser {
	return io.ReadCloser(ReaderCloserDecorator{
		Reader:    r,
		closeFunc: f,
	})
}

// Close the io.Closer interface method
func (rcw ReaderCloserDecorator) Close() error {
	if rcw.closeFunc == nil {
		return nil
	} else {
		return rcw.closeFunc()
	}
}
