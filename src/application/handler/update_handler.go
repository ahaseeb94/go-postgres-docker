package handler

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/service"
	"net/http"
)

func UpdateTodo(c *gin.Context) {
	var todo service.TodoInput
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.Param("id")
	res, err := service.UpdateTodo(id, todo)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully", "todo": res})
}
