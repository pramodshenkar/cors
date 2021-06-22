package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("student")
	router.Use(cors.Default())
	{
		v1.POST("/login", loginHandleFunc)
	}

	v2 := router.Group("")
	v2.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5500"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	v2.Use(Authenticate())
	{
		v2.POST("/dashboard", handleFunc)
	}

	router.Run(":5000")
}

func handleFunc(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Dashboard"})
}

func loginHandleFunc(c *gin.Context) {
	var student struct {
		Name string `json:"name"`
	}

	if c.BindJSON(&student) != nil {
		c.JSON(400, gin.H{"message": "Provide required details."})
		c.Abort()
		return
	}

	token := fmt.Sprintf("%s%s", "iamtokenfor", student.Name)

	c.JSON(200, gin.H{"token": token})
}

func Authenticate() gin.HandlerFunc {

	return func(c *gin.Context) {

		requiredToken := c.Request.Header["Authorization"]

		if len(requiredToken) == 0 {
			c.JSON(400, gin.H{"message": "No token found"})
			c.Abort()
			return
		}

		fmt.Println(requiredToken)

		c.Next()
	}
}
