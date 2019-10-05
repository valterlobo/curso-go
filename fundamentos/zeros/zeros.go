package main

import "fmt"

func main() {
	var a int
	var b float64
	var c bool
	var d string
	var e *int

	fmt.Printf("%v %v %v %q %v", a, b, c, d, e)
	fmt.Println("\nTypes")
	fmt.Printf("%t %t %t %t %t", a, b, c, d, e)
}
