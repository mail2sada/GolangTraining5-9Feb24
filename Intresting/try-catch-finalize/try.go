package main

import "fmt"

type Exception interface{}

type Block struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

func (b Block) Do() {
	if b.Finally != nil {
		defer b.Finally()
	}
	if b.Catch != nil {
		defer func() {
			r := recover()
			if r != nil {
				b.Catch(r)
			}
		}()
	}
	if b.Try != nil {
		b.Try()
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
