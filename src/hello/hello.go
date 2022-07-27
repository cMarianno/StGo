package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Caio"
	idade := 24
	versao := 1.2
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("O tipo da variavel nome é", reflect.TypeOf(nome))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando)
	fmt.Println("O endereço da minha variavel comando é", &comando)
	fmt.Println("O comando escolhido foi", comando)
}
