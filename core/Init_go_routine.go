package core

import (
	"go-gin-vue/core/go_runtime"
)

func init() {
	go_runtime.GoAdd()
	//go_runtime.GoTimer()
	go_runtime.SingleChannel()
	go_runtime.GoChannelGrace()
	go_runtime.GoChannelCaChe()
	go_runtime.GoChannel()
	go_runtime.GoMaxProc()
	go_runtime.GoExit()
	go_runtime.GoRoutine()
	go_runtime.GoRunTime()
}
