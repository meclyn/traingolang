package main 

import (
    "fmt"
    "bufio"
    "os"
    "strconv"
    "strings"
)


func main (){
	// cria novo leitor para ler o input do user
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Digite um número: ")

	// lê o input do user ate um caractere
	input, _ := reader.ReadString('\n')

	// Remove espaço em branco no input
	input = strings.TrimSpace(input)

	//converte a string do input para int
	number, err := strconv.Atoi(input)
	if err != nil {
		fmt.PrintIn("Erro: por favor, digite um número válido.")
		return
	}

	// verifica se o numero é par ou impar
	if number%2 == 0 {
		fmt.Printf("O número %d é par.\n", number)
	} else {
		fmt.Printf("O número %d é ímparr.\n", number)
	}
}
