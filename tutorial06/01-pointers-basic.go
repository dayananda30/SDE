package tutorial06

import (
	"fmt"
)

func PointerBaiscs() {
	x := 10

	y := &x

	fmt.Println("Value of x:", x)  // 10
	fmt.Println("Value of y:", *y) // 10

	*y = 20
	fmt.Println("Value of x after changing y:", x)  // 20
	fmt.Println("Value of y after changing y:", *y) // 20

}
