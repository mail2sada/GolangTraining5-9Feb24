# Slice
A slice is a flexible and dynamic data structure built on top of arrays. Unlike arrays, slices are not of fixed size and can grow or shrink dynamically. Slices provide a more convenient and powerful way to work with sequences of data. Here are the key features and details about slices.

# Overview of Slice
Slice is a structure that has 3 attributes.
1. len -> Number of items in slice(length)
2. cap -> How many elements can be stored without reallocating(Capacity)
3. ptr -> A pointer pointing to the consecutively allocated memory.

Whan a slice is created and not initialized it will be a nil slice. Any operations on that will result in a panic.

            
                    var intSlice []int

                    
                    intSlice = []int{1,2,3,4,5,6,7,8,9,10}
                    

# Slice Declaration and Initialization:
A slice is declared without specifying a size. Instead, you create a slice by slicing an existing array or another slice. The syntax for creating a slice is as follows:
            Syntax:
                var sliceName []Type
                    or
                sliceName := []Type{elements}

            Example 1:
                var intSlice []int
                //here in above example slice is nil.
                // any operations on this will result in panic
                    or 
                intSlice = []int{1,2,3,4,5,6,7,8,9,10}
                // in the above line we are initializing the slice.
                // now we can perform operations on this slice.
            Example 2:
                // Slice can also be initialized using a builtin function make
                var intSlice []int
                //here in above example slice is nil.
                // any operations on this will result in panic
                intSlice = make([]int, 0, 5)
                // here make function will create a slice whose capacity is 5 and length is 0
                // now we can perform any slice opeations on this.
            Example 3:
                // slice can also decalred using the short declaration operator.
                intSlice := []int{1,2,3,4,5}
                        or
                intSlice := make([]int, 0, 5) 
                // this will create a int slice with 0 lenght and capacity 5
                // we need to insert the elements in slice using append method

# Accessing Elements of Slice
We can access the elements of slice like array using the index.

            Example:
                // Creating a slice
                intSlice := []{1,2,3,4,5}
                a := intSlice[0]
                fmt.Println("Value of a ",[a], intSlice[0])
                intSlice[0] = 100
                fmt.Println("Value of ", intSlice[0])
                //in order to insert element at position 5, we should use append
                intSlice = append(intSlice, 6)
                fmt.Println(intSlice)




