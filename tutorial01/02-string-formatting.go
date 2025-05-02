package tutorial01

import (
	"fmt"
)

// PrintStringFormatting demonstrates string formatting in Go.
func PrintStringFormatting() {
	i := 42
	f := 3.14159
	b := true
	s := "GoLang"

	fmt.Printf("Integer: %d\n", i)
	fmt.Printf("Float: %f\n", f)
	fmt.Printf("Boolean: %t\n", b)
	fmt.Printf("String: %s\n", s)

	// Using Printf for formatted output
	fmt.Printf("Hello, %s!\n", "World")

	// Using Sprintf to format a string and store it in a variable
	formattedString := fmt.Sprintf("Hello, %s!", "Go")
	fmt.Println(formattedString)

	// Using Fprintf to write formatted output to a file (example)
	// file, err := os.Create("output.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer file.Close()
	// fmt.Fprintf(file, "Hello, %s!\n", "File")

	// Using Errorf to create an error with formatted message
	// err := fmt.Errorf("an error occurred: %s", "some error")
	// if err != nil {

	// 	log.Println(err)
	// }
	// Using Println for simple output

	// Right-align a string in a 10-character wide field

	// Left-align a string in a 10-character wide field
	formattedStringLeft := fmt.Sprintf("%-10s", "Gopher")
	fmt.Println("Left-Aligned String:", formattedStringLeft)

}
