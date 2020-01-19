package access

import "github.com/gin-gonic/gin"

// GET /ping
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
		"status":  200,
	})
}
