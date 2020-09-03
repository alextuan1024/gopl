package main

import (
	"log"
	"time"
)

func main() {
	bigSlowFunction()
}

func bigSlowFunction() {
	defer trace("bigSlowFunction")()
	time.Sleep(time.Second * 10)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}
