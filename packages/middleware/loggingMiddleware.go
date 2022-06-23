package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware(c *gin.Context) {
	log.Println("requested url: " + c.Request.URL.Path)
	c.Next()
}
