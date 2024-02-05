package main

import "fmt"

func main() {
	fmt.Println("Demo: Multi Dimensional arrays")

	var mulArray = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for i := 0; i < 3; i++ {
		fmt.Println("")
		for j := 0; j < 3; j++ {
			fmt.Print("  ", mulArray[i][j])
		}
		fmt.Println("\n")
	}

	for _, v := range mulArray {
		fmt.Println(v)
		for _, v1 := range v {
			fmt.Println(v1)
		}
	}

}
