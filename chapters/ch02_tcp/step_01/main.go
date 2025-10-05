package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

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
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Server is running on --> http://localhost:42069")

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		lines := getLinesChannel(connection)

		for line := range lines {
			fmt.Println(line)
		}
	}
}
