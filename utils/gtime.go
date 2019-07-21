package utils

import "time"

//NowDateTime 返回当前时间 如：2019-01-01 00:00:00
func NowDateTime () string {
	var timeFmt = "2006-01-02 15:04:05"
	return time.Now().Format(timeFmt)	
}