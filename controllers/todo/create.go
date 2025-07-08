package todo

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/gin-gonic/gin"
)

// CreateTodo godoc
// @Summary Create a new todo entry
// @Description Creates a todo (auth required)
// @Tags Todos
// @Accept  json
// @Produce  json
// @Param   input  body  models.SwaggerTodoCreate true  "Todo data"
// @Success 200 {object} models.SwaggerTodoCreateResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Security BearerAuth
// @Router /todos/ [post]
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
