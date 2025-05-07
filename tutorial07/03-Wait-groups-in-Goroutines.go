package tutorial07

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroups demonstrates how to use wait groups to synchronize goroutines.
func WaitGroups() {
	fmt.Println("Demonstrating Wait Groups ...")
	t0 := time.Now()

	var wg sync.WaitGroup
	for i := 0; i < len(dbData); i++ {
		wg.Add(1) // Increment the wait group counter
		go dbCallWithWaitGroup(i, &wg)
	}

	wg.Wait() // Wait for all goroutines to finish
	fmt.Println("Time taken to fetch data:", time.Since(t0))
}

func dbCallWithWaitGroup(i int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement the wait group counter when the goroutine completes
	fmt.Println("Fetching data from DB...")
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched from DB:", dbData[i])
}
