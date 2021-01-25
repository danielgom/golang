package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {

	// Connect to server on localhost 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	mustCopy(os.Stdout, conn)
	err = conn.Close()
	if err != nil {
		log.Fatal(err)
	}

}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
