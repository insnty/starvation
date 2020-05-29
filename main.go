/* Concurrent process starvation demonstrated with two concurrent processes,
 * where one process (starving_proc) starves the other (normal_proc) on the shared resource 'mutex' */
package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var mutex sync.Mutex
	var procs sync.WaitGroup

	var starving_count int64
	var normal_count int64

	timeout := 1 * time.Second

	starving_proc := func() {
		defer procs.Done()

		for start := time.Now(); time.Since(start) < timeout; {
			mutex.Lock()
			time.Sleep(3 * time.Nanosecond)
			mutex.Unlock()

			starving_count++
		}
		log.Printf("starving_proc: finished %v operations\n", starving_count)
	}

	normal_proc := func() {
		defer procs.Done()

		for start := time.Now(); time.Since(start) < timeout; {
			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			mutex.Lock()
			time.Sleep(1 * time.Nanosecond)
			mutex.Unlock()

			normal_count++
		}
		log.Printf("normal_proc: finished %v operations\n", normal_count)
	}

	procs.Add(2)
	go starving_proc()
	go normal_proc()

	procs.Wait()

	log.Println("all processes finished")
	log.Printf("starving_proc finished %v more operations than normal_proc\n", starving_count-normal_count)
}
