// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-19

package crawler

import (
	"bytes"
	"container/list"
	"fmt"
	"log"
	"net/http"
)

// Crawl calls and parses the assigned webpage, looking for URLs and collecting
// them in a list.
func Crawl(url string) *list.List {
	body := getPage(url)
	fmt.Print(body)
	return &list.List{}
}

// Helper function used to retrieve the contents of the specified webpage.
func getPage(url string) *string {
	// Retrieve the page
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer resp.Body.Close()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// Convert io.ReadCloser to string
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	body := buf.String()
	return &body
}
