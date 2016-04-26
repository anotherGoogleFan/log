package log

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

func BadRequest(w http.ResponseWriter, req *http.Request) *HTTPError {
	return &HTTPError{
		w:          w,
		req:        req,
		statusCode: http.StatusBadRequest,
		logIP:      true,
	}
}

func Forbidden(w http.ResponseWriter, req *http.Request) *HTTPError {
	return &HTTPError{
		w:          w,
		req:        req,
		statusCode: http.StatusForbidden,
		logIP:      true,
	}
}

func ServerError(w http.ResponseWriter, req *http.Request) *HTTPError {
	return &HTTPError{
		w:          w,
		req:        req,
		statusCode: http.StatusInternalServerError,
	}
}

func NotFound(w http.ResponseWriter, req *http.Request) *HTTPError {
	return &HTTPError{
		w:          w,
		req:        req,
		statusCode: http.StatusNotFound,
	}
}

type HTTPError struct {
	w          http.ResponseWriter
	req        *http.Request
	statusCode int
	logIP      bool
}

func (e *HTTPError) fields(v ...interface{}) Fields {
	var msg interface{} = v
	if len(v) == 1 {
		msg = v[0]
	}
	type Error struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	}
	e.w.WriteHeader(e.statusCode)
	json.NewEncoder(e.w).Encode(Error{
		Status:  fmt.Sprintf("%d %s", e.statusCode, http.StatusText(e.statusCode)),
		Message: fmt.Sprint(msg),
	})
	fields := Fields{
		"status":  e.statusCode,
		"request": e.req.RequestURI,
	}
	if e.logIP {
		fields["ip"] = realIP(e.req)
	}
	return fields
}

func realIP(req *http.Request) string {
	ip, _, _ := net.SplitHostPort(req.RemoteAddr)
	if realIP := req.Header.Get("X-Real-IP"); realIP != "" {
		ip = realIP
	}
	return ip
}

func (e *HTTPError) Warnf(format string, v ...interface{}) {
	e.fields(fmt.Sprintf(format, v...)).printf(fWarn, 1, format, v)
}

func (e *HTTPError) Errorf(format string, v ...interface{}) {
	e.fields(fmt.Sprintf(format, v...)).printf(fError, 1, format, v)
}

func (e *HTTPError) Infof(format string, v ...interface{}) {
	e.fields(fmt.Sprintf(format, v...)).printf(fInfo, 1, format, v)
}

func (e *HTTPError) Warn(v ...interface{}) {
	e.fields(v).print(fWarn, 1, v)
}

func (e *HTTPError) Error(v ...interface{}) {
	e.fields(v).print(fError, 1, v)
}

func (e *HTTPError) Info(v ...interface{}) {
	e.fields(v).print(fInfo, 1, v)
}
