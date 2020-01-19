package main

import "github.com/gin-gonic/gin"
import "app/pkg/controllers/access"
import "app/pkg/controllers/healthcheck"

func main() {

	router := gin.Default()

	router.GET("/ping", access.Ping)
	router.GET("/healthcheck", healthcheck.Get)
	router.Run()
}
