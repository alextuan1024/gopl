package main

import (
	"bytes"
	"fmt"
)

func main() {
	var old = "Hal: $foo\nBarry: '$foo', are you serious?\nHal: 'Course I'm serious. $foo\nBarry: You're an idiot, Hal.\n"
	result := expand(old, replaceStr)
	fmt.Println(result)
}

func expand(s string, f func(string) string) string {
	substr := "$foo"
	res := []byte(s)
	subBytes := []byte(substr)
	rep := f(substr[1:])
	var i = bytes.Index(res, subBytes)
	for ; i != -1; {
		res = append(res[:i], append([]byte(rep), res[i+len(subBytes):]...)...)
		i = bytes.Index(res, subBytes)
	}
	return string(res)
}

func replaceStr(src string) string {
	if src == "foo" {
		return `Who the hell is Batman?`
	}
	return ""
}
