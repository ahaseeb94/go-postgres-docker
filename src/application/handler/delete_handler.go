package handler

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/service"
	"net/http"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	if err := service.DeleteTodoByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted successfully"})
}
