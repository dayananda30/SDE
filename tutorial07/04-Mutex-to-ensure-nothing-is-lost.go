package tutorial07

import (
	"fmt"
	"sync"
	"time"
)

// Mutex demonstrates how to use mutexes to ensure that only one goroutine can access a resource at a time.
func Mutex() {
	fmt.Println("Demonstrating Mutex ...")
	t0 := time.Now()

	var mu sync.Mutex
	for i := 0; i < len(dbData); i++ {
		mu.Lock() // Lock the mutex before accessing the shared resource
		go dbCallWithMutex(i, &mu)
	}
	time.Sleep(10 * time.Second)
	fmt.Println("Time taken to fetch data:", time.Since(t0))
}

func dbCallWithMutex(i int, mu *sync.Mutex) {
	defer mu.Unlock() // Unlock the mutex when the goroutine completes
	fmt.Println("Fetching data from DB...")
	time.Sleep(2 * time.Second)
	fmt.Println("Data fetched from DB:", dbData[i])
}
