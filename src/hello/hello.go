package main

import (
	"fmt"
	"os"
)

func main() {
	/* Call Function without Return */
	exibeIntroducao()
	exibeMenu()
	/* Call Function with Return */
	comando := leComando()

	/*
		If
			if comando == 1 {
				fmt.Println("Monitorando...")
			} else if comando == 2 {
				fmt.Println("Exibindo Logs...")
			} else if comando == 0 {
				fmt.Println("Saindo do programa...")
			} else {
				fmt.Println("Não conheço este comando")
			}
	*/

	/* Swith */
	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		/* Exit Success */
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		/* Exit Error */
		os.Exit(-1)
	}
}

/* Function without Return */
func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

/* Function with Int Return */
func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
