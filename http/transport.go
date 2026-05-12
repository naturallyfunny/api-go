package http

import (
	"context"
	nethttp "net/http"
)

// Propagator reads a value from ctx and sets a corresponding HTTP header.
// It must be a no-op when the value is absent from ctx.
type Propagator func(ctx context.Context, h nethttp.Header)

// WithHeader returns a Propagator that reads the string value stored under key
// in ctx and sets it under headerName. It is a no-op when the value is absent
// or is not a non-empty string.
func WithHeader(key any, headerName string) Propagator {
	return func(ctx context.Context, h nethttp.Header) {
		if val, ok := ctx.Value(key).(string); ok && val != "" {
			h.Set(headerName, val)
		}
	}
}

// Transport is a net/http.RoundTripper that runs a set of Propagators on each
// outbound request before delegating to Base.
type Transport struct {
	// Base is the underlying RoundTripper; net/http.DefaultTransport is used when nil.
	Base nethttp.RoundTripper
	// Propagators is the ordered list of propagators applied to each request.
	Propagators []Propagator
}

func (t *Transport) RoundTrip(req *nethttp.Request) (*nethttp.Response, error) {
	ctx := req.Context()
	clone := req.Clone(ctx)

	for _, p := range t.Propagators {
		p(ctx, clone.Header)
	}

	return t.base().RoundTrip(clone)
}

func (t *Transport) base() nethttp.RoundTripper {
	if t.Base != nil {
		return t.Base
	}
	return nethttp.DefaultTransport
}
