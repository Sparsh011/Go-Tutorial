package main

import (
	"fmt"
	"strconv"
	"strings"
)

func TypeConversions() bool {
	// Converting string to number:
	numString := "-05\n" // Deliberately added a \n here to test TrimSpace function
	fmt.Println("numString before conversion: ", numString)

	numNum, err := strconv.ParseInt(
		strings.TrimSpace(numString),
		10,
		64, // , is must at the end otherwise it gives compile time error
	)

	if err != nil {
		fmt.Println("Error: ", err)
		return true
	}

	fmt.Println("numString after conversion: ", numNum)
	fmt.Printf("Type of numNum after conversion: %T", numNum)
	println()

	/*
		Note: "-05" is also converted into -5.
	*/
	return true
}
