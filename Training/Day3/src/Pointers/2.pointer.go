//Pointer to pointer

package main

import "fmt"

func main() {
	fmt.Println("Demo: Double Pointer")

	var a int = 100

	ptr := &a

	dptr := &ptr

	fmt.Println("Value of a ", a)
	fmt.Println("Addreess of a", &a)
	fmt.Println("Contents of ptr", ptr)
	fmt.Println("Address of ptr", &ptr)
	fmt.Println("Contents of dptr", dptr)

	fmt.Println("Contents of ptr using dptr", *dptr)

	fmt.Println("Contents of a using dptr", **dptr)
}
