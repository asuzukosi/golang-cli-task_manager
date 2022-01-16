package core

import (
	"fmt"
	"task_manager/database"
)

type TaskManagerService struct {
}

func NewTaskManagerService() *TaskManagerService {
	return &TaskManagerService{}
}

func (s *TaskManagerService) AddTodo(name string) error {
	todo, err := database.AddTask(name)
	if err != nil {
		return err
	}
	fmt.Println("New todo item added")
	fmt.Println(todo)
	return nil
}

func (s *TaskManagerService) ListTodos() error {
	todos, err := database.GetAllTaskItems()
	if err != nil {
		return err
	}
	if len(todos) == 0 {
		fmt.Println("No todo items!! use the add command to add todo items.")
		return nil
	}
	for _, todo := range todos {
		fmt.Printf("%d) %s\n", todo.Id, todo.Name)
	}
	return nil
}

func (s *TaskManagerService) DoTodo(id int) error {
	err := database.RemoveTaskItem(id)
	if err != nil {
		return err
	}
	fmt.Println("Todo item of id: ", id, " removed")
	return nil
}
