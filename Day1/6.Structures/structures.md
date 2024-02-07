# Structure in Golang

A struct (strucutre) is a composite data type that groups together variables (fields or members) under a single name. It is a way to define a record or an object with multiple fields. Here are the key details about structures in Go:

## Struct Declaration:
The syntax for declaring a struct is as follows:

            Syntax:
                type StructName struct {
                    Field1 DataType1
                    Field2 DataType2
                    // ... additional fields
                }
1. StructName: The name of the struct.
2. Field1, Field2, etc.: The names of the fields in the struct.
3. DataType1, DataType2, etc.: The data types of the corresponding fields.

            Example:
                package main

                import "fmt"

                // Defining a struct named Person
                type Person struct {
                    FirstName string
                    LastName  string
                    Age       int
                }

                func main() {
                    // Creating an instance of the Person struct
                    person := Person{
                        FirstName: "John",
                        LastName:  "Doe",
                        Age:       30,
                    }

                    // Accessing struct fields
                    fmt.Println("First Name:", person.FirstName)
                    fmt.Println("Last Name:", person.LastName)
                    fmt.Println("Age:", person.Age)
                }
## Accessing Struct Fields:
You can access the fields of a struct using the dot (.) operator.

                fmt.Println("First Name:", person.FirstName)

## Initializing Structs:
There are multiple ways to initialize structs:

Using a composite literal.
Using the new keyword.
Creating an instance and initializing fields one by one.

                // Using a composite literal
                person := Person{
                    FirstName: "Alice",
                    LastName:  "Wonderland",
                    Age:       21,
                }

                // Using the new keyword
                personPtr := new(Person)
                personPtr.FirstName = "Bob"
                personPtr.LastName = "Builder"
                personPtr.Age = 35

                // Creating an instance and initializing fields one by one
                var anotherPerson Person
                anotherPerson.FirstName = "Charlie"
                anotherPerson.LastName = "Chaplin"
                anotherPerson.Age = 50

## Anonymous Structs:
You can create structs without explicitly defining a type. These are called anonymous structs.

            Example:
                package main

                import "fmt"

                func main() {
                    // Creating an anonymous struct
                    person := struct {
                        FirstName string
                        LastName  string
                        Age       int
                    }{
                        FirstName: "Jane",
                        LastName:  "Smith",
                        Age:       25,
                    }

                    // Accessing struct fields
                    fmt.Println("First Name:", person.FirstName)
                    fmt.Println("Last Name:", person.LastName)
                    fmt.Println("Age:", person.Age)
                }

## Embedded Structs:
You can embed one struct within another to create a composite struct.

            Example:
                package main

                import "fmt"

                type Address struct {
                    City  string
                    State string
                }

                type Person struct {
                    FirstName string
                    LastName  string
                    Address   Address // Embedded struct
                }

                func main() {
                    // Creating a Person instance with an embedded Address
                    person := Person{
                        FirstName: "John",
                        LastName:  "Doe",
                        Address: Address{
                            City:  "New York",
                            State: "NY",
                        },
                    }

                    // Accessing fields including the embedded struct
                    fmt.Println("First Name:", person.FirstName)
                    fmt.Println("Last Name:", person.LastName)
                    fmt.Println("City:", person.Address.City)
                    fmt.Println("State:", person.Address.State)
                }
## Comparing Structs:

Structs can be compared using the == and != operators. Two structs are equal if their corresponding fields are equal.

            Example:
                person1 := Person{FirstName: "John", LastName: "Doe", Age: 30}
                person2 := Person{FirstName: "John", LastName: "Doe", Age: 30}

                if person1 == person2 {
                    fmt.Println("Equal")
                } else {
                    fmt.Println("Not Equal")
                }

                
## Final Notes
Structs in Go provide a way to define and organize complex data structures. They are used to create custom types with multiple fields, encapsulating related information. Understanding how to declare, initialize, and use structs, as well as defining methods on structs, is fundamental to effective Go programming.


 