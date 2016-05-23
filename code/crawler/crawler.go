// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-19

package crawler

import (
	"container/list"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Crawl calls and parses the assigned webpage, looking for URLs and collecting
// them in a list.
func Crawl(url string) *list.List {
	body := getPage(url)
	defer body.Close()
	list := new(list.List)
	list.Init()
	tokenizer := html.NewTokenizer(body)
	log.Println("Tokenizer created!")

	for {
		tt := tokenizer.Next()

		log.Printf("%v", tt)

		switch tt {
		case html.ErrorToken:
			if tokenizer.Err() != io.EOF {
				log.Printf("Token Error: %v", tokenizer.Err())
			}
			log.Println("Error: EOF")
			return list
		case html.StartTagToken:
			token := tokenizer.Token()
			if token.Data != "a" {
				continue
			}

			log.Println("Token found!")

			found, href := getHref(token)
			if !found {
				continue
			}

			if strings.Index(href, "/wiki/") == 0 {
				list.PushFront(href)
				log.Println("URL located!")
			}

		}
	}
}

// Helper function used to retrieve the contents of the specified webpage.
func getPage(url string) (body io.ReadCloser) {
	// Retrieve the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("HTTP Error: %v", err)
	}
	if err != nil {
		log.Fatalf("Response Error: %v", err)
	}
	body = resp.Body
	log.Println("Body received!")
	return
}

// Helper function used to identify href:s in an html file
func getHref(t html.Token) (found bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			found = true
			log.Println("Href found!")
		}
	}
	return
}
