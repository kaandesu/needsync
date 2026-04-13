package main

import (
	"needsync/internal/needs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/sync", func(c *gin.Context) {
		err := needs.SyncCreate("./needs")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "synced",
		})
	})

	r.Run(":8080")
}
