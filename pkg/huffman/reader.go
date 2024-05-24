package huffman

import "io"

type Reader struct {
	r io.Reader
}

func NewReader(r io.Reader) io.ReadCloser {
	return &Reader{r}
}

func (c *Reader) Read(p []byte) (n int, err error) {
	return c.r.Read(p)
}

func (c *Reader) Close() error {
	// TODO: finish reading the compressed data?
	return nil
}
