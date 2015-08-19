package main

import (
	"fmt"
	"os"
)

/*
	TODO: Support many command line options
*/

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("touch: missing file operand")
		return
	}
	filenames := os.Args[1:]
	for _, file := range filenames {
		os.Create(file)
	}
}
