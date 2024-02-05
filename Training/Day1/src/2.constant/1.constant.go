package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Constant declaration")
	const a int = 0
	const b, c = a, "hello"
	const e, f, g string = "", "", ""

	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

}
