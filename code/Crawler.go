// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-16

package main

import (
  "fmt"
  "bufio"
)

// A hashmap containing the visited URLs
var visited = make(map[string]int)

// A hashmap containing the visited URLs
var toVisit = make(map[string]int)

// The main function, starts the crawler on a
// given URL. TODO: add print to a file for all
// visited URLs.
func main() {
}

// Crawls the given URL looking for more URLs
func crawl(URL string) {
}
