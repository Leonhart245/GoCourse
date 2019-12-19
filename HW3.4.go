package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Message struct {
    Name string
    Body string
	Time int64
	Phone int
}

func main()  {
	a := Message{"Alice", "Hello", 214124231451513, 1294706395881547000}

   b, err := json.Marshal(a)
    if err != nil {
        log.Println(err)
        return
    }

	err1 := ioutil.WriteFile("HW4Date.txt", b, 0644)
	if err1 != nil{
		log.Fatal(err1)
	}
    fmt.Println(a)
}

