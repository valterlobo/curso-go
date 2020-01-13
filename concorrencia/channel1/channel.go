package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	var inteiro = <-ch    // recebendo dados do canal (leitura)
	fmt.Println(inteiro)
	ch <- 2
	fmt.Println(<-ch)
}
