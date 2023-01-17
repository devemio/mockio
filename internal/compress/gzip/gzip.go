package gzip

import (
	"bytes"
	"compress/gzip"
	"io"
)

func Decompress(v string) string {
	r, _ := gzip.NewReader(bytes.NewReader([]byte(v)))

	b, _ := io.ReadAll(r)

	return string(b)
}
