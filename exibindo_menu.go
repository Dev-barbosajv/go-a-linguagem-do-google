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
	fmt.Printf("Monitorando...\n\n")

	sites := lendoNomeSite()

	for i, site := range sites {
		fmt.Printf("Posicao: %d | site: %s\n", i, site)

		resp, err := http.Get(site)
		if err != nil {
			fmt.Println("Ocorreu um erro:", err)
		}
		if resp.StatusCode == 200 {
			fmt.Printf("O site: %v está ativo e rodando.\n\n", site)
		} else {

			fmt.Printf("Site: %s, está com problema. %v", site, resp.StatusCode)
		}

	}

}

func lendoNomeSite() []string {

	var sites []string

	arquivo, err := os.Open("/Users/devjhonny/estudos/go-a-linguagem-do-google/sites.txt")
	fmt.Println(arquivo)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	return sites

}
