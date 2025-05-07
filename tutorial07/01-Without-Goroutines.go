package tutorial07

import (
	"fmt"
	"time"
)

// WithoutGoroutines demonstrates how to perform a task without using goroutines.

var dbData = []string{"Data1", "Data2", "Data3", "Data4", "Data5"}

func WithoutGoroutines() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		fmt.Println("Fetching data from DB...")
		dbCall(i)
	}
	fmt.Println("Time taken to fetch data:", time.Since(t0))
}

func dbCall(i int) {
	fmt.Println("Fetching data from DB...")
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched from DB:", dbData[i])
}
