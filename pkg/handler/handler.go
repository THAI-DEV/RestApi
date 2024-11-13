package handler

import (
	"dechdev/pkg/config"
	"dechdev/pkg/util"
	"net/http"

	"github.com/THAI-DEV/dechutil"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"buildTime": config.BuildTime,
		"startTime": config.StartTime,
	})
}

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "--- pong ---")
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": dechutil.RandomString(8, true, true, true, true),
		"access":  util.CurrentDateTimeString(),
		"key":     config.Key,
	})
}
