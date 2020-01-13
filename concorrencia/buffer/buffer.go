package main

import (
	"fmt"
	"time"
)

func rotina(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	ch <- 5
	fmt.Println("Executou!")
	ch <- 6
}

func rotina_for(ch chan int) {

	var i int = 0 
	for i < 6 {
		i = i + 1
		ch <- i*10
	}
	
	fmt.Println("Executou! for chanel")

}

func main() {
	ch := make(chan int, 3)
	go rotina_for(ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	var i int = 0
	for i < 4 {
		i = i + 1
		fmt.Println(<-ch)
	}

}
