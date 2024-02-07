package main

import "fmt"

type Employee struct {
	Eid  int
	Name string
}

func main() {
	fmt.Println("Demo: Struct pointers")
	var emp Employee = Employee{Eid: 1, Name: "Suraj"}

	fmt.Println(emp)

	var ptrEmp *Employee = &emp

	fmt.Println("Contents of emp using ptrEmp", *ptrEmp)
	fmt.Println("Member of struct are ", ptrEmp.Eid, ptrEmp.Name)
}
