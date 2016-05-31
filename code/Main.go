// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-25

package main

import (
	"fmt"
	"inda-project/code/crawler"
	"runtime"
	"sync"
	"time"
)

// The library struct hold a map with visited links,
// one with links to be visited as well as a Mutex
// lock.
type library struct {
	sync.Mutex
	visited map[string]bool
	toVisit map[string]bool
}

// Initialize library maps
var lib = &library{visited: map[string]bool{}, toVisit: map[string]bool{}}

// The URL to start from
var startURL = "http://en.wikipedia.org/wiki/Main_Page"

// The main function, starts the crawler on a
// given URL and prints out all of the visited URLs
func main() {
	lib.toVisit[startURL] = true
	wg := new(sync.WaitGroup)

	// The main for-loop responsible for sending crawlers to all known pages
	for {
		for urlkey, exists := range lib.toVisit {
			// A crawler thread starts
			fmt.Println(runtime.NumGoroutine())
			for runtime.NumGoroutine() >= 200 {
				time.Sleep(10 * time.Millisecond)
			}
			wg.Add(1)
			// A crawler thread starts
			go func(urlkey string, exists bool) {
				if exists {
					lib.Lock()
					delete(lib.toVisit, urlkey)
					lib.visited[urlkey] = true
					tMap := crawler.Crawl(urlkey)
					controller(tMap)
					lib.Unlock()
					runtime.Gosched()
				}
				wg.Done()
			}(urlkey, exists)
		}
		wg.Wait()
		// If no more pages left to visit, break
		if len(lib.toVisit) == 0 {
			break
		}
	}
	// When done crawling, print found URL:s
	for urlkey, exists := range lib.visited {
		if exists {
			fmt.Println(urlkey)
		}
	}
}

// Helper function responsible for adding newly acquired
// unvisited pages to the queue.
func controller(tMap map[string]bool) {
	for i := range tMap {
		if lib.visited[i] == false {
			lib.toVisit[i] = true
		}
	}
}
