package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Missing File name arguments")
	}
	for _, file := range os.Args[1:] {
		file, err := os.Open(file)
		if err != nil {
			fmt.Println(err)
			continue
		}
		showFileContentAsReverse(file)
	}
}

func showFileContentAsReverse(file *os.File) {

}
