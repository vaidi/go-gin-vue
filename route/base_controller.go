package route

import (
	"github.com/gin-gonic/gin"
	"time"
)

type BaseController struct{}

func (base BaseController) infoMsg(c *gin.Context) {
	unix := time.Now().Unix()
	url := c.Request.URL
	println("header:" + url.String())

	println(time.Now().Unix() - unix)
}
