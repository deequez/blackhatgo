package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	// Explicitly track work that remains to prevent go routine from exiting before work is completed.
	// 1. Create sync.WaitGroup which acts as a synchronized counter.
	var wg sync.WaitGroup
	for i := 1; i < 1024; i++ {
		// 2. Increment the synchronized counter each time a go routine is created to scan a port.
		wg.Add(1)
		go func(j int) {
			// 3. Decrement the counter every time one unit of work has been performed.
			defer wg.Done()
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				// port is filtered or closed
				return
			}
			conn.Close()
			fmt.Printf("%d open\n", j)
		}(i)
	}
	// 4. Block until all work is done and counter has returned to zero.
	wg.Wait()
}
