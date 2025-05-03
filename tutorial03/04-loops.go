package tutorial03

import (
	"github.com/fatih/color"
)

// Loops are used to execute a block of code repeatedly.

// In Go, there is only one loop construct: the `for` loop.
// The `for` loop can be used in several different ways:
// 1. Traditional for loop
// 2. For loop with a condition
// 3. For loop with a range
// 4. Infinite loop
// 5. Nested loops
// 6. Break and continue statements
// 7. Labels

func Loops() {
	// Traditional for loop
	for i := 0; i < 5; i++ {
		color.Blue("Traditional for loop: %d", i)
	}

	// For loop with a condition
	i := 0
	for i < 5 {
		color.Blue("For loop with a condition: %d", i)
		i++
	}

	// For loop with a range
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		color.Blue("For loop with a range: index %d, value %d", index, value)
	}

	// Infinite loop
	// Uncomment the following lines to see an infinite loop
	// for {
	// 	color.Blue("Infinite loop")
	// }

	// Nested loops
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			color.Blue("Nested loop: i %d, j %d", i, j)
		}
	}

	// Break statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		color.Blue("Break statement: %d", i)
	}

	// Continue statement
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		color.Blue("Continue statement: %d", i)
	}

	// Labels
outerLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				break outerLoop
			}
			color.Blue("Label: i %d, j %d", i, j)
		}
	}
}
