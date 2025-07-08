package main

import (
	"log"
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
	if os.Getenv("RAILWAY_ENVIRONMENT_ID") == "" {
		godotenv.Load()
	} else {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
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
