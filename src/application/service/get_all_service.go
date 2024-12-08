package service

import "go-postgres-docker/src/infrastructure/persistence"

func GetAllTodos() ([]persistence.Todo, error) {
	var todos []persistence.Todo

	// Fetch all todos
	if err := persistence.DB.Find(&todos).Error; err != nil {
		return nil, err
	}

	return todos, nil
}
