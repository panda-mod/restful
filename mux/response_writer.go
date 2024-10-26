package mux

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"net/http"
)

type ResponseWriter struct {
	buffer     *bytes.Buffer
	writer     http.ResponseWriter
	statusCode int // 响应状态码
}

func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

func (w *ResponseWriter) Write(data []byte) (int, error) {
	return w.buffer.Write(data)
}

func (w *ResponseWriter) Flush() {
	if f, ok := w.writer.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *ResponseWriter) Push(target string, opts *http.PushOptions) {
	if f, ok := w.writer.(http.Pusher); ok {
		f.Push(target, opts)
	}
}

func (w *ResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if h, ok := w.writer.(http.Hijacker); ok {
		return h.Hijack()
	}
	return nil, nil, fmt.Errorf("http.Hijacker is not implemented")
}

func (w *ResponseWriter) Finally() {
	w.writer.WriteHeader(w.statusCode)
	w.writer.Write(w.buffer.Bytes())
}
