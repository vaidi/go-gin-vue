package route

import "github.com/gin-gonic/gin"

type ChatController struct {
}

func (chat ChatController) Chat(c *gin.Engine) {

	group := c.Group("/chat")
	group.Use()
	{
		group.GET("/girl", func(c *gin.Context) {
			c.String(200, "hello girl my love")
		})
		group.GET("/boy", func(c *gin.Context) {})

	}

}
