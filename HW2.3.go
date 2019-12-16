package main

import (
	"fmt"
	"math/big"
)

func main() {
	F1 := big.NewInt(0)
	F2 := big.NewInt(1)
	for i := 1; i < 100; i++ {
		F1.Add(F1, F2)
		F1, F2 = F2, F1

		fmt.Printf("%v\n", F1)
	}
}
