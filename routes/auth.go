package routes

import (
	"github.com/Felix-kuang/todo-golang-api/controllers/auth"

	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.RouterGroup) {
	r.POST("/register", auth.Register)
	r.POST("/login", auth.Login)

}
