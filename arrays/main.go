package main

import "fmt"

func main() {
	/*
		Note: In Golang, Arrays are not used a lot because of minimal operations that we can perform on them.
			  Eg. We can't even sort an array, hence we use SLICES most of the time in Go
	*/
	var arr = [3]string{"Hello", "Bye", "Good night"}
	fmt.Println("Array is : ", arr)
	fmt.Println("Lenght of array: ", len(arr))

	fmt.Println("Traversing array: ")

	for item := range arr {
		fmt.Println("Item index is : ", item, " and item value is : ", arr[item])
	}
}
