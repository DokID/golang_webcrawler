// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-16

package main

import (
	"fmt"
	"inda-project/code/crawler"
	"runtime"
	"sync"
)

// A hash map containing the visited URL:s
var visited = make(map[string]bool)

// A hash map containing the URLs to be visited
var toVisit = make(map[string]bool)

// The URL to start from
var startURL = "http://en.wikipedia.org/wiki/Main_Page"

// The main function, starts the crawler on a
// given URL. TODO: add print to a file for all
// visited URL:s.
func main() {
	toVisit[startURL] = true
	wg1, wg2 := new(sync.WaitGroup), new(sync.WaitGroup)
	mutex := new(sync.Mutex)

	// The main for-loop responsible for sending crawlers to all known pages
	for {
		for urlkey, exists := range toVisit {
			wg1.Add(1)
			wg2.Add(1)
			// A crawler thread starts
			go func(urlkey string, exists bool) {
				if exists {
					mutex.Lock()
					delete(toVisit, urlkey)
					visited[urlkey] = true
					tMap := crawler.Crawl(urlkey)
					controller(tMap)
					mutex.Unlock()
					runtime.Gosched()
				}
				wg1.Done()
				wg2.Done()
			}(urlkey, exists)
		}
		wg1.Wait()
		// If no more pages left to visit, break
		if len(toVisit) == 0 {
			break
		}
	}
	wg2.Wait()
	// When done crawling, print found URL:s
	for urlkey, exists := range visited {
		if exists {
			fmt.Println(urlkey)
		}
	}
}

// Helper function responsible for adding newly acquired
// unvisited pages to the queue.
func controller(tMap map[string]bool) {
	for i := range tMap {
		if visited[i] == false {
			toVisit[i] = true
		}
	}
}
