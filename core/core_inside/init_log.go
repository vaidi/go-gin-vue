package core_inside

import (
	"go.uber.org/zap"
	"time"
)

// 用来初始化log的文件
func InitLog() {
	//设置全局的log
	//zap.ReplaceGlobals()

}
func initLog() {
	logger, _ := zap.NewProduction()
	defer logger.Sync() // flushes buffer, if any
	sugar := logger.Sugar()
	var url string = "xxh"
	sugar.Infow("failed to fetch URL",
		// Structured context as loosely typed key-value pairs.
		"url", url,
		"attempt", 3,
		"backoff", time.Second,
	)
	sugar.Infof("Failed to fetch URL: %s", url)

	log := zap.NewExample()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fileds",
		zap.Int("age", 24), zap.String("agender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
	log.Panic("this is panic message")
}

func logNewProtection() {

	log, _ := zap.NewProduction()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fileds",
		zap.Int("age", 24), zap.String("agender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
}

func logNewDevelopment() {
	logger, _ := zap.NewDevelopment()
	slogger := logger.Sugar()

	slogger.Debugf("debug message age is %d, agender is %s", 19, "man")
	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s", "sprintf")
	slogger.Infow("Infow() allows tags", "name", "Legolas", "type", 1)

}
