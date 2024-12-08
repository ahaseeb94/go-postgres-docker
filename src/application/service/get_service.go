package service

import (
	"errors"
	"go-postgres-docker/src/infrastructure/persistence"
)

func GetTodoByID(id string) (persistence.Todo, error) {
	var todo persistence.Todo

	// Find the todo by ID
	if err := persistence.DB.First(&todo, id).Error; err != nil {
		return persistence.Todo{}, errors.New("todo not found")
	}

	return todo, nil
}
