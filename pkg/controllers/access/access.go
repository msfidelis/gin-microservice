package access

import "github.com/gin-gonic/gin"
import "github.com/go-redis/redis"

// GET /ping
func Ping(c *gin.Context) {

	client := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	result, err := client.Incr("counter").Result()
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"access":  result,
		"status":  200,
	})
}
