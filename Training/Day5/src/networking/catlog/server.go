package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Product struct {
	Pid         int    `json:"pid"`
	Name        string `json:"name"`
	Description string `json:"desc"`
	InStock     bool   `json:"stock"`
	Quantity    int    `json:"qty"`
	Category    string `json:"cat"`
}

func main() {
	fmt.Println("Demo: Product Catlog")
	var ProductList = []Product{Product{Pid: 1, Name: "Laptop", Description: "15 inch", InStock: true, Quantity: 10, Category: "Computers"},
		Product{Pid: 1, Name: "iPhone 15", Description: "128 GB", InStock: true, Quantity: 5, Category: "Mobile"},
	}
	//mx := http.NewServeMux()
	mx := mux.NewServeMux()
	mx.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Received request at root"))
	})
	mx.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Listing the products\n"))
		//b, err := json.Marshal(ProductList)
		b, err := json.MarshalIndent(ProductList, "\t", "\n")
		if err != nil {
			w.Write([]byte("Internal server error"))
			fmt.Println("Internal server error", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(b)
		w.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":8080", mx)
	if err != nil {
		fmt.Println("Received error", err)
	}
	fmt.Println(ProductList)

}
