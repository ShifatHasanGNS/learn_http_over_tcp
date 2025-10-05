# Task

- In one Terminal (server), run in `localhost` a TCP server that listens on port `42069`
  - open a terminal just inside the repository
  - `go run ./ch_03/step_01`
- From another Terminal (client), send a HTTP Request to `http://localhost:42069/coffee` using `curl`
  - open another terminal in any location on the same device
  - `curl http://localhost:42069/coffee`

---

For:

```shell
curl -X POST -H "Content-Type: application/json" -d '{"flavor":"dark mode"}' http://localhost:42069/coffee
```

Output:

```http
POST /coffee HTTP/1.1
Host: localhost:42069
User-Agent: curl/8.7.1
Accept: */*
Content-Type: application/json
Content-Length: 24

```

You might notice that the body `{"flavor":"dark mode"}` isn't showing up yet in the tcplistener output... and that's because it doesn't end with a newline.

So now, For:

```shell
curl -X POST \
     -H "Content-Type: application/json" \
     -d '{"flavor":"dark mode"}
     ' \
     http://localhost:42069/coffee
```

Output:

```http
POST /coffee HTTP/1.1
Host: localhost:42069
User-Agent: curl/8.7.1
Accept: */*
Content-Type: application/json
Content-Length: 24

{"flavor":"dark mode"}

```

Now it should work as expected!
