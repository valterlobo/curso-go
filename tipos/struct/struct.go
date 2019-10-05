package main

import "fmt"

type produto struct {
	cod      string
	nome     string
	preco    float64
	desconto float64
}

// Método: função com receiver (receptor)
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p produto) label() string {
	return p.cod + "-" + p.nome
}

func main() {
	var produto1 produto

	produto1 = produto{
		cod:      "pap-lap",
		nome:     "Lapis",
		preco:    1.79,
		desconto: 0.05,
	}

	fmt.Println(produto1.label())
	fmt.Println(produto1, produto1.precoComDesconto())

	produto2 := produto{"inf-note", "Notebook", 1989.90, 0.10}
	fmt.Println(produto2.label())
	fmt.Println(produto2, produto2.precoComDesconto())
}
