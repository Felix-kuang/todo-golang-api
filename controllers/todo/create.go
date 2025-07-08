package todo

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(401, gin.H{"error": "User not authenticated"})
		return
	}

	var input struct {
		Title string `json:"title"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "Invalid input"})
		return
	}

	todo := models.Todo{
		Title:  input.Title,
		UserID: userID.(uint),
	}

	if err := db.DbClient.Create(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(201, gin.H{"message": "Todo created", "todo": todo})
}
