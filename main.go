package main

import (
	"fmt"
	"os"

	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/routes"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"

	_ "github.com/Felix-kuang/todo-golang-api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	if os.Getenv("RAILWAY_STATIC_URL") == "" {
		err := godotenv.Load()
		if err != nil {
			fmt.Println("Local .env not found, but that's OK in production.")
		}
	}
}

func main() {
	db.Connect()

	router := gin.Default()

	//Documentation
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	routes.AuthRouter(auth)

	todo := router.Group("/todos")
	routes.TodoRoutes(todo)

	router.Run()
}
