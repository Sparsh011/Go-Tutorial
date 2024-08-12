package main

import (
	"fmt"
	"math/rand/v2"
	"sort"
)

func main() {
	fmt.Println("Inside Slices package")

	var sliceList = []int{1, 2, 3, 4, 5, 6}

	fmt.Println("sliceList: ", sliceList)

	/*
		Initializing a Slice is similar to initializing Array. The difference
		is that we don't pass a fixed size when initializing a Slice and also, a
		Slice must be initialized, unlike Array
	*/

	sliceList = append(sliceList, 7)

	fmt.Println("sliceList after appending: ", sliceList)

	sliceList = append(sliceList[1:3]) // Creates a subarray from 1st index (inclusive) to 3rd index (exclusive)
	fmt.Println("sliceList after using colon syntax: ", sliceList)

	highScore := make([]int, 10)
	// This initializes slice with size 10, so using highScore[10] = 50 will give
	// index out of bounds error, but we can use append method that allows DMA.

	for i := range 10 {
		highScore[i] = generateRandomNumber()
	}

	fmt.Println("\nHigh score before sorting: ", highScore)

	// Sorts in descending order. We can use this code to allow custom sorting as well
	sort.Slice(highScore, func(i, j int) bool {
		return highScore[i] > highScore[j]
	})

	fmt.Println("\nHigh score after sorting in descending order: ", highScore)

	// Sorts in ascending order
	sort.Slice(highScore, func(i, j int) bool {
		return highScore[i] < highScore[j]
	})

	fmt.Println("\nHigh score after sorting in ascending order: ", highScore)

	rand.Shuffle(len(highScore), func(i, j int) {
		highScore[i], highScore[j] = highScore[j], highScore[i]
	})

	fmt.Println("\nShuffling after sorting: ", highScore)

	fmt.Println("\n\nSorting again...")

	sort.Ints(highScore)

	fmt.Println("\n\nNew sorted slice in ascending order: ", highScore)

	var orderedSlice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("\n\nBefore removing 3rd value from orderedSlice: ", orderedSlice)
	fmt.Println("\n\nRemoving 3rd value from orderedSlice... ")

	orderedSlice = append(orderedSlice[:2], orderedSlice[3:]...)
	fmt.Println("\n\nAfter removing 3rd value from orderedSlice: ", orderedSlice)
}

func generateRandomNumber() int {
	return rand.IntN(100)
}
