package main

import (
	"github.com/bhushan-aruto/todo_go/command"
	"github.com/bhushan-aruto/todo_go/storage"
	"github.com/bhushan-aruto/todo_go/todo"
)

func main() {
	todos := todo.Todos{}
	storage := storage.NewStoarge[todo.Todos]("todos.json")
	if err := storage.Load(&todos); err != nil {
		todos = todo.Todos{}
	}
	storage.Load(&todos)
	cmdFlags := command.NewCmdFlag()
	cmdFlags.Execute(&todos)

	storage.Save(todos)

}
