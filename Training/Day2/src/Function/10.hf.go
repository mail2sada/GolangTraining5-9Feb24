package main

import "fmt"

func ConnectToNetworkResource(url string, fn func(status string)) (ok int) {
	i := 0
	fn("Initializing")
	for {
		//fmt.Println("Initializing..")

		i++
		if i == 10000 {
			break
		}

	}
	//fmt.Println("Initialized")
	fn("Initialized")

	i = 0
	fn("Connnecting")
	for {
		//fmt.Println("Connnecting...")

		i++
		if i == 10000 {
			break
		}

	}
	fn("Connected")
	//fmt.Println("Connected")
	return
}

func ConnectionStatus(status string) {
	fmt.Println("from ConnectionStatus:", status)
}

func main() {
	fmt.Println("Demo: Higher Order function")

	// ConnectToNetworkResource("https://mavenir.com", func(status string) {
	// 	// action
	// 	fmt.Println("Received status:", status)
	// })

	ConnectToNetworkResource("https://mavenir.com", ConnectionStatus)
}
