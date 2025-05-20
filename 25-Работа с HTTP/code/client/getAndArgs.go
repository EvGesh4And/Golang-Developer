package main

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

// https://site.ru/search?query=...&limit=
func main() {
	reqArgs := url.Values{}
	reqArgs.Add("query", "go syntax")
	reqArgs.Add("limit", "5")

	reqUrl, _ := url.Parse("http://site.ru/search")
	reqUrl.RawQuery = reqArgs.Encode()

	req, _ := http.NewRequest("GET", reqUrl.String(), nil)

	req.Header.Add("User-Agent", `Mozilla/5.0 Gecko/20100101 Firefox/39.0`)
	log.Printf("Тело запроса:\n%+v\n", req)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Тело ответа:\n%s", string(body))
}
