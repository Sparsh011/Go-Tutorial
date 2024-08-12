package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var result bool = TypeConversions()

	if result {
		fmt.Println("Returning because of true")
		return
	}

	var userName string = "Sparsh"                     // Declares a variable of type string
	fmt.Printf("Variable is of type: %T \n", userName) // %T is used to print the type of the variable
	var isLoggedIn = false
	fmt.Printf("Is logged in: %v \n", isLoggedIn)
	isLoggedIn = true
	fmt.Printf("Is logged in now: %v \n", isLoggedIn)

	var unassignedBool bool
	fmt.Println("Unassigned bool is ", unassignedBool)

	/*
		Note
		1. If variable is not initialized with a value and we try to use it,
		   Go does not allocate garbage value. Eg. in case of int, it assigns 0 and
		   in case of string, it assigns empty value and in case of bool, it assigns false
		2. We can also declare variables without assigning a type as well.
		3. We can also declare variabled without var keyword using :=. We must add an initializing value when using this operator
		   unlike using var keyword. This is also called Walrus operator. We cannot use walrus
		   operator outside of a method.
		4. use const for declaring constants
		5. VVIMP - In Golang, we declare public variables and methods by keeping first character
		   uppercase.
	*/

	count := "hello"
	fmt.Println("Count is ", count)

	/*
		Comma, Ok (also known as comma, err) syntax is the analogous of try catch.
		Like swift, we can use _ to denote that the error value won't be utilized so don't give us compile time error
	*/

	reader := bufio.NewReader(os.Stdin)

	input, error := reader.ReadString('\n')

	println("Input from user: ", input)
	fmt.Println("Error: ", error)
}
