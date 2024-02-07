package main

import "fmt"

func TestFuncForDefer() {
	fmt.Println("We are in test defer")                            //4                           //3
	defer fmt.Println("TestFuncForDefer:: defer call from line 7") //5
}

func main() {
	fmt.Println("Demo: defer") //1
	defer TestFuncForDefer()
	fmt.Println("After Defer")                     // 2
	defer fmt.Println("This is last line of main") // 3
}
