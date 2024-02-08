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
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	log.Println("Wrote: 5 elements")
	ch <- 6
	log.Println("Wrote: 6th element")
	ch <- 7
	log.Println("Wrote 7")
}

func ReadData(ch chan int) {
	defer wg.Done()
	log.Println("ReadData")
	time.Sleep(5 * time.Second)
	d := <-ch
	log.Println("Read value", d)
	time.Sleep(5 * time.Second)
	d = <-ch
	log.Println("Read value", d)
}

func main() {

	fmt.Println("Demo: Buffered Channel")

	ch := make(chan int, 5)
	wg.Add(2)
	go WriteData(ch)
	go ReadData(ch)
	wg.Wait()

}
