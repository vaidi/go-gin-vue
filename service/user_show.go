package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserShowService struct{}

// 展示自己
func (user UserShowService) ShowSelf(c *gin.Context) {
	c.JSON(200, gin.H{
		"myLove":    "xxh！",
		"beautiful": "大漂亮啊",
	})
}

type Wife struct {
	Age      int    `json:"age" form:"age"`
	Name     string `json:"name" form:"name"`
	Birthday string `json:"birthday" form:"birthday"`
}

func (user UserShowService) ShowMyWife(c *gin.Context) {
	var wife Wife
	if err := c.ShouldBind(&wife); err != nil {
		c.JSON(500, fmt.Sprint(err))
		return
	}
	c.String(http.StatusOK, "你好啊%s", wife.Name)

}

func ChangeAge() {

}
