package application

import (
	"github.com/gin-gonic/gin"
	"go-postgres-docker/src/application/handler"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	r.POST("/todos", handler.CreateTodo)
	r.GET("/todos", handler.GetAllTodos)
	r.GET("/todos/:id", handler.GetTodo)
	r.PUT("/todos/:id", handler.UpdateTodo)
	r.DELETE("/todos/:id", handler.DeleteTodo)

	return r
}
