package main

import "fmt"

func Divide(a, b int) (ans int, err error) {
	if b == 0 {
		err = ZeroError //fmt.Errorf("Please enter non zero value for b, current value of b is %d", b)
		return
	}
	if b == -1 {
		err = NegError //fmt.Errorf("I am not handing when b is %d", b)
		return
	}
	ans = a / b
	return
}

var ZeroError = fmt.Errorf("Please enter non zero value for b, current value of b is 0")
var NegError = fmt.Errorf("I am not handing when b is -1")

func main() {
	fmt.Println("Demo: Functions")

	const b [5]string = [5]string{"H", "M", "4", "3", "2"}

	a, e := Divide(10, 0)
	if e != nil {
		if e == ZeroError {
			//handle
		} else if e == NegError {

		}
		fmt.Println("Received error", e)
	} else {
		fmt.Println("Received answer", a)
	}

	a, e = Divide(10, -1)
	if e != nil {
		fmt.Println("Received error", e)
	} else {
		fmt.Println("Received answer", a)
	}
}
