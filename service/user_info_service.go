package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-vue/dao"
	"net/http"
)

func ListUserInfo(c *gin.Context) {
	info := dao.UserInfo{Mobile: "17610721152", RoleId: 1, Password: "123456"}
	dao.UserInfo{}.Add(info)
	fmt.Println("调用query方法")
	query := dao.UserInfo{}.Query(1)
	fmt.Printf("查询到用户%v", query)
	c.JSON(http.StatusOK, gin.H{
		"userInfo": query,
	})
}
