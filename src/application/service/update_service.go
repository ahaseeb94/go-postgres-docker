package service

import (
	"errors"
	"go-postgres-docker/src/infrastructure/persistence"
)

func UpdateTodo(id string, input TodoInput) (persistence.Todo, error) {
	var todo persistence.Todo

	// Find the todo by ID
	if err := persistence.DB.First(&todo, id).Error; err != nil {
		return persistence.Todo{}, errors.New("todo not found")
	}

	// Update the fields
	todo.Title = input.Title
	todo.Description = input.Description
	todo.Completed = input.Completed

	// Save the changes
	if err := persistence.DB.Save(&todo).Error; err != nil {
		return persistence.Todo{}, err
	}

	return todo, nil
}
