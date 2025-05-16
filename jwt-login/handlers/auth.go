package handlers

import (
	"jwt-login/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var users = make(map[string]string)

var mySingIngKey = []byte("my_secret_key")

type ErrerResponse struct {
	Message string `json:"message"`
}

func Register(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, ErrerResponse{Message: "Invalid request"})
		return
	}
	if hasedPassword, err := utils.HashPassword(user.Password); err != nil {
		c.JSON(500, ErrerResponse{Message: "Internal server error"})
		return
	} else {
		users[user.Username] = hasedPassword
		c.JSON(200, gin.H{"message": "User registered successfully"})
	}

}

func Login(c *gin.Context) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, ErrerResponse{Message: "Invalid request"})
		return
	}
	storedPassword, ok := users[user.Username]
	if !ok {
		c.JSON(401, ErrerResponse{Message: "Invalid username or password"})
		return
	}
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password))
	if err != nil {
		c.JSON(401, ErrerResponse{Message: "Invalid username or password"})
		return
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24),
	}).SignedString(mySingIngKey)
	if err != nil {
		c.JSON(500, ErrerResponse{Message: "Internal server error"})
		return
	}
	c.JSON(200, gin.H{"token": token})
}
