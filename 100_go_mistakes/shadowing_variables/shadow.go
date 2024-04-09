package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	var tracing bool
	if time.Now().Hour() > 12 {
		tracing = true
	}
	var client *http.Client
	if tracing {
		// Shadows client
		client, err := createClientWithTracing()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(client)
	} else {
		// Shadow again
		client, err := createDefaultClient()
		if err != nil {
			log.Fatal(err)
		}
		log.Println(client)
	}
	_ = client // must call since was shadowed to avoid compilation errors
	avoidShadowing()
}

func avoidShadowing() {
	var tracing bool
	if time.Now().Hour() > 12 {
		tracing = true
	}
	var client *http.Client
	var err error
	if tracing {
		client, err = createClientWithTracing() // No shadow
	} else {
		client, err = createDefaultClient() // No shadow
	}
	if err != nil {
		log.Fatal(err)
	}
	log.Println(client) // Never shadows client variable
}

func createClientWithTracing() (*http.Client, error) {
	return nil, nil
}

func createDefaultClient() (*http.Client, error) {
	return nil, nil
}
