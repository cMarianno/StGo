package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	/* Call Function without Return */
	exibeIntroducao()

	for {
		exibeMenu()
		/* Call Function with Return */
		comando := leComando()
		/* Swith */
		switch comando {
		case 1:
			iniciarMonitoramento()
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
	}
}

/* Function without Return */
func exibeIntroducao() {
	nome := "Caio"
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
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/* Array
	var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"
	*/

	/* Slice */
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"}

	/* For
	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])
	}
	*/

	/* ForEach */
	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")

}

func testaSite(site string) {

	/* Ignore Some Returns */
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}
}
