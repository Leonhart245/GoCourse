package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	var kat1 int
	var kat2 int
	var gip, plosh, per float64
	var trans string
	fmt.Println("Введите размер катета 1 : ")
	fmt.Scanln(&trans)
	kat1, err1 := strconv.Atoi(trans)
	if err1 != nil {
		log.Fatalln(err1)
	}
	fmt.Println("Введите размер катета 2 : ")
	fmt.Scanln(&trans)
	kat2, err2 := strconv.Atoi(trans)
	if err2 != nil {
		log.Fatalln(err2)
	}
	gip = math.Sqrt(math.Pow(float64(kat1), 2) + math.Pow(float64(kat2), 2))
	per = float64(kat1) + float64(kat2) + float64(gip)
	plosh = (float64(kat1) + float64(kat2)) / 2

	fmt.Println("Гипотенуза равна: ", gip)
	fmt.Println("Периметр равен: ", per)
	fmt.Println("Площадь равна: ", plosh)
}
