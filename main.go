package main

import (
	tutorial01 "github.com/dayananda30/sde/tutorial01"
	tutorial02 "github.com/dayananda30/sde/tutorial02"
	"github.com/dayananda30/sde/tutorial03"
	"github.com/dayananda30/sde/tutorial04"
	"github.com/dayananda30/sde/tutorial05"
	"github.com/fatih/color"
)

func main() {
	color.Red("--------- Tutorial 01 Basic Hello World and String Formatting ---------")
	tutorial01.PrintHelloWorld()
	tutorial01.PrintStringFormatting()
	tutorial01.Variables()

	color.Red("--------- Tutorial 02 :: Functions and Control Statements ---------")
	tutorial02.PrintMeMessage("Hello World")
	res, err := tutorial02.Division(10, 2)
	if err == nil {
		color.Yellow("Result: %d", res)
	}

	tutorial02.ControlStatements(10, 20)

	color.Red("--------- Tutorial 03 :: Arrays Slices Maps and Loops ---------")
	tutorial03.Arrays()
	tutorial03.Slices()
	tutorial03.Maps()
	tutorial03.Loops()

	color.Red("--------- Tutorial 04 :: Strings Runes and Bytes ---------")
	tutorial04.Strings()

	color.Red("--------- Tutorial 05 :: Structs and Interfaces---------")
	tutorial05.Structs()

}
