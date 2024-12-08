package service

import (
	"errors"
	"go-postgres-docker/src/infrastructure/persistence"
)

func DeleteTodoByID(id string) error {
	var todo persistence.Todo

	// Find the todo by ID
	if err := persistence.DB.First(&todo, id).Error; err != nil {
		return errors.New("todo not found")
	}

	// Delete the todo
	if err := persistence.DB.Delete(&todo).Error; err != nil {
		return err
	}

	return nil
}
