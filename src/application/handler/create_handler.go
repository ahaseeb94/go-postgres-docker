package handler

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/service"
	"net/http"
)

func CreateTodo(c *gin.Context) {
	var todo service.TodoInput
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.CreateTodo(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
