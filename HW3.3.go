package main

import (
	"fmt"
	"turn"
)

func main()  {
	turn.Push("первый текст")
	turn.Push("второй текст")
	turn.Push("третий текст")
	fmt.Println(turn.Bob())
	turn.Push("четвертый текст")
	fmt.Println(turn.Bob())
	fmt.Println(turn.Bob())
	fmt.Println(turn.Bob())
}