# Note

Parsing the Request Line (string)

For example, given:

```http
POST /coffee HTTP/1.1
Host: localhost:42069
User-Agent: curl/7.81.0
Accept: */*
Content-Length: 21

{"flavor":"dark mode"}
```

We want our HTTP parser to return a struct that looks like this:

```go
type Request struct {
    RequestLine RequestLine
    Headers     map[string]string
    Body        []byte
}
```

That way our application logic (the server code) has nicely structured HTTP data to work with.

Our goal is to take the server we created in the last section and have it parse out the start-line according to the RFC here.

The Request-Line
Remember how HTTP messages start with a start-line? Well, if it's a request (not a response), then the start-line is called the request-line and has a specific format.

```text
HTTP-name    = %s"HTTP"
HTTP-version = HTTP-name "/" DIGIT "." DIGIT
request-line = method SP request-target SP HTTP-version
```

Example:

```text
HTTP-name    = GET /coffee HTTP/1.1
HTTP-version = HTTP/1.1
```

---

## Task

Add the following structs and the function signature to `parser.go`:

```go
type RequestLine struct {
    Method        string
    RequestTarget string
    HttpVersion   string
}

type Request struct {
    RequestLine RequestLine
}

// ...
// Add what you need here
// ...

// to parse the request-line from the reader
func RequestFromReader(reader io.Reader) (*Request, error)
```

Add the following tests to `test.go`:

```go
// Test: Good GET Request line
r, err := RequestFromReader(strings.NewReader("GET / HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
require.NoError(t, err)
require.NotNil(t, r)
assert.Equal(t, "GET", r.RequestLine.Method)
assert.Equal(t, "/", r.RequestLine.RequestTarget)
assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

// Test: Good GET Request line with path
r, err = RequestFromReader(strings.NewReader("GET /coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
require.NoError(t, err)
require.NotNil(t, r)
assert.Equal(t, "GET", r.RequestLine.Method)
assert.Equal(t, "/coffee", r.RequestLine.RequestTarget)
assert.Equal(t, "1.1", r.RequestLine.HttpVersion)

// Test: Invalid number of parts in request line
_, err = RequestFromReader(strings.NewReader("/coffee HTTP/1.1\r\nHost: localhost:42069\r\nUser-Agent: curl/7.81.0\r\nAccept: */*\r\n\r\n"))
require.Error(t, err)
```

For now,

- you can push the entire request into memory using io.ReadAll and work with the entire thing as a string.
- Create a parseRequestLine function to do the parsing.
- Remember that newlines in HTTP are \r\n, not just \n.
- You can discard everything that comes after the request-line for now.
- There are always just 3 parts to the request line: strings.Split is your friend here.
- Verify that the "method" part only contains capital alphabetic characters.
- Verify that the http version part is 1.1, extracted from the literal HTTP/1.1 format, as we only support HTTP/1.1 for now.

Add more test cases to `test.go` to cover any edge cases you can think of. Here are the names of all the tests I wrote:

- Good Request line
- Good Request line with path
- Good POST Request with path
- Invalid number of parts in request line
- Invalid method (out of order) Request line
- Invalid version in Request line
