package main

import (
	"bufio"
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

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
