package service

import "go-postgres-docker/src/infrastructure/persistence"

type TodoInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func CreateTodo(input TodoInput) (persistence.Todo, error) {
	todo := persistence.Todo{
		Title:       input.Title,
		Description: input.Description,
		Completed:   input.Completed,
	}

	if err := persistence.DB.Create(&todo).Error; err != nil {
		return persistence.Todo{}, err
	}

	return todo, nil
}
