package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"strings"

	"github.com/devemio/mockio/internal/compress/gzip"
	"github.com/devemio/mockio/internal/compress/zlib"
	"github.com/devemio/mockio/internal/logger"
	"github.com/devemio/mockio/internal/types"
)

type Interceptor struct {
	w http.ResponseWriter
	r *http.Request
	l logger.Log
	c Color
	v types.Verbose
}

func (i *Interceptor) Handle(fn http.HandlerFunc) {
	if !i.v.IsVerbose() {
		fn(i.w, i.r)
		return
	}

	r, err := httputil.DumpRequest(i.r, true)
	if err != nil {
		i.l.Info(">>> failed to parse request: %v", err)
		return
	}

	i.l.Info(">>> %s %s\n%s", i.r.Method, i.r.URL.Path, r)

	rec := httptest.NewRecorder()

	fn(rec, i.r)

	i.l.Info("<<< %d %s %s%s\n%s", rec.Code, i.r.Method, i.r.URL.Path, i.headers(rec), i.decompress(rec))

	for k, v := range rec.Header() {
		i.w.Header()[k] = v
	}

	i.w.WriteHeader(rec.Code)

	_, _ = rec.Body.WriteTo(i.w)
}

func (i *Interceptor) headers(rec *httptest.ResponseRecorder) string {
	h := rec.Header()

	if i.v < types.VerbosityVeryVerbose || len(h) == 0 {
		return ""
	}

	var sb strings.Builder

	for k, values := range h {
		for _, v := range values {
			sb.WriteString(i.c.Cyan(k) + ": " + v + "\n")
		}
	}

	return "\n" + sb.String()
}

func (i *Interceptor) decompress(rec *httptest.ResponseRecorder) string {
	v := fmt.Sprintf("%v", rec.Body)
	e := rec.Header().Get("Content-Encoding")

	if e == "gzip" {
		return gzip.Decompress(v)
	}

	if e == "deflate" {
		return zlib.Decompress(v)
	}

	return v
}
