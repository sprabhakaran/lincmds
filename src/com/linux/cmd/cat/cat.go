package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		panic("Missing arguments")
	}
	for _, file := range os.Args[1:] {

	}
}
