package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"name": "anil",
	"eid": 10,
	"department": "r & d",
	"designation": "engineer"
}
*/

type Employee struct {
	Name        string `json:"name"`
	Eid         int    `json:"eid"`
	Department  string `json:"department"`
	Designation string `json:"designation"`
}

func main() {
	fmt.Println("Demo: json encoding...")

	e := Employee{Name: "Anand", Eid: 100, Department: "MDE", Designation: "Engineer"}

	b, err := json.Marshal(e)

	if err != nil {
		fmt.Println("Received errpr while parsing")
	} else {
		fmt.Println(b)

		fmt.Println(string(b))
	}
	b, err = json.MarshalIndent(e, "\t", "\n")
	if err != nil {
		fmt.Println("Received errpr while parsing")
	} else {
		fmt.Println(b)

		fmt.Println(string(b))
	}

}
