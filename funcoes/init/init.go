package main

import "fmt"

var teste string

func init() {
	teste = "init"
	fmt.Println("Inicializando...")
}

func setStatus(sts string) {
	teste = sts
}

func main() {
	fmt.Println("Main...", teste)
	setStatus("main")
	fmt.Println("Main...", teste)
}
