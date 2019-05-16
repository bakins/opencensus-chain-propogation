package chain

import (
	"net/http"

	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
)

// Chain tries multiple trace propogation formats for input
// and can inject multiple outgoing propogation formats.
type Chain struct {
	Incoming []propagation.HTTPFormat
	Outgoing []propagation.HTTPFormat
}

// SpanContextFromRequest extracts a span context from incoming requests.
func (c *Chain) SpanContextFromRequest(r *http.Request) (trace.SpanContext, bool) {
	for _, f := range c.Incoming {
		s, ok := f.SpanContextFromRequest(r)
		if ok {
			return s, ok
		}
	}
	return trace.SpanContext{}, false
}

// SpanContextToRequest modifies the given request to include the given span context.
func (c *Chain) SpanContextToRequest(s trace.SpanContext, r *http.Request) {
	for _, f := range c.Outgoing {
		f.SpanContextToRequest(s, r)
	}
}
