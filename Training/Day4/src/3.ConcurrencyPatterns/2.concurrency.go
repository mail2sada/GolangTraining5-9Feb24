package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func Read(ele ...int) chan int {
	ch := make(chan int)

	/////
	wg.Add(1)

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

	chOut := make(chan int)
	wg.Add(1)

	go func() {
		defer wg.Done()

		defer close(chOut)
		for {
			num, ok := <-ch
			if !ok {
				break
			}
			num *= num
			chOut <- num
		}
	}()

	return chOut

}

func Display(ch chan int) {
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			num, ok := <-ch
			if !ok {
				break
			}
			fmt.Println(num)
		}
	}()
}

func main() {
	fmt.Println("Pipeline ")
	/*chIn := Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	ch := Square(chIn)
	Display(ch)*/

	Display(Square(Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)))

	Display(Square(Square(Read(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))))

	wg.Wait()
}
