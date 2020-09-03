package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: elementsbytagname url tagnames...")
	}
	url := os.Args[1]
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("elementsbytagname fetch url: %s failed, error: %v", url, err)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatalf("elementsbytagname parse html error: %v", err)
	}
	names := os.Args[2:]
	nodes := ElementsByTagName(doc, names...)
	for _, node := range nodes {
		fmt.Printf("<%s", node.Data)
		for _, a := range node.Attr {
			fmt.Printf(` %s="%s"`, a.Key, a.Val)
		}
		fmt.Printf("></%s>\n", node.Data)
	}
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	var results []*html.Node
	var visit = func(n *html.Node) {
		for _, name := range names {
			if n.Type == html.ElementNode {
				if n.Data == name {
					results = append(results, n)
					break
				}
			}
		}
	}
	forEachNode(doc, visit, nil)
	return results
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
