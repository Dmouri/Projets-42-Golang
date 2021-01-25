package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {

	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
