package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		words, images, err := CountWordsAndImages(url)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "countWordsAndImages: %v\n", err)
			continue
		}
		fmt.Printf("%d words, %d images\n", words, images)
	}
}

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	doc, err := html.Parse(resp.Body)
	if err != nil {
		err = fmt.Errorf("parsing HTML: %w", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	switch n.Type {
	case html.TextNode:
		c, err := countWordsInStr(n.Data)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "encountered an error: %v\n", err)
		}
		words += c
	case html.ElementNode:
		if n.Data == "img" {
			images += 1
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return words, images
}

func countWordsInStr(s string) (int, error) {
	var count int
	input := bufio.NewScanner(strings.NewReader(s))
	input.Split(bufio.ScanWords)
	for input.Scan() {
		count++
	}
	if input.Err() != nil {
		return count, fmt.Errorf("countWordsInStr: %w", input.Err())
	}
	return count, nil
}
