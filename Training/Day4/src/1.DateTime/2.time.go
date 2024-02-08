package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo: time")

	t := time.Now()

	for i := 0; i < 10000; i++ {

	}
	fmt.Println("time taken to process", time.Since(t))

	res := TestProfiling(100)
	fmt.Println(res)

	res = TestProfiling(12179)
	fmt.Println(res)

	res = TestProfiling(31007)
	fmt.Println(res)

	res = TestProfiling(999983)
	fmt.Println(res)

}

func TestProfiling(n int) bool {
	defer func(t time.Time) {
		fmt.Println("It took", time.Since(t))

	}(time.Now())
	isNumberDivisible := true

	for i := 2; i < (n - 1); i++ {
		if n%i == 0 {
			isNumberDivisible = false
			break
		}
	}

	return isNumberDivisible
}
