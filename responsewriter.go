package violetear

import (
	"net/http"
	"time"
)

// ResponseWriter wraps the standard http.ResponseWriter
type ResponseWriter struct {
	http.ResponseWriter
	requestID    string
	size, status int
	start        time.Time
}

// NewResponseWriter returns ResponseWriter
func NewResponseWriter(w http.ResponseWriter, rid string) *ResponseWriter {
	rw := &ResponseWriter{
		ResponseWriter: w,
		start:          time.Now(),
	}
	if rid != "" {
		rw.requestID = w.Header().Get(rid)
		rw.Header().Set(rid, rw.requestID)
	}
	return rw
}

// Status provides an easy way to retrieve the status code
func (w *ResponseWriter) Status() int {
	return w.status
}

// Size provides an easy way to retrieve the response size in bytes
func (w *ResponseWriter) Size() int {
	return w.size
}

// Start retrieve the start time
func (w *ResponseWriter) RequestTime() string {
	return time.Since(w.start).String()
}

// RequestID retrieve the Request ID
func (w *ResponseWriter) RequestID() string {
	return w.requestID
}

// Write satisfies the http.ResponseWriter interface and
// captures data written, in bytes
func (w *ResponseWriter) Write(data []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	size, err := w.ResponseWriter.Write(data)
	w.size += size
	return size, err
}

// WriteHeader satisfies the http.ResponseWriter interface and
// allows us to catch the status code
func (w *ResponseWriter) WriteHeader(statusCode int) {
	w.status = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}
