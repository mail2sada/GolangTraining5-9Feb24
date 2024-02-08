package main

import (
	"fmt"
	"sync"
	"time"
)

var counter = 0

var wg = sync.WaitGroup{}

// var mutex = sync.Mutex{}
var mutex = sync.RWMutex{}

func IncrementCounter() {
	defer wg.Done()
	for i := 0; i < 1000; i++ {
		mutex.Lock()
		counter++
		mutex.Unlock()
	}
}

func PrintCounterValue() {
	defer wg.Done()
	for {

		mutex.RLock()
		fmt.Println("Value of Counter", counter)
		mutex.RUnlock()
		if counter == 5000 {
			break
		}
		time.Sleep(10 * time.Microsecond)
	}
}

func main() {
	defer func(t time.Time) {
		fmt.Println(time.Since(t))
	}(time.Now())
	fmt.Println("Demo: RWMutex")
	wg.Add()
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	go IncrementCounter()
	go PrintCounterValue()
	go PrintCounterValue()
	wg.Wait()

	fmt.Println("Final Counter Value", counter)

}
