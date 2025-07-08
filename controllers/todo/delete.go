package todo

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/gin-gonic/gin"
)

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	userID := c.MustGet("user_id").(uint)

	var todo models.Todo
	if err := db.DbClient.
		Where("id = ? AND user_id = ?", id, userID).
		First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	todo.IsDeleted = true

	if err := db.DbClient.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete todo"})
		return
	}

	c.JSON(200, gin.H{"message": "Todo deleted"})
}
