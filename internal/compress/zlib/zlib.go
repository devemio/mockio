package zlib

import (
	"bytes"
	"compress/zlib"
	"io"
)

func Decompress(v string) string {
	r, _ := zlib.NewReader(bytes.NewReader([]byte(v)))

	b, _ := io.ReadAll(r)

	return string(b)
}
