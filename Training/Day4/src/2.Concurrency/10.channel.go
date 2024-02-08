package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func Write(ch chan int) {
	defer wg.Done()
OuterLoop:
	for {
		select {
		case ch <- 1:
			fmt.Println("Data is written")
			break OuterLoop
		case i := <-ch:
			fmt.Println(i)
		default:
			print("Hello")

		}
	}

}

func Read(ch chan int) {
	defer wg.Done()

	time.Sleep(5 * time.Second)
	i := <-ch
	fmt.Println(i)
}

func main() {
	fmt.Println("Demo: select")
	wg.Add(2)
	ch := make(chan int)
	go Write(ch)
	go Read(ch)
	wg.Wait()
}
