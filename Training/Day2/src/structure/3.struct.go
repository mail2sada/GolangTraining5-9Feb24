package main

import "fmt"

type Student struct {
	RollNo           int
	Name, Class, Div string
	Score            Marks
}

type Marks struct {
	RollNo                                     int
	Lang1, Lang2, SocialStudies, Maths, Sciene int
}

func main() {
	fmt.Println("Demo: Embedded struct")

	var Std Student = Student{RollNo: 1,
		Name:  "Anil",
		Class: "10",
		Div:   "b",
		Score: Marks{RollNo: 1,
			Lang1:         100,
			Lang2:         99,
			SocialStudies: 100,
			Maths:         100,
			Sciene:        99},
	}

	fmt.Println(Std)

}
