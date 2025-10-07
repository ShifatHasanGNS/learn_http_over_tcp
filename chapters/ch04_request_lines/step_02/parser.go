package main

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type RequestLine struct {
	Method        string
	RequestTarget string
	HttpVersion   string
}

type Request struct {
	RequestLine RequestLine
}

var (
	ErrorInvalidRequestLine     = fmt.Errorf("invalid request line")
	ErrorUnsupportedHTTPVersion = fmt.Errorf("unsupported http version; only HTTP/1.1 is supported")
	Separator                   = "\r\n"
)

func (r *RequestLine) isValid() bool {
	return r.HttpVersion == "1.1"
}

func parseRequestLine(line string) (*RequestLine, string, error) {
	idx := strings.Index(line, Separator)
	if idx == -1 {
		return nil, "", ErrorInvalidRequestLine
	}

	request_line_parts := strings.SplitN(line[:idx], " ", 3)
	if len(request_line_parts) != 3 {
		return nil, line, ErrorInvalidRequestLine
	}

	remaining_request := line[idx+len(Separator):]
	request_line := &RequestLine{
		Method:        request_line_parts[0],
		RequestTarget: request_line_parts[1],
		HttpVersion:   strings.TrimPrefix(request_line_parts[2], "HTTP/"),
	}
	if !request_line.isValid() {
		return nil, line, ErrorUnsupportedHTTPVersion
	}

	return request_line, remaining_request, nil
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, errors.Join(
			fmt.Errorf("failed to do `io.ReadAll`: %w", err),
			err,
		)
	}

	request_line, _, err := parseRequestLine(string(data))
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *request_line,
	}, nil
}
