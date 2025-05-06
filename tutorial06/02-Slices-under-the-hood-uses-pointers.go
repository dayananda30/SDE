package tutorial06

import (
	"fmt"
)

// SlicesUnderTheHood demonstrates how slices are implemented using pointers
// and how they can be used to create a new slice that shares the same underlying array.

func SlicesUnderTheHood() {
	mySlice := []int{1, 2, 3, 4, 5}
	var myAnotherSlice = mySlice

	mySlice[0] = 10

	fmt.Println("Original Slice:", mySlice)       // Output: Original Slice: [10 2 3 4 5]
	fmt.Println("Another Slice:", myAnotherSlice) // Output: Another Slice: [10 2 3 4 5]
}
