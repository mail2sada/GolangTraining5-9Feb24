package main

import "fmt"

type Marks struct {
	TotalMarks int
	AvgMarks   float64
}

type Student struct {
	Rlno int
	Name string
	Marks
}

func (s *Student) SetName(name string) *Student {
	s.Name = name
	return s
}

func (s *Student) SetRollNo(rno int) *Student {
	s.Rlno = rno
	return s
}

func (m *Marks) SetTotal(total int) *Marks {
	m.TotalMarks = total
	return m
}
func (m *Marks) SetAvg(avg float64) *Marks {
	m.AvgMarks = avg
	return m
}

func main() {
	fmt.Println("Demo: Prompted Method")
	var s Student
	s.SetRollNo(1).SetName("Hello").SetTotal(499).SetAvg(99.99)
	fmt.Println(s)
}
