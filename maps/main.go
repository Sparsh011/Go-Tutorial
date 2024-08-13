package main

import (
	"fmt"
)

func main() {
	var myMap = make(map[string]string)

	myMap["Sp"] = "Sparsh"
	myMap["Ch"] = "Chadha"

	fmt.Println("\nMap: ", myMap)
	fmt.Println("\nGetting Sp: ", myMap["Sp"])
	fmt.Println("\nGetting a value that doesn't exist: ", myMap["Sparsh"]) // prints empty

	delete(myMap, "Ch") // Used to delete something from Map. Can also be used to delete from slices

	fmt.Println("\nMap after deletion: ", myMap)

	myMap["Ch"] = "Chadha"

	fmt.Println("\nIterating through myMap...")

	for key, value := range myMap {
		fmt.Printf("\nFor key %v, value is %v\n", key, value)
	}
}
