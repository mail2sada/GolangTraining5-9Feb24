package main

import "fmt"

type Student struct {
	rollNo      int
	name, class string
}

func (s Student) Print() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("RollNo\t\t Name\t\t Class")
	fmt.Println("---------------------------------------------------")
	fmt.Println(s.rollNo, "\t\t", s.name, "\t\t", s.class)
	fmt.Println("---------------------------------------------------")

}

func main() {
	fmt.Println("Demo: Methods")

	s := Student{rollNo: 1, name: "Anand", class: "10"}

	s.Print()
}
