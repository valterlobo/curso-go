package main

import "fmt"

type item struct {
	produtoID int
	qtde      int
	preco     float64
}

type pedido struct {
	userID int
	itens  []item
	total  float64
}

func (i item) valorTotal() float64 {

	var total float64
	total = i.preco * float64(i.qtde)
	return total
}

func (p *pedido) valorTotal() float64 {
	total := 0.0
	for _, item := range p.itens {
		total += item.preco * float64(item.qtde)
	}
	p.total = total
	return total
}

func (p *pedido) addItem(i item) {

	p.itens = append(p.itens, i)
	p.valorTotal()
}

func (p pedido) printItens() {

	for _, item := range p.itens {
		fmt.Println("ITEM:", item)
		fmt.Println("TOTAL ITEM", item.valorTotal())
	}

}

func main() {
	pedido := pedido{
		userID: 1,
		itens: []item{
			item{produtoID: 1, qtde: 2, preco: 12.10},
			item{2, 1, 23.49},
			item{11, 100, 3.13},
		},
	}

	pedido.printItens()
	pedido.addItem(item{produtoID: 10, qtde: 2, preco: 22.10})
	fmt.Println(" ****************************")
	pedido.printItens()
	fmt.Printf("Valor total do pedido é R$ %.2f", pedido.valorTotal())
	fmt.Println(" Pedido:", pedido.total)
}
