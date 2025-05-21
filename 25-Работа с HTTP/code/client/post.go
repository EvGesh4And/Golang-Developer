package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

var (
	errUnexpectedHTTPStatus  = errors.New("errUnexpectedHTTPStatus")
	errUnexpectedContentType = errors.New("errUnexpectedContentType")
)

type AddRequest struct {
	Id    int    `json: "id"`
	Title string `json: "title"`
	Text  string `json: "text"`
}

func main() {
	addReq := &AddRequest{
		Id:    123,
		Title: "for loop",
		Text:  "...",
	}

	jsonBody, err := json.Marshal(addReq)
	if err != nil {
		log.Fatal(err)
	}
	req, err := http.NewRequest("POST", "http://site.ru/add_item", bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Fatal(fmt.Errorf("do request: %w", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		log.Fatal(fmt.Errorf("%w: %s", errUnexpectedHTTPStatus, resp.Status))
	}
	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		log.Fatal(fmt.Errorf("%w: %s", errUnexpectedContentType, ct))
	}
	body, err := io.ReadAll(resp.Body)
	fmt.Println(body)
}
