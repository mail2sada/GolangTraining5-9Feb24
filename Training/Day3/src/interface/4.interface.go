package main

import "fmt"

func Print(element ...interface{}) {
	for _, v := range element {
		fmt.Println(v)
	}
}

func AddAnything(e ...interface{}) (int, error) {
	result := 0
	for _, v := range e {
		switch v.(type) {
		case int:
			fmt.Println("int")
			result += v.(int)
		case float32:
			fmt.Println("float32")

			result += int(v.(float32))

		case float64:
			fmt.Println("float64")

			result += int(v.(float64))
		case uint:
			fmt.Println("uint")

			result += int(v.(uint))
		default:
			fmt.Println("Unhandled type")
		}
	}
	return result, nil
}

func main() {
	fmt.Println("Variadic function using empty interface")
	Print(1, "hello", 3.14159, true, false)

	res, _ := AddAnything(1, 3.1, 5, "Hello")

	fmt.Println(res)
}
