// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-19

package crawler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Crawls the given URL looking for more URLs
func crawl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatalf("Error: %v", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("Error: %v", err)
		log.Fatalf("Error: %v", err)
	}
	fmt.Printf("%s", body)
}
