package tutorial02

import (
	"errors"

	"github.com/fatih/color"
)

func PrintMeMessage(message string) {
	color.Yellow("PrintMeMessage: %s", message)
}

func Division(a int, b int) (int, error) {
	color.Yellow("Division: %d / %d", a, b)
	if b == 0 {
		return 0, errors.New("division by zero")
	}

	return a / b, nil
}
