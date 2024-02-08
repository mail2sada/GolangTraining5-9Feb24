package main

import (
	"fmt"
	"sync"
)

const FanOutCount = 3

var wg = sync.WaitGroup{}

func Read(ele ...int) chan int {
	wg.Add(1)
	ch := make(chan int)
	go func() {
		defer wg.Done()
		defer close(ch)
		for _, v := range ele {
			ch <- v
		}
	}()

	return ch
}

func Square(ch chan int) chan int {
	wg.Add(1)
	chOut := make(chan int)

	swg := sync.WaitGroup{}

	for i := 0; i < FanOutCount; i++ {
		swg.Add(1)
		go func() {
			defer swg.Done()
			for num := range ch {
				num *= num
				chOut <- num
			}
		}()
	}

	go func() {
		defer wg.Done()
		defer close(chOut)
		swg.Wait()
	}()

	return chOut
}

func Display(ch chan int) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for num := range ch {
			fmt.Println(num)
		}
	}()
}

func main() {
	fmt.Println("Demo: FanOut-FanIn")

	//Display(Square(Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15)))

	Display(Square(Square(Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12))))
	wg.Wait()
}
