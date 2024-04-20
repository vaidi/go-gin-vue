package basic_core

import (
	"fmt"
	"log"
	"os"
)

func LogInit() {
	logTest1()
	logTest2()

}
func logTest1() {
	log.Printf("这是一条%s日志\n", "普通的")
	//log.Fatalln("这是一条fatal的日志")
	//这个会引发推出
	log.SetPrefix("[pprof]")
	//log.Panicln("这是一条panic的日志")
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("这是一条很普通的日志。")
}
func logTest2() {
	file, err := os.OpenFile("./test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed", err)
		return
	}
	log.SetOutput(file)
	log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
	log.Println("记录日志啊")
	log.SetPrefix("【大漂亮】")
	log.Println("增加前缀后的日志")

}
