package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Date and time")

	t := time.Now()

	expTime := t.Add(45 * time.Minute)

	fmt.Println(t, expTime)

	
}
