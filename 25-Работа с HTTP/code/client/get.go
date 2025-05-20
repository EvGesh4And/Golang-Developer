package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("http://0.0.0.0:8083/api/analyzer/1747394522")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Print(resp.Status)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(body)
}
