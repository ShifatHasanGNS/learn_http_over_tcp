package main

import (
	"bufio"
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

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			out <- scanner.Text()
		}

		if err := scanner.Err(); err != nil {
			log.Println("Scanner error:", err)
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
