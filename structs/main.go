package main

import "fmt"

func main() {
	/*
		There is no inheritance in Golang!
	*/

	sparsh := User{
		"Sparsh Chadha",
		"ABCD@sparsh.com",
		true,
	}

	fmt.Println("Printing sparsh: ", sparsh)

	// We can also use printf to get a more detailed print statement using $+v
	fmt.Printf("\nPrinting Sparsh with clarity: %+v\n", sparsh)

	// Accessing a field of object
	fmt.Printf("Name is: %v and email is %v\n", sparsh.Name, sparsh.Email)

	if num := 3; num < 10 {
		fmt.Println("\nNumber is less than 10")
	} else {
		fmt.Println("\nNumber is > 10")
	}
}

type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
}
