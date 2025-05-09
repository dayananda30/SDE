package main

import (
	tutorial01 "github.com/dayananda30/sde/tutorial01"
	tutorial02 "github.com/dayananda30/sde/tutorial02"
	"github.com/dayananda30/sde/tutorial03"
	"github.com/dayananda30/sde/tutorial04"
	"github.com/dayananda30/sde/tutorial05"
	"github.com/dayananda30/sde/tutorial06"
	"github.com/dayananda30/sde/tutorial07"
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

	hp_laptop := tutorial05.Laptop{
		Brand: "HP",
		Model: "Pavilion",
		Price: 50000,
	}

	hp_laptop.PrintDetails()

	hp_laptop.UpdatePrice(60000)
	hp_laptop.PrintDetails()

	rectangle := tutorial05.Rectangle{
		Width:  10,
		Height: 5,
	}

	circle := tutorial05.Circle{
		Radius: 7,
	}

	tutorial05.PrintShapeDetails(&rectangle)
	tutorial05.PrintShapeDetails(&circle)

	color.Red("--------- Tutorial 06 :: Pointers ---------")
	tutorial06.PointerBaiscs()
	tutorial06.SlicesUnderTheHood()

	color.Red("--------- Tutorial 07 :: Goroutines and Channels ---------")
	tutorial07.WithoutGoroutines()
	tutorial07.WithGoroutines()
	tutorial07.WaitGroups()
	tutorial07.Mutex()

}
