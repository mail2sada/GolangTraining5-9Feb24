package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo: Multiple tickers")

	t1 := time.NewTicker(1 * time.Second)
	t2 := time.NewTicker(2 * time.Second)
	t3 := time.NewTicker(3 * time.Second)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case t := <-t1.C:
				fmt.Println("t1 expired", t)
			case t := <-t2.C:
				fmt.Println("t2 expired", t)
			case t := <-t3.C:
				fmt.Println("t3 expired", t)
				// default:
				// 	fmt.Println("timers are running...")
			}
		}

	}()
	wg.Wait()
}
