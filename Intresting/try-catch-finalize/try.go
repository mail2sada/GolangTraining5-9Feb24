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

func main() {
	fmt.Println("Demo Try Catch Finalize in Go...")
	Block{
		Try: func() {
			fmt.Println("We are in Try block")
		},
		Catch: func(e Exception) {
			fmt.Println("Received exception e")
		},
		Finally: func() {
			fmt.Println("We have are in finally")
		},
	}.Do()
}
