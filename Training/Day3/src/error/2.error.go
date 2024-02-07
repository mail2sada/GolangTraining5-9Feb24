package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Please enter non zzero for b")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Demo: Error")

	ans, err := Divide(100, 20)
	if err != nil {
		fmt.Println("Received error", err)
	} else {
		fmt.Println(ans)
	}

	ans, err = Divide(100, 0)
	if err != nil {
		fmt.Println("Received error", err)
	} else {
		fmt.Println(ans)
	}

}
