package main

import (
	"fmt"

	"github.com/0zl/pikacha/configuration"
	"github.com/gin-gonic/gin"
)

func ProcessUploader(c *gin.Context, cfg *configuration.Config) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(400, gin.H{
			"error": "Failed to retrieve file",
		})
		return
	}

	if cfg.MaximumFileSize > 0 && file.Size > cfg.MaximumFileSize {
		c.JSON(400, gin.H{
			"error": "File too large",
		})
		return
	}

	err = c.SaveUploadedFile(file, fmt.Sprintf("%s/%s", "static", file.Filename))
	if err != nil {
		panic(err)
	}

	c.JSON(200, gin.H{
		"url": "/static/" + file.Filename,
	})
}
