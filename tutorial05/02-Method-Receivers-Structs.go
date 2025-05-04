package tutorial05

// https://dev.to/jpoly1219/structs-methods-and-receivers-in-go-5g4f

/*

structs are like classes in other languages, but they do not support inheritance.

Method receivers are like methods in other languages, but they are not bound to a specific instance of a struct.
They are bound to a specific type, which can be a struct or an interface.

*/

type Laptop struct {
	Brand string
	Model string
	Price int
}

func (l *Laptop) UpdatePrice(newPrice int) {
	l.Price = newPrice
}

func (l *Laptop) PrintDetails() {
	println("Brand:", l.Brand)
	println("Model:", l.Model)
	println("Price:", l.Price)
}
