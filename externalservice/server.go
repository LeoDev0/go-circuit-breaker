package externalservice

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitServer() {
	startTime := time.Now()
	e := gin.Default()
	e.GET("/ping", func(ctx *gin.Context) {
		if time.Since(startTime) < 5*time.Second {
			ctx.JSON(http.StatusInternalServerError, gin.H{"ping": "error"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"ping": "pong"})
	})

	e.Run(":8080")
}
