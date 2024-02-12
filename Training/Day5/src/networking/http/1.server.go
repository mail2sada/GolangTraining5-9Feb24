package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server with multiple routes")
	/*
		/
		/details
	*/

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We have hit root function"))
	})

	http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is test web server, we are traing this out\n Please contact it@mavenir.com"))
	})

	http.ListenAndServe(":8080", nil)
}
