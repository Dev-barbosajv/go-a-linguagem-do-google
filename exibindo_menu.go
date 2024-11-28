package main

import "fmt"

func main() {
	fmt.Println("1 - Iniciar monitoramento.")
	fmt.Println("2 - Exibir logs.")
	fmt.Println("0 -  Sair do programa.")

	var comando int
	fmt.Scanf("Escolha uma opção: %d", &comando)
	fmt.Println("O endereço da minha variavel é:", &comando)
	fmt.Println("O comando escolhido foi:", &comando)
}
