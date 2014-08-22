package main

import (
    "fmt"
    "time"
    "math/rand"
)


func main() {
    // https://gobyexample.com/channels

    // make a channel that can buffer 100 integers
    myChannel := make(chan int, 100)

    go func() {
        for {
            m := <-myChannel
            fmt.Println(m)
        }
    }()

    for {
        myChannel <- rand.Intn(1000)
        time.Sleep(250 * time.Millisecond)
    }


}
