package main

import (
	"net/http"

	"github.com/0zl/pikacha/configuration"
	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configuration.NewConfig()

	gin.SetMode(gin.ReleaseMode)
	sv := gin.Default()

	sv.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pikacha!",
		})
	})

	println("pikacha is running.")
	sv.Run(config.GetAdress())
}
