package main

type Tarefa struct {
	descricao  string
	prioridade int
}

func NewTarefa(descricao string, prioridade int) *Tarefa {
	return &Tarefa{
		descricao:  descricao,
		prioridade: prioridade,
	}
}
