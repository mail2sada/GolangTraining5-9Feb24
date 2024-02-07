package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
}

func (s Student) String() string {
	fmt.Println("String Mehod called")
	return fmt.Sprintln("RollNo:", s.RollNo, "\t Name:", s.Name, "\t Class:", s.Class)
}

func main() {
	fmt.Println("Demo: Reflection")
	std := Student{RollNo: 1, Name: "Pradeep", Class: "Btech"}
	fmt.Printf("%v", std)
}
