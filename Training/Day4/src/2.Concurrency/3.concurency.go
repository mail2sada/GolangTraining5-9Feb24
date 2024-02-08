package main

import (
	"fmt"
	"sync"
)

func spawnGoroutine(msg string, wg *sync.WaitGroup) {
	wg.Add(1)

	go PrintMessage(msg, wg)
}

func PrintMessage(msg string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(msg)
}

func main() {
	fmt.Println("Demo: Wait groups")
	var wg sync.WaitGroup = sync.WaitGroup{}
	spawnGoroutine("Hello", &wg)
	spawnGoroutine("Welcome", &wg)
	spawnGoroutine("To", &wg)
	spawnGoroutine("Go", &wg)
	spawnGoroutine("Concurrency Training", &wg)
	wg.Wait()

	fmt.Println("All threads executed succefully")
	fmt.Println("Now terminating main")
}
