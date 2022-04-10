package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	sl := flag.String("sl", "en", "source language")
	tl := flag.String("tl", "id", "translated language")
	q := flag.String("q", "words", "text to be translated")
	flag.Parse()

	postBody := url.Values{}
	postBody.Set("sl", *sl)
	postBody.Set("tl", *tl)
	postBody.Set("q", *q)

	url := "https://translate.google.com/translate_a/single?client=at&dt=t&dt=ld&dt=qca&dt=rm&dt=bd&dj=1&ie=UTF-8&oe=UTF-8&inputm=2&otf=2&iid=1dd3b944-fa62-4b55-b330-74909a99969e"
	contentType := "application/x-www-form-urlencoded"
	resp, err := http.Post(url, contentType, strings.NewReader(postBody.Encode()))
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var pot map[string]interface{}
	err = json.Unmarshal(body, &pot)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	sentence := pot["sentences"].([]interface{})[0].(map[string]interface{})
	fmt.Println(sentence["trans"])
}
