package main

import (
	"flag"
	"fmt"
	"os"
)

/*
 add -a 100 -b 100
*/

func main() {
	fmt.Println("Demo: Command line tool")
	a := flag.Int("a", -1, "Please key in a value for a to add")
	b := flag.Int("b", -1, "Please key in a value for b to add")

	flag.Parse()

	if *a == -1 && *b == -1 {
		flag.Usage()
		return
	}

	fmt.Println("We received arguments", *a, *b)

	sum := *a + *b

	fmt.Println("Sum of ", *a, " &", *b, "is:", sum)

	for _, v := range os.Args {
		fmt.Println("Received argument ", v)
	}

}
