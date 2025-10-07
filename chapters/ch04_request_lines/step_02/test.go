package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRequestLineParse(t *testing.T) {
	r, err := RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err, "error should be nil")
	require.NotNil(t, r, "request should not be nil")
	assert.Equal(t, "GET", r.RequestLine.Method, "method should be GET")
	assert.Equal(t, "/", r.RequestLine.RequestTarget, "request target should be /")
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion, "http version should be 1.1")

	// Test: Good GET Request line with path
	r, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/2.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.NoError(t, err, "error should be nil")
	require.NotNil(t, r, "request should not be nil")
	assert.Equal(t, "GET", r.RequestLine.Method, "method should be GET")
	assert.Equal(t, "/coffee", r.RequestLine.RequestTarget, "request target should be /coffee")
	assert.Equal(t, "1.1", r.RequestLine.HttpVersion, "http version should be 1.1")

	// Test: Invalid number of parts in request line
	_, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
	require.Error(t, err, "error should not be nil")
}
