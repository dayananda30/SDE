package tutorial05

import (
	"fmt"
)

// Structs are a composite data type in Go that groups together variables (fields) under a single name.
// They are similar to classes in other programming languages, but they do not support inheritance.
// Structs are defined using the following syntax:

// type StructName struct {
// 	FieldName fieldType
// 	FieldName fieldType
// 	...
// }

// For example, to define a struct that represents a person, you would write:

type Person struct {
	Name    string
	Age     int
	Address string
}

// You can also define methods for structs, which are functions that operate on the struct's fields.
// Methods are defined using the following syntax:
//
//	func (s *StructName) MethodName() {
//		// method body
//	}
//
// For example, to define a method that prints the person's name, you would write:
func (p *Person) PrintName() {
	fmt.Println("Name:", p.Name)
}
func (p *Person) PrintAge() {
	fmt.Println("Age:", p.Age)
}
func (p *Person) PrintAddress() {
	fmt.Println("Address:", p.Address)
}

// You can create a new instance of a struct using the following syntax:
// person := StructName{FieldName: value, FieldName: value, ...}
// For example, to create a new person, you would write:
// person := Person{Name: "Alice", Age: 25, Address: "123 Main St"}

// You can also create a new instance of a struct using the following syntax:
// person := new(StructName)

func Structs() {
	// Create a new instance of the Person struct
	person := Person{Name: "Alice", Age: 25, Address: "123 Main St"}

	// Accessing struct fields
	fmt.Println("Name:", person.Name)
	fmt.Println("Age:", person.Age)
	fmt.Println("Address:", person.Address)

	// Modifying struct fields
	person.Name = "Bob"
	person.Age = 30
	person.Address = "456 Elm St"

	// Accessing modified struct fields
	fmt.Println("Modified Name:", person.Name)
	fmt.Println("Modified Age:", person.Age)
	fmt.Println("Modified Address:", person.Address)

	// Calling struct methods
	person.PrintName()
	person.PrintAge()
	person.PrintAddress()
}
