package main

import (
	"net/http"

	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	start := time.Now()
	uptime := time.Since(start)
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		// fmt.Printf("Request received: %s\n", c.String())
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
			"uptime": uptime,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
