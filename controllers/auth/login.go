package auth

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/models"
	"github.com/Felix-kuang/todo-golang-api/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Login godoc
// @Summary Login User
// @Description Register by username & password
// @Tags Auth
// @Accept  json
// @Produce  json
// @Param   input  body  models.SwaggerLoginInput true  "User credentials"
// @Success 200 {object} models.SwaggerSuccessLoginResponse
// @Failure 400 {object} models.SwaggerErrorResponse
// @Router /auth/login [post]
func Login(c *gin.Context) {
	//cek input
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&input); err != nil || input.Username == "" || input.Password == "" {
		c.JSON(400, gin.H{"error": "Username or password is missing"})
		return
	}

	//ambil data user
	var user models.User
	if err := db.DbClient.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "Invalid credentials"})
		return
	}

	// validasi password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate token", "message": err})
		return
	}

	//return user
	c.JSON(200, gin.H{
		"message": "Login success",
		"token":   token,
	})
}
