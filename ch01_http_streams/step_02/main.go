package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const file_messages = "assets/messages.txt"

func main() {
	f, err := os.Open(file_messages)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 8)
	for {
		n, err := f.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		if n == 0 {
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buf[:n])
	}
}
