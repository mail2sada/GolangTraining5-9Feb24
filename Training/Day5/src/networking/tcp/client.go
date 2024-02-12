package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Demo: tcp client")
	address := ":4567"
	protocol := "tcp"
	conn, err := net.Dial(protocol, address)
	if err != nil {
		fmt.Println("Received error")
	} else {
		defer conn.Close()
		b := make([]byte, 1024, 1024)
		for {
			n, err := conn.Read(b)
			if err != nil {
				continue
			}
			fmt.Println("read ", n, "bytes")
			fmt.Println(string(b))
			conn.Write([]byte("Client hello"))
		}
	}
}
