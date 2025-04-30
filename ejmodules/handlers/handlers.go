package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Helloworld(c *gin.Context) {
	fmt.Println("Hello World")
}
