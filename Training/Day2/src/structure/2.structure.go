package main

import "fmt"

/*
1. define student entity
2. write details
3. print details
*/

type Student struct {
	RollNo int
	Name   string
	Class  string
	Div    string
}

func main() {
	fmt.Println("Demo: Student struct")
	var std Student = Student{RollNo: 1, Name: "Akash Kumar", Class: "10", Div: "B"}

	fmt.Println(std)

	fmt.Printf("%v\n", std)

	fmt.Println("----------------------------------------")
	fmt.Println("Rollno\tName\t\tClass\tDivision")
	fmt.Println("----------------------------------------")
	fmt.Println(std.RollNo, "\t", std.Name, "\t", std.Class, "\t", std.Div)
	fmt.Println("----------------------------------------")
}
