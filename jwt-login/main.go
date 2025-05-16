package main

import (
	"fmt"
	"jwt-login/handlers"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	fmt.Println("Server running on port 8080")
	log.Fatal(router.Run(":8080"))

	router.Run()
}
