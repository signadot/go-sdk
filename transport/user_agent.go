package transport

import (
	"net/http"
)

type userAgent struct {
	inner http.RoundTripper
	agent string
}

func (ua *userAgent) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", ua.agent)
	return ua.inner.RoundTrip(r)
}
