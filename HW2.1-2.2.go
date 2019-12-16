package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Четные числа: ", i)
		}
	}

	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			fmt.Println("Числа делимые на 3 без остатка: ", i)
		}
	}
}
