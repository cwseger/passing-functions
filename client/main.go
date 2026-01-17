package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	// Connect to server
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Open and transfer source file
	file, err := os.Open("source/source.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(conn, file)
	if err != nil {
		log.Fatal(err)
	}
}
