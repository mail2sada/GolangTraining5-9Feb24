package main

import (
	"fmt"
	"time"
)

func PrintMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	fmt.Println("Demo: go routine")

	go PrintMessage("Hello")
	go PrintMessage("Welcome")
	go PrintMessage("to")
	go PrintMessage("Go")
	go PrintMessage("Concurency training")

	time.Sleep(10 * time.Minute)

}
