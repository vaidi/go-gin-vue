package core_inside

import (
	"fmt"
	"go-gin-vue/global"
	"gopkg.in/ini.v1"
)

var MysqlConfig global.MySqlConfig
var ServerConfig global.ServerConfig

func InitConfig() {
	fmt.Println("用来初始化配置")
	cfg, _ := ini.Load("./config/my.ini")
	fmt.Println("获取到配置", cfg.Section("mysql").Key("datasource").String())
	var config global.Config
	err := ini.MapTo(&config, "./config/my.ini")
	if err != nil {
		fmt.Printf("load ini failed,err:%v\n", err)
		return
	}
	MysqlConfig = config.MysqlConfig
	ServerConfig = config.ServerConfig
	fmt.Println("获取到配置", config, ServerConfig, MysqlConfig)
}
