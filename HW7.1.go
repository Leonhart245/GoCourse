package main

import (
	"fmt"
	"time"
)

const N = 10

// Spinner demo in goroutine
func main() {
	go spinner(50 * time.Millisecond)
	//	const n = 45
	//	fibN := fibonacci(n)
	//	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)

	time.Sleep(N * time.Second) // need sleep for N seconds only
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}