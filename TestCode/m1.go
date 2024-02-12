package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("http service")
	http.NewServeMux()
}
