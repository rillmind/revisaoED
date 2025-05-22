package main

import "fmt"

type ListaTarefa struct {
	inicio *NoTarefa
}

func (lt *ListaTarefa) AdicionarTarefa(tarefa *Tarefa) {
	novo := NewNoTarefa(tarefa)

	if lt.inicio == nil {
		lt.inicio = novo

	} else {
		atual := lt.inicio

		for atual.proximo != nil {
			atual = atual.proximo
		}

		atual.proximo = novo
	}
}

func (lt *ListaTarefa) Listar() {
	atual := lt.inicio

	fmt.Println("Tarefas: {")
	for atual != nil {
		fmt.Println(
			"  {\n",
			" Descrição:", atual.tarefa.descricao, "\n",
			" Prioridade:", atual.tarefa.prioridade, "\n",
			" },",
		)

		atual = atual.proximo
	}
	fmt.Println("}")
}

func (lt *ListaTarefa) Remover(descricao string) bool {
	return true
}

func (lt *ListaTarefa) EstaVazia(descricao string) bool { return true }

func (lt *ListaTarefa) Tamanho() int { return 1 }
