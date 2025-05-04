package tutorial05



import (
	"fmt"
)


// Interfaces are a powerful feature in Go that allow you to define a set of methods that a type must implement.
// They are similar to interfaces in other programming languages, but they are not bound to a specific type.
// An interface is defined using the following syntax:
// type InterfaceName interface {
// 	MethodName() returnType
// 	MethodName() returnType
// 	...
// }


// For example, to define an interface that represents a shape, you would write:
type Shape interface {
	Area() float64
	Perimeter() float64
}

// You can then define a struct that implements the interface by providing implementations for all of its methods.
type Rectangle struct {
	Width  float64
	Height float64
}
func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius int
}

func (c *Circle) Area() float64 {
	return 3.14 * float64(c.Radius*c.Radius)
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * float64(c.Radius)
}


func PrintShapeDetails(s Shape) {
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

