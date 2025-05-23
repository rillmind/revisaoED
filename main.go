package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var lista ListaTarefa

func main() {
	sc := bufio.NewScanner(os.Stdin)

	for {
		menu()

		sc.Scan()
		opcao := sc.Text()
		fmt.Println()

		switch opcao {
		case "1":
			fmt.Println("Criando tarefa...")
			fmt.Print("Descrição: ")
			sc.Scan()
			descricao := sc.Text()
			fmt.Print("Prioridade: ")
			sc.Scan()
			prioridadeStr := sc.Text()
			prioridade, err := strconv.Atoi(prioridadeStr)

			if err != nil {
				fmt.Print(err)
			}

			tarefa := NewTarefa(descricao, prioridade)

			lista.AdicionarTarefa(tarefa)

		case "2":
			lista.Listar()

		case "3":
			fmt.Println("Removendo tarefa...")
			fmt.Print("Descrição: ")
			sc.Scan()
			descricao := sc.Text()
			lista.Remover(descricao)

		case "4":
			fmt.Print(lista.Tamanho())
		}

		if opcao == "0" {
			break
		}
	}
}

func menu() {
	fmt.Println()
	fmt.Println("1) Adicionar tarefa")
	fmt.Println("2) Listar tarefas")
	fmt.Println("3) Remover tarefa")
	fmt.Println("4) Tamanho da lista")
	fmt.Println("0) Sair")
	fmt.Print("Escolha: ")
}
