package conditionalstmt

import "fmt"

func main() {
	fmt.Println("Logical operators")
	a, b, c := 10, 20, 30

	if (a > b) && (a > c) {
		fmt.Println("a is biggest")
	}
	// ||
	// !
}
