package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo: Create file")
	fname := "./test.txt"
	f, err := os.Create(fname)
	if err != nil {
		fmt.Println("Received error")
		return
	}
	defer f.Close()
	f.Write([]byte("This is the test content of our file"))
}
