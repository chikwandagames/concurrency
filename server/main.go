package main

import (
	"io"
	"log"
	"net"
	"time"
)

// To run
// 1. cd server
// 2. ./server

func main() {
	// TODO: write server program to handle concurrent client connections.
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			// If theres an error we continue to get the next client connection
			continue
		}
		// To serve multiple clients, simply turn the call to handleConn in goroutine
		// so that we create a seperate goroutine to handle each client connection
		go handleConn(conn)
	}

}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		// Write a response to the connection handler every second
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
