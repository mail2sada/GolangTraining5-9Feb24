package main

import "fmt"

type Display interface {
	Print()
}

type Getter interface {
	GetID() int
	GetName() string
}

type Student struct {
	Roll int
	Name string
}

func (s Student) GetID() int {
	return s.Roll
}

func (s Student) GetName() string {
	return s.Name
}

func (s Student) Print() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Roll No\t\tName")
	fmt.Println("--------------------------------------------------")
	fmt.Println(s.Roll, "\t\t", s.Name)
	fmt.Println("--------------------------------------------------")
}

type Employee struct {
	Eid  int
	Name string
}

func (s Employee) GetID() int {
	return s.Eid
}

func (s Employee) GetName() string {
	return s.Name
}

func (e Employee) Print() {
	fmt.Println("--------------------------------------------------")
	fmt.Println("Eid No\t\tName")
	fmt.Println("--------------------------------------------------")
	fmt.Println(e.Eid, "\t\t", e.Name)
	fmt.Println("--------------------------------------------------")
}

func main() {
	fmt.Println("Demo:Interface")

	var v Display
	e := Employee{Eid: 1, Name: "xyz"}
	v = e
	v.Print()

	s := Student{Roll: 10, Name: "Anand"}
	v = s
	v.Print()

	var g Getter = e
	fmt.Println(g.GetID())
	fmt.Println(g.GetName())

	g = s

	fmt.Println(g.GetID())
	fmt.Println(g.GetName())

}
