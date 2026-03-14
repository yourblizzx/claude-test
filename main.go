package main

import (
	"fmt"

	"claude-test/calc"
)

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	values, err := calc.ComputeValues(5)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	for _, v := range values {
		fmt.Println("i =", v)
	}
}
