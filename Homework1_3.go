package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	var sum int
	var per int
	var trans string
	fmt.Println("Введите сумму вклада : ")
	fmt.Scanln(&trans)
	sum, err1 := strconv.Atoi(trans)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("Введите % : ")
	fmt.Scanln(&trans)
	per, err2 := strconv.Atoi(trans)
	if err2 != nil {
		log.Fatalln(err2)
	}
	var result, result2, result3, result4, result5 float64

	result = float64(sum) + float64(sum)*(float64(per)/100)
	result2 = float64(result) + float64(result)*(float64(per)/100)
	result3 = float64(result2) + float64(result2)*(float64(per)/100)
	result4 = float64(result3) + float64(result3)*(float64(per)/100)
	result5 = float64(result4) + float64(result4)*(float64(per)/100)

	fmt.Println("Сумма вклада за 1 год равна: ", result)
	fmt.Println("Сумма вклада за 2 год равна: ", result2)
	fmt.Println("Сумма вклада за 3 год равна: ", result3)
	fmt.Println("Сумма вклада за 4 год равна: ", result4)
	fmt.Println("Сумма вклада за 5 год равна: ", result5)
}
