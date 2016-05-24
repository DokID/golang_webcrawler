// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-16

package main

import (
	"container/list"
	"fmt"
	"inda-project/code/crawler"
	"sync"
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
	toVisit[startURL]  = true
	var wg sync.WaitGroup

	for {
		for urlkey, exists := range toVisit {
			wg.Add(1)
			go func() {
				if exists == true {
					toVisit[urlkey] = false
					visited[urlkey] = true
					tList := crawler.Crawl(urlkey)
					controller(tList)
				}
				wg.Done()
			}()
		}
	}
	wg.Wait()
	for urlkey, exists := range visited {
		if exists {
			fmt.Println(urlkey)
		}
	}
}

func controller(list List) {
	for i := range list {
		if toVisit[i] == false && visited[i] == false {
			toVisist[i] = true
		}
	}
}
