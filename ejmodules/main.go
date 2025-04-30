package main

import (
	"ejmodules/handlers"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", handlers.Helloworld)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println("Servidor iniciado 8080")
}
