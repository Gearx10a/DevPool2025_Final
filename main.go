package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Register struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Fullname string `json:"fullname" binding:"required"`
}

type User struct {
	gorm.Model
	Username string
	Password string
	Fullname string
}

func main() {
	r := gin.Default()
	r.POST("/register", func(c *gin.Context) {
		var json Register
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// Create User
		encryptedPassword, _ := bcrypt.GenerateFromPassword([]byte(json.Password), 10)
		//user := User{Username: json.Username, Password: string(encryptedPassword), Fullname: json.Fullname}
		//if user.ID > 0 {
		//c.JSON(http.StatusOK, gin.H{"status": "ok", "message": "User Create Success", "userID": user.ID})
		//} else {
		//c.JSON(http.StatusOK, gin.H{"status": "error", "message": "User Create Failed"})
		//}
		c.JSON(200, gin.H{
			"register":         json,
			"encrytedPassword": encryptedPassword,
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"massage": "pong",
		})
	})
	r.Run("localhost:8080")
}
