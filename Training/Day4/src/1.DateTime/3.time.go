package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: epoc time")

	t := time.Now()

	fmt.Println(t)

	s := t.Unix()

	fmt.Println(s)

	t = time.Unix(s, 2334003)

	fmt.Println(t)

	t = time.Unix(121223345, 0)

	fmt.Println(t)
}
