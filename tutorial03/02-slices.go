package tutorial03

import (
	"github.com/fatih/color"
)

// Slices are a more flexible and powerful data structure than arrays.
// They allow you to create dynamic arrays that can grow and shrink in size.
// Slices are defined using the following syntax:
// var sliceName []dataType
// For example, to declare a slice of integers, you would write:
// var numbers []int

func Slices() {
	// Declare and initialize a slice of integers
	numbers := []int{1, 2, 3, 4, 5}
	color.Blue("Slice: %v", numbers)

	// Accessing slice elements
	color.Blue("First element: %d", numbers[0])
	color.Blue("Second element: %d", numbers[1])

	// Modifying slice elements
	numbers[0] = 10
	color.Blue("Modified first element: %d", numbers[0])

	// Length of the slice
	color.Blue("Length of the slice: %d", len(numbers))

	// Appending elements to a slice
	numbers = append(numbers, 6, 7, 8)
	color.Blue("Appended slice: %v", numbers)
	// Slicing a slice
	slicedNumbers := numbers[1:4]
	color.Blue("Sliced slice: %v", slicedNumbers)

	var mySlice []int = make([]int, 3, 5)
	color.Blue("Slice with make: %v", mySlice)
}
