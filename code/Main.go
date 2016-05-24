// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-16

package main

import (
	"fmt"
	"inda-project/code/crawler"
)

// A hash map containing the visited URLs
var visited = make(map[string]bool)

// A hash map containing the URLs to be visited
var toVisit = make(map[string]bool)

// The URL to start from
var startURL = "http://en.wikipedia.org/wiki/Main_Page"

// The main function, starts the crawler on a
// given URL. TODO: add print to a file for all
// visited URLs.
func main() {
	m := crawler.Crawl(startURL)
	fmt.Println(len(m))
	for key, _ := range m {
		fmt.Println(key)
	}
}
