package main

import (
	"fmt"
	"strconv"
)

type Student struct {
	RollNo int
	Name   string
}

func (s Student) String() string {
	return `{"RollNo":` + strconv.Itoa(s.RollNo) + `,"Name":"` + s.Name + `"}`
}

func main() {
	fmt.Println("Demo Reflection")
	s := Student{RollNo: 1, Name: "Mayur Bhosale"}

	fmt.Println(s)
	str := fmt.Sprintln(s)
	fmt.Println(str)
}
