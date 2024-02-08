package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Demo:Timer")

	var tm *time.Timer = time.NewTimer(5 * time.Second)
	fmt.Println("Timer Started at", time.Now())
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		t := <-tm.C
		fmt.Println("Timer Expired at", t)
	}()
	wg.Wait()
	tm.Stop()

}
