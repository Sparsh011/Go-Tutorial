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

	myObj := MethodInStructStruct{
		"Sparsh",
		"hello",
		true,
	}

	if myObj.GetNonPublicField() {
		fmt.Println("Non public field is true, so returning from here...")
		fmt.Println("")
		return
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

	checkVarArgType(2, 3, 4, 5, 6)
}

func checkVarArgType(values ...int) {
	fmt.Printf("\n\nType of values: %T\n", values) // It's type is slice because it prints []int. In case of array, [size]int would've been printed
	// var a [5]int
	// fmt.Printf("\n\nType of a: %T\n", a)
}

type User struct {
	Name       string
	Email      string
	IsLoggedIn bool
}

type MethodInStructStruct struct {
	Name           string
	Email          string
	nonPublicField bool
}

// This is a method technically in golang terms
func (obj MethodInStructStruct) GetNonPublicField() bool {
	// Like this we can access private fields
	// Also, obj is a value type parameter, not a reference type so even if we change any field of obj, it won't reflect in the original instance of MethodInstanceStructStruct
	return obj.nonPublicField
}
