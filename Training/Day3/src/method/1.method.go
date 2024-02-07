package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
}

func (s *Student) SetStudent(rlno int, name string) {
	s.RollNo = rlno
	s.Name = name

	fmt.Println(s)
}

func (s *Student) SetRollNo(rlno int) *Student {
	s.RollNo = rlno
	return s
}

func (s *Student) SetName(name string) *Student {
	s.Name = name
	return s
}

func main() {
	fmt.Println("Demo: Methods")

	var str Student

	str.SetStudent(1, "hello")

	fmt.Println(str)

	var newStd Student

	newStd.SetRollNo(1).
		SetName("Hello")
	fmt.Println(newStd)
}
