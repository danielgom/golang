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

	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx,
		http.MethodGet, "https://www.google.com", nil)

	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.DefaultClient.Do(req)
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
