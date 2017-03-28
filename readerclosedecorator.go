package utils

import "io"

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
