package huffman

import (
	"io"
)

type writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) io.WriteCloser {
	return &writer{w}
}

func (w *writer) Write(p []byte) (int, error) {
	c := compress(p)
	return w.w.Write(c)
}

func (w *writer) Close() error {
	// TODO: finish writing the compressed data. Maybe replace it with Flush()?
	return nil
}
