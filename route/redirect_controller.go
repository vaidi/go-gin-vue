package route

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RedirectController struct{}

func (red RedirectController) Redirect(engin *gin.Engine) {
	engin.GET("redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/user/login")
	})

}
