package main

import (
	"ToDo/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)
	router.GET("/todo/:id", handlers.GetToDo)
	router.POST("/edit/:id", handlers.EditToDo)
	router.POST("/delete/:id", handlers.DeleteToDo)
	router.POST("/add", handlers.AddToDo)
	router.POST("/validate/:username", handlers.CheckUsername)
	router.POST("/validate/", handlers.EmptyUsername)
	router.Run("localhost:4000")
}
