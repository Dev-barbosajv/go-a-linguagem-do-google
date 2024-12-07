package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
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
			fmt.Println("Ocorreu um erro ao acessar o site:", err)
			continue // Pula para o próximo site, caso haja erro
		}

		// Verifique se a resposta foi bem-sucedida
		if resp != nil && resp.StatusCode == 200 {
			fmt.Printf("O site: %v está ativo e rodando.\n\n", site)
		} else {
			// Tratar o caso onde a resposta é diferente de 200
			fmt.Printf("Site: %s está com problema. Status Code: %v\n", site, resp.StatusCode)
		}
	}
}

func lendoNomeSite() []string {
	var sites []string

	// Tente abrir o arquivo e verifique se houve erro
	arquivo, err := os.Open("/Users/devjhonny/estudos/go-a-linguagem-do-google/sites.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return nil // Retorna um slice vazio em caso de erro
	}
	defer arquivo.Close() // Certifique-se de fechar o arquivo ao final

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		// Só adiciona ao slice se a linha não for vazia
		if linha != "" {
			sites = append(sites, linha)
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println("Erro ao ler o arquivo:", err)
			break
		}
	}

	return sites
}
