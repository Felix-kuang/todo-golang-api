package routes

import (
	"github.com/Felix-kuang/todo-golang-api/controllers/todo"
	"github.com/Felix-kuang/todo-golang-api/middleware" // kalau udah ada JWT middleware
	"github.com/gin-gonic/gin"
)

func TodoRoutes(r *gin.RouterGroup) {
	protected := r.Group("/")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/", todo.GetTodo)
	protected.POST("/", todo.CreateTodo)
	protected.PATCH("/:id", todo.ToggleTodo)
	protected.DELETE("/:id", todo.DeleteTodo)
}
