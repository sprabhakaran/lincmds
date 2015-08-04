package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func exploreFolder(dirPath string) {
	files, _ := ioutil.ReadDir(dirPath)
	fmt.Println(dirPath, files)
	for i := range files {
		file := files[i]

		if file.IsDir() {
			fmt.Println("Dir :: ", file.Name())
			exploreFolder(dirPath + "/" + file.Name())
		} else {
			fmt.Println("\t", file.Name())
		}
	}
}

func main() {
	dstPath := os.Args[1]

	exploreFolder(dstPath)
}
