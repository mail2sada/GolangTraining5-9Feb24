package main

import (
	"fmt"
	"sync"
)

/*
ReadNumber -> Squere -> Read and Print
*/
var wg = sync.WaitGroup{}

func ReadNumber(ch chan int, elements ...int) {
	defer wg.Done()
	defer close(ch)
	for _, v := range elements {
		ch <- v
	}
}

func Square(chIn chan int, chOut chan int) {
	defer wg.Done()
	defer close(chOut)
	for {
		num, ok := <-chIn
		if !ok {
			break
		}
		num *= num
		chOut <- num
	}
}

func Display(ch chan int) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(num)
	}

}

func main() {
	fmt.Println("Concurrency in Pattern Pipeline")
	wg.Add(3)
	readChannel := make(chan int)
	squareChannel := make(chan int)

	go ReadNumber(readChannel, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	go Square(readChannel, squareChannel)
	go Display(squareChannel)
	wg.Wait()
}
