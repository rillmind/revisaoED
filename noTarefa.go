package main

type NoTarefa struct {
	tarefa  *Tarefa
	proximo *NoTarefa
}

func NewNoTarefa(tarefa *Tarefa) *NoTarefa {
	return &NoTarefa{
		tarefa:  tarefa,
		proximo: nil,
	}
}
