package queue

import (
	"fmt"
	"net/http"
)

type QueueResponseWriter struct {
	body       []byte
	statusCode int
	header     http.Header
}

func NewQueueResponseWriter() *QueueResponseWriter {
	return &QueueResponseWriter{
		header: http.Header{},
	}
}

func (w *QueueResponseWriter) Header() http.Header {
	return w.header
}

func (w *QueueResponseWriter) Write(b []byte) (int, error) {
	w.body = b
	// implement it as per your requirement
	return 0, nil
}

func (w *QueueResponseWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
}

var okFn = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	r := &http.Request{
		Method: http.MethodPost,
	}
	w := NewQueueResponseWriter()
	okFn(w, r)
	fmt.Println(w.statusCode)
}
