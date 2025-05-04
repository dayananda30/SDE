package tutorial04

import (
	"strings"

	"github.com/fatih/color"
)

// Strings are a built-in data type in Go that represent a sequence of characters.
// They are similar to strings in other programming languages.

// Strings are immutable, meaning that once a string is created, it cannot be changed.
// However, you can create a new string by concatenating existing strings or modifying them.
// Strings are defined using double quotes ("") or backticks (``) for raw strings.

func Strings() {
	// Declare and initialize a string
	var str string = "Hello, World!"
	color.Magenta("String: %s", str)

	// String length
	color.Magenta("String length: %d", len(str))

	// String concatenation
	var str2 string = " How are you?"
	var str3 string = str + str2
	color.Magenta("Concatenated string: %s", str3)

	// String indexing
	color.Magenta("First character: %c", str[0])
	color.Magenta("Last character: %c", str[len(str)-1])

	// String slicing
	color.Magenta("Substring: %s", str[7:12])

	// String comparison
	if str == "Hello, World!" {
		color.Magenta("Strings are equal")
	} else {
		color.Magenta("Strings are not equal")
	}

	// String functions
	color.Magenta("String to upper case: %s", strings.ToUpper(str))
	color.Magenta("String to lower case: %s", strings.ToLower(str))
}
