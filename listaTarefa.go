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

	fmt.Print("Tarefas: {")
	for atual != nil {
		fmt.Print(`
  {
  Descrição:`, atual.tarefa.descricao, `
  Prioridade:`, atual.tarefa.prioridade, `
  },`,
		)

		atual = atual.proximo
	}
	fmt.Println("\n}")
}

func (lt *ListaTarefa) Remover(descricao string) bool {
	var anterior *NoTarefa
	atual := lt.inicio

	for atual != nil {
		if atual.tarefa.descricao == descricao && atual == lt.inicio {
			lt.inicio = atual.proximo
		} else if atual.tarefa.descricao == descricao {
			anterior.proximo = atual.proximo
		}
		anterior = atual
		atual = atual.proximo
	}
	return false
}

func (lt *ListaTarefa) EstaVazia(descricao string) bool { return true }

func (lt *ListaTarefa) Tamanho() int { return 1 }
