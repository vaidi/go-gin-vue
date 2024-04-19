package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//type MsgInfoMiddleware struct {
//}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()
		fmt.Println("中间件开始了")
		c.Set("request", "中间件")
		c.Next()
		status := c.Writer.Status()
		println(status)
		since := time.Since(now)
		fmt.Println("总消耗时间：", since.String())
	}
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("token"); err == nil {
			if cookie.Value == "yel" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//拒绝
		c.Abort()
		return

	}
}
