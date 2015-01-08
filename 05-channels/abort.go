package main

// example of sending an abort message to stop running
// go routines from sending data

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	num_producers = 5
)

func main() {

	var wg sync.WaitGroup
	rand.Seed(time.Now().UnixNano())

	out := make(chan int)

	// this channel is used across all the producers in a select
	// makes use of the pattern that RECEIVES on closed channels
	// do not block
	done := make(chan bool)
	// defer close(done) <-- putting a defer here would be idiomatic

	// start a bunch of goroutines
	for i := 1; i <= num_producers; i++ {
		go func(id int) {
			defer wg.Done()

			wait := time.Duration((100 + rand.Int63n(900))) * time.Millisecond
			fmt.Printf("Producer:%v waiting %v\n", id, wait)

			after := time.After(wait)

			for {
				// print out because the output is interesting
				// the first producer to send to `out` will
				fmt.Println("  > ", id, "before switch{}")
				select {
				case <-after:
					fmt.Println("  ! ", id, "is first!")
					out <- id
				case <-done: // will run when done is closed
					return
				}
				fmt.Println("  > ", id, "after switch{}")
			}
		}(i)
	}

	wg.Add(num_producers)

	// wait for all producers to finish before
	// closing `out` as sends to a closed channel
	// result in a panic
	go func() {
		wg.Wait()
		close(out)
	}()

	// block on a first response
	fmt.Printf("------\n  > Producer %v finished first\n", <-out)
	fmt.Println("Closing done channel")
	close(done)

	// at this point main() will exit, but the runtime will wait
	// for the rest of the application to finish, it seems that
	// a few things are cleaned up like output to the console
	// so fmt.PrintX doesn't work anywhere...
	//
	// however `go run --race abort.go` doesn't report any race
	// conditions so it should be ok.
}
