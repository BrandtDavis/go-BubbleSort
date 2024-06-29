package main

import (
	"fmt"
	"math/rand"
)

func makeRandomSlice(numItems, max int) []int {
	randomSlice := make([]int, numItems)
	for i := 0; i < numItems; i++ {
		randomSlice[i] = rand.Intn(max)
	}

	return randomSlice
}

func printSlice(slice []int, numItems int) {
	if numItems > len(slice) {
		numItems = len(slice)
	}

	for i := 0; i < numItems; i++ {
		fmt.Println(slice[i])
	}
}

func checkSorted(slice []int) {
	for i := 1; i < len(slice); i++ {
		if slice[i-1] > slice[i] {
			fmt.Println("The slice is NOT sorted!")
			return
		}
	}
	fmt.Println("The slice is sorted")
}

func bubbleSort(slice []int) {
	for i := 0; i < len(slice); i++ {
		swapped := false

		for j := 1; j < len(slice); j++ {
			if slice[j-1] > slice[j] {
				temp := slice[j-1]
				slice[j-1] = slice[j]
				slice[j] = temp

				swapped = true
			}
		}

		if !swapped {
			break
		}
	}
}

func main() {
	// Get the number of items and maximum item value.
	var numItems, max int
	fmt.Printf("# Items: ")
	fmt.Scanln(&numItems)
	fmt.Printf("Max: ")
	fmt.Scanln(&max)

	// Make and display an unsorted slice.
	slice := makeRandomSlice(numItems, max)
	printSlice(slice, 40)
	fmt.Println()

	// Sort and display the result.
	bubbleSort(slice)
	printSlice(slice, 40)

	// Verify that it's sorted.
	checkSorted(slice)
}
