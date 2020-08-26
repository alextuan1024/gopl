package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	freq, err := wordfreq(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "wordfreq: %v", err)
		os.Exit(1)
	}
	for k, v := range freq {
		fmt.Printf("%s: %d\n", k, v)
	}
}

func wordfreq(in io.Reader) (map[string]int, error) {
	freq := make(map[string]int)
	input := bufio.NewScanner(in)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		freq[word]++
	}
	if input.Err() != nil {
		return nil, input.Err()
	}

	return freq, nil
}
