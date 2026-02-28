// This file contains the source the client will send to the server and then be executed by the server.
package source

import (
	"io"
	"log"
	"net/http"
)

func Function() {
	log.Println("--- Making http request ---")
	r, err := http.NewRequest(http.MethodGet, "https://www.google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Request created:", r.URL)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Println("Response received:", resp.Status)
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Response body length (bytes):", len(b))
}
