package tutorial03

import (
	"github.com/fatih/color"
)

// Maps are a built-in data type in Go that associates keys with values.
// They are similar to dictionaries in Python or hash tables in other languages.
// Maps are defined using the following syntax:
// var mapName map[keyType]valueType
// For example, to declare a map that associates strings with integers, you would write:
// var ages map[string]int

// You can also initialize a map at the time of declaration:
// var ages = map[string]int{"Alice": 25, "Bob": 30}

func Maps() {
	// Declare and initialize a map of string to int
	ages := map[string]int{"Alice": 25, "Bob": 30}
	color.Blue("Map: %v", ages)

	// Accessing map elements
	color.Blue("Alice's age: %d", ages["Alice"])
	color.Blue("Bob's age: %d", ages["Bob"])

	// Modifying map elements
	ages["Alice"] = 26
	color.Blue("Modified Alice's age: %d", ages["Alice"])

	// Adding new elements to the map
	ages["Charlie"] = 35
	color.Blue("Added Charlie's age: %d", ages["Charlie"])

	// Length of the map
	color.Blue("Length of the map: %d", len(ages))

	// Deleting an element from the map
	delete(ages, "Bob")
	color.Blue("Deleted Bob's age: %d", ages["Bob"])

	// Iterating over a map
	for name, age := range ages {
		color.Blue("%s's age: %d", name, age)
	}

	// Checking if a key exists in the map
	age, exists := ages["Alice"]
	if exists {
		color.Blue("Alice's age: %d", age)
	} else {
		color.Blue("Alice not found")
	}

	var myMap map[string]int = make(map[string]int)
	myMap["Alice"] = 25

	// nested maps
	var nestedMap = map[string]map[string]int{ // key: value
		"Alice": {"age": 25, "height": 170},
		"Bob":   {"age": 30, "height": 180},
	}
	color.Blue("Nested Map: %v", nestedMap)
}
