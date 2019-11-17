package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("TESTE DATE")

	layoutISO := "2006-01-02"

	date, err := time.Parse(layoutISO, "2010-10-10")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Print(date)

}
