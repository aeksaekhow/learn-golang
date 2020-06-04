package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://google.com")

	m := memoryBuffer{}
	fmt.Println("Address of m", &m)
	io.Copy(&m, resp.Body)
	//fmt.Println(m.message)

}

type memoryBuffer struct {
	message string
}

func (c *memoryBuffer) Write(p []byte) (n int, err error) {
	c.message += string(p)
	return len(p), nil
}
