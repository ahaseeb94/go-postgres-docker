package handler

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/service"
	"net/http"
)

func GetTodo(c *gin.Context) {
	id := c.Param("id")
	todo, err := service.GetTodoByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)
}
