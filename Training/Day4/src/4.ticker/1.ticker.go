package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: Ticker")

	tck := time.NewTicker(2 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range tck.C {
			fmt.Println("Ticker Expired", v)
		}

	}()
	wg.Wait()
}
