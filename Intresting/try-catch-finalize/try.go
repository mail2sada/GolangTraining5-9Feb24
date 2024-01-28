package main

import "fmt"

type Block struct {
	Try      func()
	Catch    func(any)
	Finalize func()
}

func (b Block) Do() {
	if b.Finalize != nil {
		defer b.Finalize()
	}
}

func main() {
	fmt.Println("Demo Try Catch Finalize in Go...")
}
