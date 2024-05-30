package main

import (
	"net/http"
	"os"

	"github.com/0zl/pikacha/configuration"
	"github.com/gin-gonic/gin"
)

func main() {
	config, _ := configuration.NewConfig()

	err := os.Mkdir("static", 0775)
	if err != nil {
		println("static directory already exists.")
	}

	gin.SetMode(gin.ReleaseMode)
	sv := gin.Default()

	sv.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "https://github.com/0zl/pikacha")
	})

	sv.Static("/nya", "./static")

	sv.POST("/upload", func(c *gin.Context) {
		ProcessUploader(c, &config)
	})

	sv.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "OwO?")
	})

	println("pikacha is running.")
	sv.Run(config.GetAdress())
}
