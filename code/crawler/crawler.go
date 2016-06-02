// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-25

package crawler

import (
	"bytes"
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
			return m
		case html.StartTagToken:
			// Start tag has been located
			token := tokenizer.Token()
			// Check if token is an <a> tag
			if token.Data != "a" {
				continue
			}

			// Look for href in token
			found, href := getHref(token)
			if !found {
				continue
			}
			// Check if href is relevant
			if strings.Index(href, "/wiki/") == 0 {
				// Check if href is a hastag or special category
				if strings.Contains(href, "#") || strings.Contains(href, ":") {
					continue
				}
				// Append domain
				href = append(href)
				m[href] = true
			}
		}
	}
}

// Helper function used to retrieve the contents of the specified webpage.
func getPage(url string) (body io.ReadCloser) {
	// Retrieve the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Response Error: %v", err)
	}
	resp.Close = true
	body = resp.Body
	return
}

// Helper function used to identify href:s in an html file
func getHref(t html.Token) (found bool, href string) {
	for _, a := range t.Attr {
		if a.Key == "href" {
			href = a.Val
			found = true
		}
	}
	return
}

// Helper function used to append the start of the URL to the href
func append(href string) string {
	buffer := new(bytes.Buffer)
	buffer.WriteString("http://simple.wikipedia.org")
	buffer.WriteString(href)
	return buffer.String()
}
