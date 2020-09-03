package main

import "fmt"

func main() {
	fmt.Println(invisible(10))
}

func invisible(n int) (v int) {
	defer func() {
		if p := recover(); p != nil {
			v, _ = p.(int)
		}
	}()

	panic(n)
}
