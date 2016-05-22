// Gábor Nagy and Niklas Ingemar Bergdahl 2016-05-19

package crawler

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Crawls the given URL looking for more URLs
func Crawl(url string) *list.List {
	body := getPage(url)
	return &list.List{}
}

func getPage(url string) *[]byte {
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
	log.Printf("%s", body)
	return &body
}
