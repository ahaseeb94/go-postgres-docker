package handler

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/service"
	"net/http"
)

func GetAllTodos(c *gin.Context) {
	todos, err := service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}
