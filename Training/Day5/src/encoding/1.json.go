package main

import (
	"encoding/json"
	"fmt"
)

type Marks struct {
	S1    int     `json:"Science"`
	S2    int     `json:"Maths"`
	S3    int     `json:"MentalAbility"`
	Total int     `json:"Total"`
	Avg   float64 `json:"Average"`
}

func (m *Marks) Calculate() {
	m.Total = m.S1 + m.S2 + m.S3

	m.Avg = float64(m.Total) / float64(3)
}

type Student struct {
	RollNo int    `json:"roll-no"`
	Name   string `json:"name"`
	Class  string `json:"class"`
	Marks  `json:"Marks"`
}

func main() {
	fmt.Println("Demo: nested marshling")

	s := Student{RollNo: 1, Name: "Anil Kumar", Class: "10A", Marks: Marks{S1: 100, S2: 99, S3: 100}}
	s.Calculate()

	b, err := json.MarshalIndent(s, "\t", "")
	if err != nil {
		fmt.Println("Received error while parsing", err)
	} else {
		fmt.Println("Binary Content", b)
		fmt.Println("String Content", string(b))
	}
}
