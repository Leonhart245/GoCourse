package main

import (
    "fmt"
    //"time"
)

func main() {
    ch := make(chan int, 3)
    ch <- 1
    ch <- 2
    ch <- 3

    go func() {
        ch <- 4
    }()

    //time.Sleep(time.Second)

    for i := 0; i < 4; i++ {
        fmt.Println(<-ch)
    }
    fmt.Println(ch)
}