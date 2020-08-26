package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.3: %v\n", err)
		os.Exit(1)
	}
	printText(doc)
}

func printText(n *html.Node) {
	for ; n != nil; n = n.NextSibling {
		if n.Data != "script" && n.Data != "style" {
			printText(n.FirstChild)
		}
		if n.Type == html.TextNode && n.Data != "" {
			fmt.Printf("Text:%+v\n", n.Data)
		}
	}
}
