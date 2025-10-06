package main

import (
	"bytes"
	"fmt"
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

	str := ""
	for {
		data := make([]byte, 8)
		_, err := f.Read(data)
		if err != nil {
			break
		}

		if i := bytes.IndexByte(data, '\n'); i != -1 { // got a newline
			str += string(data[:i])
			fmt.Println(str)
			str = ""
			data = data[i+1:] // remaining data after newline
		}
		str += string(data)
	}
}
