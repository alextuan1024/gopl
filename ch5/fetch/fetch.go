package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: fetch url")
	}
	url := os.Args[1]
	filename, size, err := fetch(url)
	if err != nil {
		log.Printf("fetch from url: %s error: %v\n", url, err)
		os.Exit(1)
	}
	log.Printf("saved url: %s as file: %s, size: %v", url, filename, size)
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	n, err = io.Copy(f, resp.Body)
	// close file, but prefer error from copy, if any
	if cErr := f.Close(); cErr != nil {
		err = cErr
	}
	return local, n, err
}
