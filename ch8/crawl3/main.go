package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"os"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string) //de-duplicated urls

	go func() { worklist <- os.Args[1:] }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					worklist <- foundLinks
				}()
			}
		}()
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) (urls []string) {
	fmt.Println(url)
	urls, err := links.Extract(url)
	if err != nil {
		fmt.Print(err)
	}
	return urls
}
