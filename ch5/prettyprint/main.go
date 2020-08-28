package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
)

var depth int
var w io.Writer

func main() {
	if len(os.Args) != 2 {
		_, _ = fmt.Fprintf(os.Stderr, "prettyprint: usage: prettyprint url")
		os.Exit(1)
	}
	url := os.Args[1]
	res, err := http.Get(url)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "prettyprint: fetch from url: %s error: %v", url, err)
		os.Exit(1)
	}
	defer res.Body.Close()
	doc, err := html.Parse(res.Body)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "prettyprint: parse html failed: %v", err)
		os.Exit(1)
	}
	w = os.Stdout
	forEachNode(doc, startTag, endTag)
}

func forEachNode(n *html.Node, pre, post func(*html.Node)) {
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

func startTag(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(w, ` %s="%s"`, a.Key, a.Val)
		}
		if n.FirstChild == nil { // has no child at all
			fmt.Fprintf(w, "/>\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
	case html.TextNode:
		if n.Data != "" {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", n.Data)
		}
	case html.CommentNode:
		fmt.Fprintf(w, "%*s<!--%s-->", depth*2, "", n.Data)
	}
	depth++
}

func endTag(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		depth--
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	case html.TextNode:
		fallthrough
	case html.CommentNode:
		depth--
	}
}
