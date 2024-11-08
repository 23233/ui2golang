package driver

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net"
	"net/http"
	"time"
)

// Requester defines the structure for an HTTP request.
type Requester struct {
	Url     string            // The URL to send the request to.
	Method  string            // The HTTP method to use (e.g., GET, POST).
	Headers map[string]string // A map of headers to include in the request.
	Data    interface{}       // The data to send with the request, if any.
}

var (
	// dialer is a custom net.Dialer with specific timeout and resolver settings.
	dialer = &net.Dialer{
		Timeout:   30 * time.Second, // Connection timeout duration.
		KeepAlive: 30 * time.Second, // Keep-alive period for the connection.
		Resolver: &net.Resolver{
			PreferGo: true, // Prefer Go's built-in DNS resolver.
			Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
				d := net.Dialer{}
				return d.DialContext(ctx, "udp", "114.114.114.114:53") // Custom DNS server.
			},
		},
	}
	// transport is an HTTP transport that uses the custom dialer.
	transport = &http.Transport{
		DialContext: dialer.DialContext,
	}
)

// Request sends an HTTP request based on the provided Requester options.
// Parameters:
//   - opt: A pointer to a Requester struct containing request details.
//
// Returns:
//   - map[string]any: A map containing the response data if successful.
//   - error: An error object if the request fails.
func Request(opt *Requester) (map[string]any, error) {
	var res = make(map[string]any)

	var data []byte
	if opt.Data != nil {
		jsonData, err := json.Marshal(opt.Data)
		if err == nil {
			data = jsonData
		}
	}

	req, err := http.NewRequest(opt.Method, opt.Url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	if opt.Headers == nil {
		opt.Headers = make(map[string]string)
	}

	// Set default content type if not provided.
	if _, ok := opt.Headers["Content-Type"]; !ok {
		opt.Headers["Content-Type"] = "application/json"
	}

	if opt.Headers != nil {
		for k, v := range opt.Headers {
			req.Header.Set(k, v)
		}
	}

	// Create an HTTP client with the custom transport.
	client := &http.Client{
		Transport: transport,
	}

	// Send the HTTP request.
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(body, &res); err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, errors.New(resp.Status)
}
