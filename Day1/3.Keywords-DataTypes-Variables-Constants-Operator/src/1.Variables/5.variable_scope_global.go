package main

import "fmt"

// WelcomeMessage is global variable
var welcomeMessage = "Hello and Welcome to Golang training Day1"

func main() {
	fmt.Println("Demo Global Variable")
	fmt.Println("This is WelcomeMessage ", welcomeMessage)
	//Calling Test,
	testFunc()
}

func testFunc() {
	fmt.Println("Printing Welcome Message from test function", welcomeMessage)
}
