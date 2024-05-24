package huffman

import (
	"fmt"
	"io"
)

type writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) io.WriteCloser {
	return &writer{w}
}

func (w *writer) Write(p []byte) (int, error) {
	c, err := compress(p)
	if err != nil {
		return 0, fmt.Errorf("failed to compress data: %w", err)
	}
	return w.w.Write(c)
}

func (w *writer) Close() error {
	// TODO: finish writing the compressed data. Maybe replace it with Flush()?
	return nil
}
