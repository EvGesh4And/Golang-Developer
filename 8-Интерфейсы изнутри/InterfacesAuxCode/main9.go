package main

import (
	"fmt"
	"net/http"
)

type IHTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

func main() {
	var c IHTTPClient
	fmt.Println("value of client is", c)
	fmt.Printf("type of client is %T\n", c)
	fmt.Println("(c == nil) is", c == nil)
}
