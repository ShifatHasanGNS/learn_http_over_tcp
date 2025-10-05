package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

const file_messages = "assets/messages.txt"

func getLinesChannel(f io.ReadCloser) <-chan string {
	out := make(chan string, 1)

	go func() {
		defer f.Close()
		defer close(out)

		str := ""
		for {
			data := make([]byte, 8)
			_, err := f.Read(data)
			if err != nil {
				break
			}

			if i := bytes.IndexByte(data, '\n'); i != -1 { // got a newline
				str += string(data[:i])
				out <- str
				str = ""
				data = data[i+1:] // remaining data after newline
			}
			str += string(data)
		}
	}()

	return out
}

func main() {
	f, err := os.Open(file_messages)
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesChannel(f)

	for line := range lines {
		fmt.Println(line)
	}
}
