package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Running Multiple http Server")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hit default route"))
	})

	http.HandleFunc("/details", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("We are in details page"))
	})

	go func() {
		router := http.NewServeMux()

		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method)
			w.Write([]byte("We are in control server"))
		})
		router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method)

			w.Write([]byte("Starting Application server"))
		})
		router.HandleFunc("/stop", func(w http.ResponseWriter, r *http.Request) {
			fmt.Println(r.Method)

			w.Write([]byte("Stopping Application Server"))
		})
		http.ListenAndServe(":8081", router)
	}()

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Failed to Start Server")
	}

}
