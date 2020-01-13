package main

import (
    "bytes"
    "encoding/csv"
    "fmt"
)

func main() {
    data := bytes.NewBufferString(`COL1,COL2
val1,val2
val1,val2
val1,val2`)

    reader := csv.NewReader(data)
    reader.Read()
    lines, err := reader.ReadAll()

    if err != nil {
        panic(err)
    }

    fmt.Println(lines)
}