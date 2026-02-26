package routes

import (
	"todolist/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)

	authenticated.GET("/todos", getTodos)
	authenticated.POST("/todos", createTodo)
	authenticated.GET("/todos/:id", getTodo)
	authenticated.PUT("/todos/:id", updateTodo)
	authenticated.DELETE("/todos/:id", deleteTodo)
}
