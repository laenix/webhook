package lark

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func SendMessage(apiUrl, reqdata, time string) {
	client := &http.Client{}
	var req *http.Request
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(reqdata))
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("contentType", "application/json")
	req.Header.Add("Accept", "application/json, text/plain, */*")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	html, _ := goquery.NewDocumentFromReader(resp.Body)
	fmt.Println(html.Text())
}
