package main

import (
	"fmt"
	"io/ioutil"
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
		showFileContent(file)
	}
}

func showFileContent(file *os.File) {
	fmt.Println("\nFile : ", file.Name(), "\n", "\b--------------------------------------------------")
	bt, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bt))
}

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		return os.IsNotExist(err) == false
	}
	return true
}
