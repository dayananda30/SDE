package tutorial02

import (
	"github.com/fatih/color"
)

func PrintMeMessage(message string) {
	color.Yellow("PrintMeMessage: %s", message)
}
