package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Time")

	var t time.Time = time.Now()

	fmt.Println("Current time is", t)

	fmt.Println("Today is Day is", t.Day())
	fmt.Println("Today week day is", t.Weekday())
	fmt.Println("Current month is", t.Month())
	fmt.Println("Year", t.Year())

	fmt.Printf("Current time is %d:%d:%d\n", t.Hour(), t.Minute(), t.Second())
}
