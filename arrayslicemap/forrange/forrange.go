package main

import "fmt"

func main() {
	numeros := [...]int{10, 20, 30, 40, 50} // compilador conta!

	for i, numero := range numeros {
		fmt.Printf("%d) %d\n", i+1, numero)
	}

	for _, num := range numeros {
		fmt.Println(num)
	}
}
