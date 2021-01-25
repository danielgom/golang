package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Write sever program to handle concurrent connections
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			continue
		}
		// To accept multiple client calls we just add go to our connection handler
		go handleCon(conn)
	}
}

func handleCon(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n ")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}
