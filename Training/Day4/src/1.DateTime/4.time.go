package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Date object creation")

	// 1/1/2024 10:10:10

	t := time.Date(2024, 1, 1, 10, 10, 10, 0, time.UTC)

	fmt.Println("Time is ", t)
}
