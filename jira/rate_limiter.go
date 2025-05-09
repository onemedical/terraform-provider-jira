package jira

import (
	"net/http"
	"time"

	"golang.org/x/time/rate"
)

// RateLimitedTransport wraps an http.RoundTripper and limits the number of requests
// that can be made per second.
type RateLimitedTransport struct {
	transport http.RoundTripper
	limiter   *rate.Limiter
}

// NewRateLimitedTransport creates a new rate-limited transport with the specified
// requests per second limit.
func NewRateLimitedTransport(transport http.RoundTripper, rps float64) *RateLimitedTransport {
	return &RateLimitedTransport{
		transport: transport,
		limiter:   rate.NewLimiter(rate.Limit(rps), 1), // Allow bursts of 1 request
	}
}

// RoundTrip implements the http.RoundTripper interface.
func (t *RateLimitedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Wait for rate limiter
	if err := t.limiter.Wait(req.Context()); err != nil {
		return nil, err
	}
	
	// Forward the request to the underlying transport
	return t.transport.RoundTrip(req)
}

// Client returns an *http.Client that uses the RateLimitedTransport.
func (t *RateLimitedTransport) Client() *http.Client {
	return &http.Client{
		Transport: t,
		Timeout:   time.Minute, // Set a reasonable timeout
	}
}

// WrapClient wraps an existing http.Client with rate limiting.
func WrapClientWithRateLimiter(client *http.Client, rps float64) *http.Client {
	if client == nil {
		client = http.DefaultClient
	}
	
	transport := client.Transport
	if transport == nil {
		transport = http.DefaultTransport
	}
	
	client.Transport = NewRateLimitedTransport(transport, rps)
	return client
}