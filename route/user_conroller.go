package route

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-vue/core/go_runtime"
	"go-gin-vue/middleware"
	"go-gin-vue/service"
	"log"
	"time"
)

type UserController struct {
	BaseController
}

func (user UserController) UserInfo(engine *gin.Engine) {
	//自定义分组路由
	group := engine.Group("/user")
	group.Use(middleware.MiddleWare())
	{
		group.GET("/login", middleware.TimeMiddleware(), func(context *gin.Context) {
			user.infoMsg(context)
			fmt.Println("你来登录了 小可爱")
			go_runtime.Log(context)
			//休眠5s
			time.Sleep(1 * time.Second)
			fmt.Println("登录结束了 小可爱")
			context.JSON(200, gin.H{
				"code":   200,
				"myLove": "xxh",
			})
			//这里一定要异步
			c := context.Copy()
			go func() {
				time.Sleep(2 * time.Second)
				c.Set("xxh", "love")
				log.Println("异步执行：" + c.Request.URL.Path + "," + c.GetString("xxh"))
			}()

		})
		group.GET("cookie", middleware.AuthMiddleWare(), func(context *gin.Context) {
			cookie, err := context.Cookie("key_cookie")
			if err != nil {
				cookie = "NotSet"
				context.SetCookie("key_cookie", "value_cooke", 20, "/", "localhost", false, true)
			}
			log.Println("cookie:" + cookie)

		})

		group.GET("submit", service.UserShowService{}.ShowSelf)
		group.GET("love", service.UserShowService{}.ShowMyWife)

	}

}
