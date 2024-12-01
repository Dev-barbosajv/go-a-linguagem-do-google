package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("1- Iniciar monitoramento.")
	fmt.Println("2- Exibir logs.")
	fmt.Println("0-  Sair do programa.")

	comando := commandLeen()

	switch comando {
	case 1:
		iniciandoMonitoramento()
	case 2:
		fmt.Println("Exibir logs.")
	case 0:
		fmt.Println("Sair do programa.")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)

	}
}

func commandLeen() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}

func iniciandoMonitoramento() {
	fmt.Println("Monitorando...")

	site := "https://www.alura.com.br"
	resp, _ := http.Get(site)
	fmt.Println(resp)
}
