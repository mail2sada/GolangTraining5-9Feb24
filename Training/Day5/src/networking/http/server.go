package main

import (
	"fmt"
	"net/http"
	"os"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received http request at root")
	fmt.Println(r)
	w.Write([]byte("We have received a request at root"))
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("Welcome to Test Http Server")

	http.HandleFunc("/", HandleRoot)

	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println("Received error")
		os.Exit(-1)
	}
}
