package main

import "fmt"

func main() {
	fmt.Println("1- Iniciar monitoramento.")
	fmt.Println("2- Exibir logs.")
	fmt.Println("0-  Sair do programa.")

	var comando int
	fmt.Scan(&comando)
	//fmt.Println("O comando escolhido foi:", comando)

	if comando == 1 {
		fmt.Println("Iniciar monitoramento.")
	} else if comando == 2 {
		fmt.Println("Exibir logs.")
	} else {
		fmt.Println("Sair do programa.")
	}
}
