package main

import (
	"bufio"
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
	f, err := os.Open(file_messages)
	if err != nil {
		log.Fatal(err)
	}

	lines := getLinesChannel(f)

	for line := range lines {
		fmt.Println(line)
	}
}
