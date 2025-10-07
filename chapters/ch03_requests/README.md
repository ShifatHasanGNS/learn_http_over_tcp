# Note

## Resources

- [HTTP/1.1](https://en.wikipedia.org/wiki/HTTP)
- [RFC 9112 Section 2.1](https://datatracker.ietf.org/doc/html/rfc9112#name-message-format)

## Terminology

- RFC = Request for Comments
- CRLF = Carriage Return Line Feed = `\r\n`

## HTT-Message Format

```http
<start-line> CRLF
<header-field> CRLF
<header-field> CRLF
...
<header-field> CRLF
CRLF
<message-body>
```

A very simple but comprehensive example:

```http
GET /coffee HTTP/1.1\r\n
Host: localhost:42069\r\n
User-Agent: curl/8.7.1\r\n
Accept: */*\r\n
\r\n
{
    "message": "Hello, World!",
    "status": "success",
}
```
