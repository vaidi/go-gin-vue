package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func TimeMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		since := time.Since(start)
		fmt.Println("程序用时：", since)
	}

}
