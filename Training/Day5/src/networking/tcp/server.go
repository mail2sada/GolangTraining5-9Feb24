package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Demo: tcp server")

	address := ":4567"
	protocol := "tcp"

	list, err := net.Listen(protocol, address)
	if err != nil {
		fmt.Println("failed to listen")
	} else {
		for {
			conn, err := list.Accept()
			if err != nil {
				fmt.Println("Received error while accepting connection")
			} else {
				go handleConnection(conn)
			}

		}
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	b := make([]byte, 1024, 1024)
	for {
		conn.Write([]byte("Server Hello..."))

		n, err := conn.Read(b)
		if err != nil {
			fmt.Println("Received error")
			continue
		}
		fmt.Println("Read ", n, "bytes")
	}
}
