package main

import "fmt"

func DisplayMapWithKeyName(m map[string]int) {

	for key, val := range m {
		fmt.Println(key, ":", val)
	}
}

func MapIterator(m map[string]func()) {
	for _, val := range m {
		val()
	}
}

func main() {
	fmt.Println("Demo: Maps with functions")

	m := make(map[string]int)

	m["NewDelhi"] = 10000001
	m["Banglore"] = 10000002
	m["Mumbai"] = 1000003
	m["Chennai"] = 2000001
	DisplayMapWithKeyName(m)

	var functionMap = map[string]func(){}

	fmt.Println(functionMap)

	functionMap["Key1"] = func() { fmt.Println("Key1 function") }
	functionMap["Key2"] = func() { fmt.Println("Key2 function") }

	MapIterator(functionMap)

}
