package todo

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/gin-gonic/gin"
)

// ToggleTodo godoc
// @Summary Toggles a new todo entry
// @Description Toggle a todo to done/undone (auth required)
// @Tags Todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.SwaggerToggleSuccessResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Security BearerAuth
// @Router /todos/{id} [patch]
func ToggleTodo(c *gin.Context) {
	id := c.Param("id")
	var todo models.Todo
	userID := c.MustGet("user_id").(uint)

	if err := db.DbClient.
		Where("id = ? AND user_id = ?", id, userID).
		First(&todo).Error; err != nil {
		c.JSON(404, gin.H{"error": "Todo not found"})
		return
	}

	todo.Done = !todo.Done

	if err := db.DbClient.Save(&todo).Error; err != nil {
		c.JSON(500, gin.H{"error": "Failed to toggle todo"})
		return
	}

	c.JSON(200, gin.H{"message": "Todo Updated", "todo": todo})
}
