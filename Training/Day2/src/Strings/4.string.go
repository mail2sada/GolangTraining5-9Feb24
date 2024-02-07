package main

import "fmt"

func main() {
	fmt.Println("Demo: String concat")

	str1 := "Hello"
	str2 := "Go"

	result := str1 + str2

	fmt.Println(result)

	result1 := fmt.Sprint(str1, str2)
	result2 := fmt.Sprintln(str1, str2)

	result3 := fmt.Sprintf("%s%s", str1, str2)

	fmt.Println("0", result)
	fmt.Println("1", result1)

	fmt.Println("2", result2)

	fmt.Println("3", result3)

}
