package core

import "go-gin-vue/core/core_inside"

func init() {
	core_inside.InitConfig()
	core_inside.InitLog()
	core_inside.DbSqlActuator()

}
