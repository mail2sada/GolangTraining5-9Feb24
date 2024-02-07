package main

import "fmt"

type Exception interface{}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (tcf Block) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			r := recover()
			if r != nil {
				tcf.Catch(r)
			}
		}()
	}
	if tcf.Try != nil {
		tcf.Try()
	}
}

func Divide(a, b int) int {
	defer func() {
		rec := recover()
		if rec != nil {
			fmt.Println("Recover from exception ", rec)
		}
	}()
	if b == 0 {
		panic("Please enter a non zero value for b")
	}
	return a / b
}

func main() {

	fmt.Println("Demo error")

	fmt.Println(Divide(10, 2)) //

	fmt.Println(Divide(10, 0)) //

	fmt.Println("TAG", Divide(10, 5))
	/*
		Block{
			Try: func() {
				fmt.Println(Divide(10, 2))

				fmt.Println(Divide(10, 0))
			},
			Catch: func(e Exception) {
				fmt.Println("recevied exeception", e)
			},
			Finally: func() {
				fmt.Println("We have hit finally")
			},
		}.Do()
	*/
}
