# Hello World!

Every go source code file has to be saved with .go extension, every go source file should belong to a package. package name is declared using below syntax

 "package <PackageName>"

 package is a keyword and <PackageName> is a of your choice, however it should follow below rules.

 1. Package name should begin with an alphabet
 2. Package name should not contain any space in between
 3. Package name can't use special charecters other than "_" (underscore)
 4. Package name can contain intergers

 Example:
    package math

    package fmt

    package maths1

    package math_1

    etc.

 Every go executable, has to have a main package. main package is known as executable package, func main is the stating point of the execution.  


 It is defined as 
    
    func main() {

    }

func is a keyword to define functions.

    package main 
    import "fmt"

    func main() {
        fmt.Println("Welcome to Go training")
    }

import is a keyword to import the external packages

import "fmt" imports the standard package fmt and its functionalities.

func main  defines the function main.
opening "{" and closing "}" defines the main function boundries
fmt.Println("Welcome to Go training") lines calls the function Println from fmt package, this outputs "Welcome to Go training" to console.

Next Topic: Keywords