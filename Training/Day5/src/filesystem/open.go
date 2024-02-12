package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Demo: Openfile")
	f, err := os.OpenFile("./test.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("received error while opening file")
		return
	}
	defer f.Close()
	/*b, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("Received error", err)
		return
	}*/
	b := make([]byte, 10, 10)
	n, err := f.Read(b)
	fmt.Println(n, err)
	for n > 0 && err == nil {
		fmt.Println("Read bytes ", n)
		fmt.Println(string(b))
		n, err = f.Read(b)
	}
}
