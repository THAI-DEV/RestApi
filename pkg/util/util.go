package util

import "time"

func CurrentDateTimeString() string {
	return time.Now().In(time.FixedZone("GMT+7", 7*3600)).Format("2006-01-02 15:04:05")
}
