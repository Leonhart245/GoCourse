package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var rub string
	const USD float64 = 64
	fmt.Println("Введите количество рублей: ")
	fmt.Scanln(&rub)
	num, err := strconv.Atoi(rub)
	if err != nil {
		log.Fatalln(err)
	}

	var result float64
	result = float64(num) / USD
	fmt.Println("Получите: ", result, "$")
}
