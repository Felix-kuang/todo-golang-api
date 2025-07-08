package main

import (
	"github.com/Felix-kuang/todo-golang-api/db"
	"github.com/Felix-kuang/todo-golang-api/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db.Connect()

	router := gin.Default()

	auth := router.Group("/auth")
	routes.AuthRouter(auth)

	todo := router.Group("/todos")
	routes.TodoRoutes(todo)

	router.Run()
}
