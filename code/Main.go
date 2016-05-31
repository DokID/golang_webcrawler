// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-25

package main

import (
	"fmt"
	"inda-project/code/crawler"
	"log"
	"os"
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
	ch := make(chan string, 200)

	ch <- startURL

	// Start filewriter
	go save(ch)

	// The main for-loop responsible for sending crawlers to all known pages
	for {
		for urlkey := range ch {
			// A crawler thread starts
			for runtime.NumGoroutine() >= 200 {
				time.Sleep(20 * time.Millisecond)
			}
			wg.Add(1)
			// A crawler thread starts
			go func(urlkey string) {
				lib.Lock()
				delete(lib.toVisit, urlkey)
				lib.visited[urlkey] = true
				tMap := crawler.Crawl(urlkey)
				controller(tMap, ch)
				fmt.Println(len(lib.visited))
				lib.Unlock()
				runtime.Gosched()
				wg.Done()
			}(urlkey)
		}
		wg.Wait()
		// If no more pages left to visit, break
		if len(lib.toVisit) == 0 || len(lib.visited) >= 500000 {
			close(ch)
			break
		}
	}
}

// Helper function responsible for adding newly acquired
// unvisited pages to the queue.
func controller(tMap map[string]bool, ch chan<- string) {
	for i := range tMap {
		if lib.visited[i] == false {
			lib.toVisit[i] = true
			ch <- i
		}
	}
}

// Helper function used to flush the contents of the visited map to a
// CSV file.
func save(ch <-chan string) {
	file, err := os.Create("nodes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("pageId:ID,name,:LABEL\n")
	if err != nil {
		log.Fatal(err)
	}
	for key := range ch {
		_, err := file.WriteString(key + "," + key + ",WEBPAGE\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
