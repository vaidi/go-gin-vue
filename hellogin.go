package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "go-gin-vue/core"
	"go-gin-vue/core/core_inside"
	"go-gin-vue/route"
)

func main() {
	engine := gin.Default()
	//设置中间件
	engine.Use()
	route.UserController{}.UserInfo(engine)
	route.ChatController{}.Chat(engine)
	route.RedirectController{}.Redirect(engine)
	engine.Run(":" + core_inside.ServerConfig.ServerPort)
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
func submit(c *gin.Context) {
	param := c.Param("paramkey")
	c.String(200, fmt.Sprintf("hello %s\n", param))
}

/**
//api 参数
engine.GET("/user/:name/*action", func(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	//截取/
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
})
//url 参数
engine.GET("/user", func(c *gin.Context) {
	//指定默认值
	//http://localhost:8080/user 才会打印出来默认的值
	name := c.DefaultQuery("name", "枯藤")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
})
//表单参数
engine.POST("/form", func(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("userpassword")
	// c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	c.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
})

//限制上传最大尺寸 上传文件参数
engine.MaxMultipartMemory = 8 << 20
engine.POST("/upload", func(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	// c.JSON(200, gin.H{"message": file.Header.Context})
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
})

// 限制表单上传大小 8MB，默认为32MB 上传多个文件
engine.MaxMultipartMemory = 8 << 20
engine.POST("/upload", func(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
	}
	// 获取所有图片
	files := form.File["files"]
	// 遍历所有图片
	for _, file := range files {
		// 逐个存
		if err := c.SaveUploadedFile(file, file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
			return
		}
	}
	c.String(200, fmt.Sprintf("upload ok %d files", len(files)))
})

//自定义分组路由
v1 := engine.Group("/v1")
{
	v1.GET("/login", login)
	v1.GET("submit", submit)
}
v2 := engine.Group("v2")
{
	v2.POST("/login", login)
	v2.POST("/submit", submit)
}
*/
