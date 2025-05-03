package tutorial03

// Arrays are a fixed-size data structure that can hold multiple values of the same type.
// They are useful when you know the number of elements in advance and want to store them in a single variable.
// In Go, arrays are defined using the following syntax:
// var arrayName [size]dataType
// For example, to declare an array of integers with a size of 5, you would write:
// var numbers [5]int

// You can also initialize an array at the time of declaration:
// var numbers = [5]int{1, 2, 3, 4, 5}

import (
	"github.com/fatih/color"
)

func Arrays() {
	// Declare and initialize an array of integers
	numbers := [5]int{1, 2, 3, 4, 5}
	color.Blue("Array: %v", numbers)

	// Accessing array elements
	color.Blue("First element: %d", numbers[0])
	color.Blue("Second element: %d", numbers[1])

	// Modifying array elements
	numbers[0] = 10
	color.Blue("Modified first element: %d", numbers[0])

	// Length of the array
	color.Blue("Length of the array: %d", len(numbers))
}
