package tutorial07

import (
	"fmt"
	"time"
)

// WithGoroutines demonstrates how to perform a task using goroutines.

func WithGoroutines() {
	fmt.Println("Demonstrating Goroutines ...")
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		// It spawns up the process by running dbCall in a separate goroutine but didn't wait for it to finish.
		// This means that the main function will not wait for the goroutine to finish before moving on to the next iteration.
		go dbCall(i)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Time taken to fetch data:", time.Since(t0))
}
