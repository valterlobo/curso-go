package main

import (
	"fmt"
	"strings"
)

func main() {

	var str string
	str = "valter"

	fmt.Println("STRING", str)

	fmt.Println("length", len(str))

	for index, char := range str {
		fmt.Printf("character at index %d is %c\n", index, char)
	}

	s := `Hello,
	My Big Blue
	"World"!`

	fmt.Println(s)

	fmt.Println(strings.EqualFold("Go", "go"))
}
