package handler

import (
	"dechdev/pkg/config"
	"dechdev/pkg/util"
	"net/http"
	"strings"

	"github.com/THAI-DEV/dechutil"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	c.JSON(200, gin.H{
		"lastDeploy": strings.Replace(config.BuildTime, "_", " ", -1),
		"lastStart":  config.StartTime,
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

func ReadData(c *gin.Context) {
	var responseError *string
	var data *string

	str := dechutil.RandomString(8, true, true, true, true)
	str = str + "|" + dechutil.TimeToStringDateTimeFull(dechutil.CurrentBkkTime())
	data = &str

	result, err := util.ReadFile("data.txt")
	if err != nil {
		str := err.Error()
		responseError = &str
		data = nil
	} else {
		responseError = nil
		strResult := string(result)
		data = &strResult
	}

	c.JSON(200, gin.H{
		"message": data,
		"error":   responseError,
	})
}

func WriteData(c *gin.Context) {
	var responseError *string
	var data *string

	str := dechutil.RandomString(8, true, true, true, true)
	str = str + "|" + dechutil.TimeToStringDateTimeFull(dechutil.CurrentBkkTime())
	data = &str

	err := util.WriteFile("data.txt", []byte(*data))
	if err != nil {
		str := err.Error()
		responseError = &str
	} else {
		responseError = nil
	}

	c.JSON(200, gin.H{
		"message": data,
		"error":   responseError,
	})
}
