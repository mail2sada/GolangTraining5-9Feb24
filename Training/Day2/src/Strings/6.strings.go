package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Demo: String Builder")

	var builder strings.Builder = strings.Builder{}

	builder.WriteByte(65)
	builder.WriteRune('B')
	builder.Write([]byte{67, 68, 69})
	builder.WriteString("FGH")

	str := builder.String()

	fmt.Println(str)
	fmt.Println(builder)

}
