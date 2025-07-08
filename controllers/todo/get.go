package todo

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/gin-gonic/gin"
)

// GetTodo godoc
// @Summary Retrieve todos
// @Description Returns all todos (auth required)
// @Tags Todos
// @Accept  json
// @Produce  json
// @Success 200 {object} models.SwaggerGetTodosResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Security BearerAuth
// @Router /todos/ [get]
func GetTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	var todos []models.Todo
	if err := db.DbClient.
		Where("user_id = ? AND is_deleted = false", userID).
		Order("created_at DESC").
		Find(&todos).Error; err != nil {
		//return error
		c.JSON(500, gin.H{"error": "Failed to retrieve todos"})
		return
	}

	c.JSON(200, gin.H{
		"message": "Success retrieve todo list",
		"todos":   todos,
	})
}
