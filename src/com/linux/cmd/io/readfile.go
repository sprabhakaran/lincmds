package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Give atleast one file")
	}
	filearg := os.Args[1]
	file, err := os.Open(filearg)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
	}
	// readLineByLine(file)
	readWordByWord(file)
	// readByteByByte(file)
}

func readLineByLine(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readWordByWord(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readByteByByte(file *os.File) {
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		fmt.Print(scanner.Text())
	}
}
