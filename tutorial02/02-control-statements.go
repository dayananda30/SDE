package tutorial02

import (
	"github.com/fatih/color"
)

func ControlStatements(num1, num2 int) {
	// if statement
	if num1 > num2 {
		color.Yellow("Num1 is greater than Num2")
	} else if num1 < num2 {
		color.Yellow("Num1 is less than Num2")
	} else {
		color.Yellow("Num1 is equal to Num2")
	}

	// switch statement
	switch num1 {
	case 0:
		color.Yellow("Num1 is zero")
	case 1:
		color.Yellow("Num1 is one")
	default:
		color.Yellow("Num1 is greater than one")
	}

	// for loop
	for i := 0; i < 5; i++ {
		color.Yellow("For loop iteration: %d", i)
	}

	// while loop
	i := 0
	for i < 5 {
		color.Yellow("While loop iteration: %d", i)
		i++
	}
}
