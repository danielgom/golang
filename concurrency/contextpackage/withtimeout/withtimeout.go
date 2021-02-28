package main

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	req, err := http.NewRequest("GET", "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond) // If the server does not respond within 100
	// milliseconds , then it is going to throw an error

	defer cancel()

	resp, err := http.DefaultClient.Do(req.WithContext(ctx))
	if err != nil {
		log.Println("ERROR", err)
		return
	}

	// Write the response to stdout

	io.Copy(os.Stdout, resp.Body)

	// Close the response body on the return

	err = resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

}
