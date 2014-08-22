package main

import (
    "fmt"
)

func main() {
    // https://gobyexample.com/channels

    /* non buffered channels only fill if there is something
       waiting to flush it
     */
    nonBuffChan  := make(chan int)

    // need something to continue flushing it
    go func() {
        for {
            _ = <-nonBuffChan
        }
    }()

    nonBuffChan <- 1

    /* buffered channels can take multiple values and
       then be flushed later. If we go over "3" here
       an error will be generated saying all goroutines are
       asleep
     */

    bufferedChan := make(chan int, 3)
    for i:=0; i<3; i++ {
        bufferedChan <- i
    }

    // here we can flush all the values in the buffered
    // channel, but if we try for 4 we'll get the
    // all goroutines are asleep error ...
    for i:=0; i<3; i++ {
        fmt.Printf("%d\n", <-bufferedChan)
    }

}
