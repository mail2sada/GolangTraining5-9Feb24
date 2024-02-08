package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: Parse time")

	s := "2024/01/01 10:10:10 PM"

	layout := "2006/01/02 03:04:05 PM"

	t, e := time.Parse(layout, s)

	if e != nil {
		fmt.Println("Received error while Parsing time", s, layout, e)
	} else {
		fmt.Println(t)
	}

}
