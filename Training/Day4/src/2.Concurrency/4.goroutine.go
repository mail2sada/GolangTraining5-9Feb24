package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

var wg = sync.WaitGroup{}

var mutex = sync.Mutex{}

func IncrementCounter() {
	defer wg.Done()
	mutex.Lock()

	for i := 0; i < 1000; i++ {
		counter++
		time.Sleep(1 * time.Nanosecond)
	}
	mutex.Unlock()

}

func main() {

	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(time.Now())
	fmt.Println("Data: Race")
	wg.Add(5)
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	wg.Wait()
	fmt.Println("Value of Counter is", counter)

}
