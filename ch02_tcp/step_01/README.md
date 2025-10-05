# Task

- In one Terminal (server), run in `localhost` a TCP server that listens on port `42069`
  - open a terminal just inside the repository
  - `go run ./ch_02/step_01`
- From another Terminal (client), send string "Hello, World!" to `http://localhost:42069`
  - open another terminal in any location on the same device
  - `printf "Hello World...\r\n" | nc -w 1 localhost 42069`
