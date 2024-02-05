package main

import "fmt"

func main() {
	fmt.Println("Demo: Nested loop with go to statement")

	i := 0

LABEL:
	i++
	for {
		i++
		fmt.Println(i)
		j := 0
		if i%3 == 0 {
			continue
		}
		for {
			j++
			if i%10 == 0 {
				goto LABEL
			}
			if j%10 == 3 {
				continue
			}
			fmt.Println(i, j)
			if j > 100 && j%5 == 0 {
				break
			}
		}
		if i > 20 {
			break
		}
	}
}
