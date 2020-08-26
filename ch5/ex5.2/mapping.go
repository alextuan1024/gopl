package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex5.2: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	for k, v := range count(m, doc) {
		fmt.Printf(" %v: %v\n", k, v)
	}
}

func count(m map[string]int, n *html.Node) map[string]int {
	for ; n != nil; n = n.NextSibling {
		m = count(m, n.FirstChild)

		if n.Type == html.ElementNode {
			v, ok := m[n.Data]
			if !ok {
				m[n.Data] = 1
			} else {
				m[n.Data] = v + 1
			}
		}
	}
	return m
}
