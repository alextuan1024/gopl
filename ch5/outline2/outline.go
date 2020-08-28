package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

var depth int

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "not enough arguments. usage: outline url...")
		os.Exit(1)
	}
	for _, url := range os.Args[1:] {
		res, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline: cannot get from url: %s, error: %v\n", url, err)
			continue
		}
		defer res.Body.Close()
		doc, err := html.Parse(res.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "outline2: cannot parse html: %v\n", err)
			continue
		}
		forEachNode(doc, startElement, endElement)
	}
}

// forEachNode calls the funcitons pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional
// pre is called before the children are visited (preorder) and
// post is called after (postorder)
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

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}
