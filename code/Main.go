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
	ch_count, ch_save := make(chan string, 20000000), make(chan string, 20000000)
	// Add starting URL to counting channel
	ch_count <- startURL
	// Start filewriter
	go save(ch_save)
	// The main for-loop responsible for sending crawlers to all known pages
	for urlkey := range ch_count {
		// Count running go-routines
		for runtime.NumGoroutine() >= 1000 {
			// If too many, sleep
			time.Sleep(100 * time.Millisecond)
		}
		// A crawler routine starts
		go func(urlkey string) {
			ch_save <- urlkey
			lib.Lock()
			delete(lib.toVisit, urlkey)
			lib.visited[urlkey] = true
			lib.Unlock()
			tMap := crawler.Crawl(urlkey)
			lib.Lock()
			controller(tMap, ch_count)
			fmt.Println(len(lib.visited), len(lib.toVisit), len(ch_count))
			lib.Unlock()
			runtime.Gosched()
		}(urlkey)
	}
	// Close all channels
	close(ch_count)
	close(ch_save)
}

// Helper function responsible for adding newly acquired
// unvisited pages to the queue.
func controller(tMap map[string]bool, ch_count chan<- string) {
	for i := range tMap {
		// Check if the address has been visited or is in queue
		if lib.visited[i] == false && lib.toVisit[i] == false {
			lib.toVisit[i] = true
			ch_count <- i
		}
	}
}

// Helper function used to flush the contents of the visited map to a
// CSV file.
func save(ch_save <-chan string) {
	file, err := os.Create("nodes.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	_, err = file.WriteString("pageId:ID,name,:LABEL\n")
	if err != nil {
		log.Fatal(err)
	}
	for key := range ch_save {
		_, err := file.WriteString(key + "," + key + ",WEBPAGE\n")
		if err != nil {
			log.Fatal(err)
		}
	}
}
