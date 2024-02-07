package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Demo: String Package")

	var str1 = "Hello"
	var str2 = "Hi"
	fmt.Println(str1, str2)

	strings.Contains(str1, "a")
}
