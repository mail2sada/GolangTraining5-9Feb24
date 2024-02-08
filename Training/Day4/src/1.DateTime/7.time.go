package main

import (
	"fmt"
	"time"
)

// day/month/year hour:min:second timezone
func main() {
	fmt.Println("Demo: Time formating")

	t := time.Now()

	fmt.Println("Current time is", t)
	layout := "2/1/2006 03:04:05 PM MST"
	s := t.Format(layout)
	fmt.Println(s)
}
