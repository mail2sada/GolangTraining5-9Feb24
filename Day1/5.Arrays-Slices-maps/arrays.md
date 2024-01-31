# Arrays
An array is a fixed-size, ordered collection of elements of the same type. The size of an array is specified at the time of declaration and cannot be changed during the program's execution. Go arrays are zero-indexed, meaning the index of the first element is 0, the second element is 1, and so on.

## Array Declaration:
The syntax for declaring an array in Go is as follows:
            Syntax:
                var variableName [size]dataType

                var is the keyword used to declare variables.
                variableName is the name of the array variable.
                size is the fixed size of the array.
                dataType is the type of elements that the array can hold.
            Example
                var numbers [5]int   // An array of 5 integers
                var colors [3]string // An array of 3 strings

## Array Initialization:
Arrays in Go can be initialized when declared or later using the index.
        Example
                numbers := [5]int{1, 2, 3, 4, 5}

                // Initializing an array without specifying the size (size is determined by the number of elements)
                colors := [...]string{"red", "green", "blue"}
### Initializing Specific Elements of an Array
We can initialize only specific elements of an array instead of intializing all the elements. Other elements will be by default intialized to 0

        Example
                number := [5]int{0:1, 2:10, 4:100}
                here number[0] = 1, number[2] = 10 and number[4] = 100.
                remaining all elements will be initialized with 0

## Accessing Elements:
You can access elements of an array using their index. The index starts from 0 and goes up to size-1.

        Example:
            firstNumber := numbers[0]    // Accessing the first element
            secondColor := colors[1]     // Accessing the second element
            numbers[2] = 10              // Modifying the value of the third element

## Length of an Array:
The length of an array is the number of elements it can hold. It is determined by the size specified during declaration.

        Example:
            lengthOfNumbers := len(numbers)    // len() function returns the length of the array

## Looping through elements of array
Using simple for loop we can access elements of an arrays

        Example:
            array := [5]int {0,1,2,3,4}
            for i:=0 ;i < len(5); i++ {
                fmt.Println(array[i])
            }

## Iterating Over an Array:
You can use a for loop with the range keyword to iterate over the elements of an array.

        Example:
            for index, value := range numbers {
                fmt.Printf("Index: %d, Value: %d\n", index, value)
            }
## Limitations of Arrays:
1. Arrays have a fixed size, which means you need to know the size at compile time.
2. Copying an array creates a new copy of all elements, which can be inefficient for large arrays.

Next Topic: Slice
