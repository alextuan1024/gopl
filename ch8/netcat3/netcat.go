package main

import (
	"io"
	"log"
	"net"
	"os"
)

// NOTE: 使用channel同步两个goroutine;等待后台goroutine
func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn) // NOTE: ignoring errors
		log.Println("done")
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin) // NOTE: 由于调用的io.Copy，所以只会在读取出错时返回，否则阻塞
	conn.Close()
	<-done // waite for background goroutine to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
