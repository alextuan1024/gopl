package main

import (
	"fmt"
	"github.com/alextuan1024/gopl/ch5/links"
	"log"
	"os"
)

var (
	slots = make(chan struct{}, 20)
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: crawl url")
	}
	var seen = make(map[string]bool)
	var worklist = make(chan []string)
	var n int
	n++
	go func() {
		worklist <- os.Args[1:]
	}()
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(lk string) {
					worklist <- crawl(lk)
				}(link)
			}
		}
	}
}

func crawl(url string) (urls []string) {
	fmt.Println(url)
	slots <- struct{}{} // acquire a token
	urls, err := links.Extract(url)
	if err != nil {
		fmt.Print(err)
	}
	<-slots // release the token
	return
}
