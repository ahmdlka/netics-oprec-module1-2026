package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var timeStart = time.Now()

func main() {
	r := gin.Default()

	r.GET("/health", getHealth)

	r.Run(":8080")
}

func getHealth(c *gin.Context) {
	c.JSON(
		http.StatusOK,
		gin.H{
			"nama":      "Sersan Mirai Afrizal",
			"nrp":       "5025241999",
			"status":    "UP",
			"timestamp": time.Now().Format(time.RFC3339),
			"uptime":    time.Since(timeStart).String(),
		},
	)
}
