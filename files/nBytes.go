package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// Simple read only open. We will cover actually reading
	// and writing to files in examples further down the page
	file, err := os.Open("rez.go")
	if err != nil {
		fmt.Println(err)
	}
	byteSlice := make([]byte, 18)
	numBytesRead, err := io.ReadFull(file, byteSlice)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(byteSlice))
	fmt.Println(numBytesRead)
	file.Close()
}
