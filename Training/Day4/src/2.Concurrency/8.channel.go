package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func WriteData(ch chan int) {
	defer wg.Done()
	log.Println("Write: Wating for 1 second")
	time.Sleep(1 * time.Second)
	log.Println("Write: Writing data to channel")
	ch <- 1
	log.Println("Write: Data Written")
}

func ReadData(ch chan int) {
	defer wg.Done()

	log.Println("Read:")
	i := <-ch
	log.Println("Read Done", i)
}

func main() {
	wg.Add(2)
	fmt.Println("Demo: channels read write")
	ch := make(chan int)

	go WriteData(ch)
	go ReadData(ch)
	wg.Wait()
}
