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
	sList := []Student{}
	s.Calculate()
	sList = append(sList, s)
	s.RollNo = 2
	s.Name = "abc"
	s.Marks.S1 = 80

	s.Calculate()
	sList = append(sList, s)

	s.RollNo = 3
	s.Name = "xyz"
	s.Marks.S1 = 1000
	s.Calculate()
	sList = append(sList, s)

	b, err := json.MarshalIndent(sList, "\t", "")
	if err != nil {
		fmt.Println("Received error while parsing", err)
	} else {
		fmt.Println("Binary Content", b)
		fmt.Println("String Content", string(b))
	}

	newList := make([]Student, 0, 0)

	err = json.Unmarshal(b, &newList)

	if err != nil {
		fmt.Println("Failed to unmarshal", err)
	} else {
		for _, v := range newList {
			fmt.Println(v)
		}
	}
}
