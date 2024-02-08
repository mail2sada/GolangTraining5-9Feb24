package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func WriteToChannel(ch chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}

func ReadFromChannel(gid int, ch chan int) {
	defer wg.Done()
	for {
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println("Gid:", gid, i)
	}
}

func main() {
	fmt.Println("Demo: Channels")

	var ch chan int = make(chan int)
	wg.Add(3)
	go WriteToChannel(ch)
	go ReadFromChannel(1, ch)
	go ReadFromChannel(2, ch)

	wg.Wait()

}
