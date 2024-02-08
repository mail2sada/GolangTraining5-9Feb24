package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Timer")
	wg := sync.WaitGroup{}
	wg.Add(1)
	time.AfterFunc(5*time.Second, func() {
		defer wg.Done()
		fmt.Println("Timer expired...")
	})
	fmt.Println("123232")
	wg.Wait()
}
