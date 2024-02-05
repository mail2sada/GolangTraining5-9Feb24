package main

import "fmt"

/*
	syntax
	var <name of array> [size]type = [size]type{val, val1, val2 ...}
*/

func main() {
	fmt.Println("Demo: arrays")

	var weekDay [7]string = [7]string{
		"Sunday", "Monday", `Tuesday`, "Wednsday", "Thurday", "Friday",
	}

	fmt.Println(weekDay[0])
	fmt.Println(weekDay[6])

	weekDay[6] = `Saturday`

	fmt.Println(weekDay[6])

}
