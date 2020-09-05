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
	// if local is "", path.Base() returns a "."
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	defer func() {
		// modify err if any error happened while closing f
		if cErr := f.Close(); cErr != nil {
			err = cErr
		}
	}()
	n, err = io.Copy(f, resp.Body)
	return local, n, err
}
