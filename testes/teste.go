package main

import "fmt"

type Task struct {
	ID   int
	NAME string
}

// Definição do método Process associado à struct Task
func (t *Task) Process() {
	fmt.Printf("Processing task with ID: %d\n", t.ID)
	fmt.Printf("Name task with ID: %v\n", t.NAME)
}

func Teste() {
	//cria uma instancia de Task
	task := Task{ID: 42}
	task.NAME = "Teste"

	//chama o método Process na instância11111111111111111111111111111111111111111111111111111111111 de Task
	task.Process()

}
