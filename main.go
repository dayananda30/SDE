package main

import (
	tutorial01 "github.com/dayananda30/sde/tutorial01"
	tutorial02 "github.com/dayananda30/sde/tutorial02"
	"github.com/fatih/color"
)

func main() {
	color.Red("--------- Tutorial 01 ---------")
	tutorial01.PrintHelloWorld()
	tutorial01.PrintStringFormatting()
	tutorial01.Variables()

	color.Red("--------- Tutorial 02 ---------")
	tutorial02.PrintMeMessage("Hello World")

}
