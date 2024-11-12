package handler

import (
	"dechdev/config"
	"net/http"
	"time"

	"github.com/THAI-DEV/dechutil"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "-- pong --")
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": dechutil.RandomString(8, true, true, true, true),
		"access":  time.Now().In(time.FixedZone("GMT+7", 7*3600)).Format("2006-01-02 15:04:05"),
		"key":     config.Key,
	})
}
