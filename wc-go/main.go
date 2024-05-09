package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var filePath string

	if len(os.Args) > 1 {
		filePath = strings.Join(os.Args[1:], " ")
		} else {
		fmt.Println("Path: ", os.Args[1])
		filePath = os.Args[1]
	}

	if len(filePath) == 0 {
		fmt.Println("Please input a place name")
		return
	}
	file, err := os.Open(filePath);

	panicOnErr(err)
	
	var lineCount, wordCount, charCount, byteCount = 0, 0, 0, 0

	scanner := bufio.NewScanner(file)

	line := ""

	for scanner.Scan() {
		lineCount++
		line = scanner.Text()
		charCount += len(line)
		byteCount += len([]byte(line))
		words := strings.Split(line, " ")
		wordCount += len(words)
	}
	
	fmt.Println(lineCount, wordCount, charCount, byteCount)

}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}