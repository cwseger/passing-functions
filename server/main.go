package main

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)

func main() {
	// Start TCP server
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Listening on :9000")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	// Create file and copy contents to it
	out, err := os.Create("received/source/source.go")
	if err != nil {
		log.Fatal(err)
	}
	// defer out.Close()

	_, err = io.Copy(out, conn)
	if err != nil {
		log.Fatal(err)
	}
	out.Close()

	if err := os.Chdir("received"); err != nil {
		log.Fatal(err)
	}

	// Run the go program contained in the transferred file
	cmd := exec.Command("go", "run", "main.go")

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	if err := os.Chdir("../"); err != nil {
		log.Fatal(err)
	}
}
