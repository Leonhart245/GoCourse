package main

import "fmt"

func main() {
	var S [1000]int
	var srez []int
	S[1] = 0
	for k := 2; k < 1000; k++ {
		S[k] = 1
	}

	for k := 2; k*k < 1000; k++ {
		if S[k] == 1 {
			for l := k * k; l < 1000; l += k {
				S[l] = 0
			}
		}
	}
	for index, value := range S {
		if value == 1 {
			srez = append(srez, index)
		}
	}
	fmt.Println("Числа Эрастофена: ", srez)
}
