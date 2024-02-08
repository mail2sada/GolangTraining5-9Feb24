package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: go routines")

	go fmt.Println("I want to execute this line in other go routine")
	go fmt.Println("This is second go routine")

	time.Sleep(1 * time.Second)
}
