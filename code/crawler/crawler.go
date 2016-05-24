// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-19

package crawler

import (
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Crawl calls and parses the assigned webpage, looking for URLs and collecting
// them in a list.
func Crawl(url string) map[string]bool {
	body := getPage(url)
	defer body.Close()
	m := make(map[string]bool)
	tokenizer := html.NewTokenizer(body)

	for {
		tt := tokenizer.Next()

		switch tt {
		case html.ErrorToken:
			// Error, if not EOF, preserve
			if tokenizer.Err() != io.EOF {
				log.Printf("Token Error: %v", tokenizer.Err())
			}
			// EOF, no more links to be found
			log.Println("Error: EOF")
			return m
		case html.StartTagToken:
			// Start tag has been located
			token := tokenizer.Token()
			// Check if token is an <a> tag
			if token.Data != "a" {
				continue
			}

			log.Println("Token found!")
			// Look for href in token
			found, href := getHref(token)
			if !found {
				continue
			}
			// Check if href is relevant
			if strings.Index(href, "/wiki/") == 0 {
				m[href] = true
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
