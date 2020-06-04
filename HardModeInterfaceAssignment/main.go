package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("No filename argument")
		return
	}

	filename := os.Args[1]
	file, _ := os.Open(filename)
	io.Copy(os.Stdout, file)
}
