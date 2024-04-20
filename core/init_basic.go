package core

import (
	"go-gin-vue/core/basic_core"
	"go-gin-vue/core/core_inside"
)

// 检验标准库的东西
func init() {
	basic_core.InitContext()
	basic_core.LogInit()
	basic_core.TimeInit()
	basic_core.FmtInit{}.BasicPrint()
	core_inside.DbSqlActuator()
}
