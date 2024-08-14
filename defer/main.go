package main

import "fmt"

func main() {
	// Defer follows a LIFO pattern when multiple defer statements are used
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	fmt.Println("Starting reverse counting: ")

	fmt.Println("\n\n\nNow starting loop")

	DeferInLoop()
}

func DeferInLoop() {
	// These 2 printlns add a new line after Now starting loop fp
	fmt.Println("")
	fmt.Println("")

	// These 2 printlns add new lines after all the items of loop are printed.
	defer fmt.Println("")
	defer fmt.Println("")

	for i := 0; i < 10; i++ {
		defer fmt.Println("Item : ", i)
	}
}
