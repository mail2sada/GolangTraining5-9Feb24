package main

import "fmt"

func PrintWelcomeMessage(name string) {
	fmt.Println(name, "Welcome to go training")
}

func main() {
	fmt.Println("Demo: Functions Arguments")
	PrintWelcomeMessage("Anil")
	PrintWelcomeMessage("Pradeep")
	PrintWelcomeMessage("Anand")
}
