package main

import (
	"flag"
	"fmt"
)

/*
-n integer
-s string
*/

func main() {
	fmt.Println("Demo: Command Line Arguments")
	var num int
	var s string
	flag.IntVar(&num, "n", 0, "Please enter a non zero value for -n option")
	flag.StringVar(&s, "s", "", "please enter a valid string for -s")

	flag.Parse()

	if num == 0 || s == "" {
		flag.Usage()
	} else {
		fmt.Println("Value of -n", num, "Value of -s", s)
	}

}
