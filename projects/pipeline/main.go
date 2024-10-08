package main

import "fmt"

func removeDuplicates(inputStream, outputStream chan string) {
	var str string
	for value := range inputStream {
		if value != str {
			str = value
			outputStream <- value
		}
	}
	close(outputStream)
}

func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go removeDuplicates(inputStream, outputStream)

	go func() {
		defer close(inputStream)

		for _, r := range "112334456" {
			inputStream <- string(r)
		}
	}()

	for x := range outputStream {
		fmt.Print(x)
	}
	// 123456
}
