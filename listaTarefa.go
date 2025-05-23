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
			return true
		} else if atual.tarefa.descricao == descricao {
			anterior.proximo = atual.proximo
			return true
		}
		anterior = atual
		atual = atual.proximo
	}
	return false
}

func (lt *ListaTarefa) EstaVazia() bool {
	return lt.Tamanho() == 0
}

func (lt *ListaTarefa) Tamanho() int {
	cont := 0
	atual := lt.inicio

	for atual != nil {
		cont += 1
		atual = atual.proximo
	}
	return cont
}
