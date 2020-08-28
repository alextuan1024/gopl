package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "elementbyid usage: elementbyid url id")
		os.Exit(1)
	}
	url, id := os.Args[1], os.Args[2]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementbyid fetch from url: %s failed, error: %v", url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "elementbyid parse failed, error: %v", err)
		os.Exit(1)
	}
	n := ElementByID(doc, id)
	if n != nil {
		fmt.Printf(`<%s`, n.Data)
		for _, a := range n.Attr {
			fmt.Printf(` %s="%s"`, a.Key, a.Val)
		}
		if n.FirstChild != nil {
			fmt.Printf("></%s>\n", n.Data)
		} else {
			fmt.Printf("/>\n")
		}
		return
	}
	fmt.Printf("no element with id as: %s", id)

}

func ElementByID(doc *html.Node, id string) *html.Node {
	var target *html.Node
	forEachNode(doc, &target, id, detectElement, nil)
	return target
}

func forEachNode(n *html.Node, targ **html.Node, id string, pre, post func(*html.Node, **html.Node, string) bool) {
	if pre != nil {
		if pre(n, targ, id) {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, targ, id, pre, post)
	}
	if post != nil {
		if post(n, targ, id) {
			return
		}
	}
}

func detectElement(n *html.Node, t **html.Node, id string) (stop bool) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				*t = n
				stop = true
				return
			}
		}
	}
	return
}
