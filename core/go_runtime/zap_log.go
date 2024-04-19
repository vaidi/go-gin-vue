package go_runtime

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var log *zap.Logger

func Log(c *gin.Context) {

	log = zap.NewExample()
	log.Debug("NewExample:" + "This is a DEBUG message")
	log.Info("NewExample:" + "This is an INFO message")
	log3, _ := zap.NewDevelopment()
	log3.Debug("NewDevelopment:" + "This is a DEBUG message")
	log3.Info("NewDevelopment:" + "This is an INFO message")
	log1, _ := zap.NewProduction()
	log1.Debug("NewProduction:" + "This is a DEBUG message")
	log1.Info("NewProduction:" + "This is an INFO message")

}
