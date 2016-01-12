package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type MarvelResponse struct {
	Data MarvelResponseData `json:"data"`
}

type MarvelResponseData struct {
	Results []interface{} `json:"results"`
	Count   json.Number   `json:"count"`
	Limit   json.Number   `json:"limit"`
	Offset  json.Number   `json:"offset"`
	Total   json.Number   `json:"total"`
}

type Character struct {
	Id     json.Number            `json:"id"`
	Name   string                 `json:"name"`
	Url    string                 `json:"resourceURI"`
	Series map[string]interface{} `json:"series"`
}

type Series struct {
}

var (
	publicKey  string = os.Getenv("MARVEL_PUBLIC")
	privateKey string = os.Getenv("MARVEL_PRIVATE")
)

func characters(payload map[string]interface{}) {
	ch := make(chan interface{})
	go get("characters", payload, ch)
	m := <-ch
	b, err := json.MarshalIndent(m, "", "  ")
	fmt.Println(string(b))
	if err != nil {
		log.Fatal("bad json", err)
	}
	var c []Character
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.UseNumber()
	if err = dec.Decode(&c); err != nil {
		log.Fatal("bad json", err)
	}
	fmt.Println("%#v", c[0])
}

func get(path string, payload map[string]interface{}, ch chan interface{}) {
	defer close(ch)
	ts := time.Now().Format("20060102150405")
	payload["ts"] = ts
	payload["apikey"] = publicKey

	md5 := md5.Sum([]byte(fmt.Sprintf("%s%s%s", ts, privateKey, publicKey)))
	payload["hash"] = fmt.Sprintf("%x", md5)

	var payloadSlice []string
	for key, value := range payload {
		payloadSlice = append(payloadSlice, fmt.Sprintf("%s=%s", key, value))
	}
	payloadString := strings.Join(payloadSlice, "&")

	url := fmt.Sprintf("http://gateway.marvel.com/v1/public/%s?%s", path, payloadString)
	log.Println(url)

	response, _ := http.Get(url)

	defer response.Body.Close()

	var res MarvelResponse
	dec := json.NewDecoder(response.Body)
	dec.UseNumber()
	if err := dec.Decode(&res); err != nil {
		log.Println(err)
		return
	}
	ch <- res.Data.Results
}

func main() {
	payload := map[string]interface{}{}
	characters(payload)
	// b, _ := json.MarshalIndent(m, "", "  ")
	// fmt.Print(string(b))
}
