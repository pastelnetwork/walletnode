package log

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/pastelnetwork/go-commons/log"
	"github.com/pastelnetwork/go-commons/log/hooks"
	"github.com/pastelnetwork/go-commons/random"

	httpmiddleware "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// Log logs incoming HTTP requests and outgoing responses.
// It uses the request ID set by the RequestID middleware or creates a short unique request ID if missing for each incoming request
// and logs it with the request and corresponding response details.
func Log() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqID := r.Context().Value(middleware.RequestIDKey)
			if reqID == nil {
				reqID, _ = random.String(8, random.Base62Chars)
			}
			started := time.Now()

			WithField("from", logFrom(r)).
				WithField("req", r.Method+" "+r.URL.String()).
				Debugf("[%v] Request", reqID)

			rw := httpmiddleware.CaptureResponse(w)
			h.ServeHTTP(rw, r)

			WithField("status", rw.StatusCode).
				WithField("bytes", rw.ContentLength).
				WithField("time", time.Since(started).String()).
				Debugf("[%v] Response", reqID)
		})
	}
}

// logFrom makes a best effort to compute the request client IP.
func logFrom(req *http.Request) string {
	if f := req.Header.Get("X-Forwarded-For"); f != "" {
		return f
	}
	f := req.RemoteAddr
	ip, _, err := net.SplitHostPort(f)
	if err != nil {
		return f
	}
	return ip
}

// ErrorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func ErrorHandler(ctx context.Context, w http.ResponseWriter, err error) {
	id := ctx.Value(middleware.RequestIDKey).(string)
	_, _ = w.Write([]byte("[" + id + "] encoding: " + err.Error()))
	Errorf("[%s] %s", id, err.Error())
}

func init() {
	log.AddHook(hooks.NewContextHook(middleware.RequestIDKey, func(entry *log.Entry, ctxValue interface{}) {
		entry.Message = fmt.Sprintf("[%v] %s", ctxValue, entry.Message)
	}))
}
