package main

import (
	"fmt"
	"gopl.io/ch5/links"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var host, home string

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "findlinks: usage: findlinks first...")
		os.Exit(1)
	}
	first := os.Args[1]
	firstUrl, err := url.Parse(first)
	if err != nil {
		log.Fatalf("findlinks: parse url failed, error:%v", err)
	}
	host = firstUrl.Host
	home, err = ioutil.TempDir("", "crawldata")
	if err != nil {
		log.Fatalf("findlinks: create temp dir failed, error: %v", err)
	}
	breadthFirst(crawl, []string{first})
}

// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(raw string) []string {
	fmt.Println(raw)
	save(home, raw)
	list, err := links.Extract(raw)
	if err != nil {
		log.Print(err)
	}
	return list
}

func save(home string, raw string) {
	u, err := url.Parse(raw)
	if err != nil {
		log.Printf("parse url: %s failed, error: %v\n", raw, err)
		return
	}
	if u.Host == host {
		res, err := http.Head(u.String())
		if err != nil {
			log.Printf("head url: %s failed, error: %v\n", u.String(), err)
			return
		}
		// save html only
		ctype := res.Header.Get("Content-Type")
		if strings.Contains(ctype, "text/html") {
			// save
			dir := filepath.Join(home, u.Host, u.Path)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				log.Printf("create directory: %s failed, error: %v\n",
					dir, err)
				return
			}
			abs := filepath.Join(dir, "file.html")
			f, err := os.OpenFile(abs, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				log.Printf("create file failed, error: %v\n", err)
				return
			}
			defer f.Close()
			resp, err := http.Get(u.String())
			if err != nil {
				log.Printf("fetch from url: %s, error: %v\n", u.String(), err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Printf("read response failed: error: %v\n", err)
				return
			}
			if _, err := f.Write(body); err != nil {
				log.Printf("write response into file failed: %v\n", err)
				return
			}
		}
	}
}
