package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getLinks() {
	url := "https://www.facebook.com/"
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("Status error code %d %s", resp.StatusCode, resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	check(err)

	f := func(i int, s *goquery.Selection) bool {
		link, _ := s.Attr("href")
		return strings.HasPrefix(link, "http")
	}

	doc.Find("body a").FilterFunction(f).Each(func(_ int, tag *goquery.Selection) {

		link, _ := tag.Attr("href")
		linkText := tag.Text()
		fmt.Printf("%s: %s\n", linkText, link)

	})
}

func main() {
	getLinks()
}
