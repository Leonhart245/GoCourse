package main

import (
	"fmt"
	"sort"
)

func main() {
	people := []struct {
		Name  string
		Phone int
	}{
		{"Stan", 123456789},
		{"Kile", 987654321},
		{"Kenny", 135792468},
		{"Cartman", 246813579},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("Сортировка по имени:", people)
}