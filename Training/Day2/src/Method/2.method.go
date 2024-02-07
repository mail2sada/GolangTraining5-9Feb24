package main

import "fmt"

type Marks struct {
	rollNo             int
	s1, s2, s3, s4, s5 int
}

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

func (m Marks) Print() {
	fmt.Println("---------------------------------------------------")
	fmt.Println("S1\t\t S2\t\tS3\t\tS4\t\tS5\t\t")
	fmt.Println("---------------------------------------------------")
	fmt.Println(m.s1, "\t\t", m.s2, "\t\t", m.s3, "\t\t", m.s4, "\t\t", m.s5)
	fmt.Println("---------------------------------------------------")
}

func main() {
	fmt.Println("Demo: Methods")

	s := Student{rollNo: 1, name: "Anand", class: "10"}

	m := Marks{rollNo: 1, s1: 100, s2: 90, s3: 85, s4: 99, s5: 100}

	s.Print()
	m.Print()

}
