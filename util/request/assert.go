package request

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
)

// HttpRequestConfig http request config
type HttpRequestConfig struct {
	Method  string
	URL     string
	Param   io.Reader
	Header  map[string]string
	Context context.Context
}

// AssertHttpRequest check http request
func AssertHttpRequest(handler http.Handler, config HttpRequestConfig) *httptest.ResponseRecorder {
	r := httptest.NewRequest(config.Method, config.URL, config.Param)
	w := httptest.NewRecorder()

	if config.Header != nil {
		for k, v := range config.Header {
			r.Header.Set(k, v)
		}
	}

	if config.Context != nil {
		r.WithContext(config.Context)
	}

	handler.ServeHTTP(w, r)

	return w
}
