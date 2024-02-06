# Strings
Strings in Go are a sequence of UTF-8 encoded characters and are immutable. In Go, a string is represented as a sequence of bytes, and the string is a built-in type of Go. We will see how to declare a string, using string and some common operations on string.

## Strings are Immutable:
In Go, strings are immutable, meaning you cannot modify a string by changing its characters. If you need to modify a string, you'll need to create a new one.

## String Declaration:
The basic syntax for declaring a string is as follows:
                syntax:
                var str string
                //or
                str := "Hello, Go!"

## String Length:
The length of a string can be obtained using the len function:

                str := "Hello, Go!"
                length := len(str)

## Accessing Characters:
Individual characters of a string can be accessed using indexing. Note that strings in Go are zero-indexed:

                char := str[0] // Accessing the first character

## String Concatenation:
1. Strings can be concatenated using the + operator:

                str1 := "Hello"
                str2 := "Go"
                result := str1 + " " + str2 // "Hello Go"

2. Strings can be concatinated using the fmt.Sprintf, fmt.Sprint, fmt.Sprintln

            Example:
                str1 := "Hello"
                str2 := "Go"
                result := fmt.Sprint(str1,str2)

            Example
                str1 := "Hello"
                str2 := "Go"
                result := fmt.Sprintln(str1,str2)
            
            Example
                str1 := "Hello"
                str2 := "Go"
                result := fmt.Sprintf("%s%s",str1,str2)

## Iterating Over Strings:
You can iterate over the characters of a string using a for loop:

            Example
                for index, char := range str {
                    fmt.Printf("Index: %d, Character: %c\n", index, char)
                }

## Converting to Byte Slice:
You can convert a string to a byte slice using the []byte type conversion:
            Example
                byteSlice := []byte(str)

## Converting to Rune Slice:
To iterate over Unicode characters in a string, convert it to a rune slice:
            Example
                for _, char := range []rune(str) {
                    fmt.Printf("Character: %c\n", char)
                }

## String Comparison:
1. Strings can be compared using the == and != operators:
            Example
                str1 := "Hello"
                str2 := "Hello"

                if str1 == str2 {
                    fmt.Println("Equal")
                } else {
                    fmt.Println("Not Equal")
                }
2. Strings can compared using standar library package method reflect.DeepEqual
            Example
                str1 := "Hello"
                str2 := "Hello"

                if reflect.DeepEqual(str1, str2) {
                    fmt.Println("Equal")
                } else {
                    fmt.Println("Not Equal")
                }
## Strings Package:
The strings package in Go provides a variety of functions for manipulating strings, such as strings.Contains, strings.HasPrefix, strings.ToLower, strings.ToUpper, etc.

            Example
                import "strings"

                contains := strings.Contains(str, "Go")

## String Formatting:
The fmt package provides formatting options for strings using Printf:
            Example
                name := "Alice"
                age := 25

                fmt.Printf("Name: %s, Age: %d\n", name, age)

## String Conversion:
You can convert other types to strings using the strconv package:
            Example
                import "strconv"

                number := 42
                str := strconv.Itoa(number)

## Raw String
In Go, raw string literals are enclosed in backticks (`) and are used to create strings that preserve the literal value of each character within the string, including newline characters and special characters.

            Example:
                package main

                import "fmt"

                func main() {
                    // Regular string literal
                    regularStr := "This is a regular string.\nIt has a newline."

                    // Raw string literal
                    rawStr := `This is a raw string.
                It preserves newlines and special characters.`

                    fmt.Println("Regular String:")
                    fmt.Println(regularStr)

                    fmt.Println("\nRaw String:")
                    fmt.Println(rawStr)
                }

In this example, regularStr is a string created with double quotes, and rawStr is created using backticks. The raw string preserves the newline character without the need for escape sequences.

It's important to note that backticks should not be used for regular string declarations. Backticks are specifically for creating raw string literals. If you need to include escape sequences or special characters, you should use double quotes.

Here's an example of using escape sequences in a regular string:

            Example
                escapedStr := "This is a string with escape sequences.\nNewline\tTab"
                fmt.Println(escapedStr)

Both regular strings and raw strings have their use cases, and the choice depends on the specific requirements of your string content.

## Final Notes:
    Understanding how strings work in Go, their immutability, and common operations such as concatenation, iteration, and conversions is crucial for effective Go programming. The built-in functions and packages provided by Go, such as strings and strconv, offer powerful tools for working with strings in a clean and efficient manner.





            