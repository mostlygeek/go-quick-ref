package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Creates a channel that outputs a bunch of integers
func IntPipe(num int) chan int {
	out := make(chan int)

	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	// generate integers inside a goroutine
	go func() {
		for i := 0; i < num; i++ {
			out <- i
			time.Sleep(time.Duration(10+rand.Intn(50)) * time.Millisecond)
		}

		// important to close the channel otherwise
		// we will have a "all goroutines asleep" crash
		close(out)
	}()

	return out
}

func main() {

	// flushing the channel using range example
	intChan1 := IntPipe(5)
	for i := range intChan1 {
		fmt.Printf("range outC example: %v\n", i)
	}

	fmt.Println("\n\nSELECT EXAMPLE\n")

	// select example
	intChan2 := IntPipe(20)
	intChan3 := IntPipe(20)

	intChan2_done, intChan3_done := false, false

	for {
		if intChan2_done && intChan3_done {
			return
		}

		// select can be used to receive from multiple
		// channels (and different types)

		select {
		case i, chan_ok := <-intChan2:
			if chan_ok {
				fmt.Printf("A:%v\n", i)
			} else {
				intChan2_done = true
			}

		case i, chan_ok := <-intChan3:
			if chan_ok {
				fmt.Printf("     B:%v\n", i) // indented for visibility
			} else {
				intChan3_done = true
			}
		}

	}
}
