package main

import "fmt"

type User struct {
	Email string `json:"email" validate:"required,email"`
	Name  string `json:"name" validate:"required,min=2,max=100"`
}

func main() {
	u := User{}
	err := validate.Validate(a)

	fmt.Fatal(err)
}
