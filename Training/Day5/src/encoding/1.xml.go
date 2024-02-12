package main

import (
	"encoding/xml"
	"fmt"
)

type Student struct {
	Roll  int    `xml:"RollNo"`
	Name  string `xml:"Name"`
	Class string `xml:"Class"`
}

func main() {
	fmt.Println("Demo: XML handling")

	s := Student{Roll: 1, Name: "Anand Kumar", Class: "10B"}

	b, e := xml.Marshal(s)
	if e != nil {
		fmt.Println("Received error while Marshling", e)
	} else {
		fmt.Println("Binary content of encoded data", b)
		fmt.Println("XML content", string(b))
	}
}
